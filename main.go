package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// doGet request to orion url and parse result as json
func doGet(ctx context.Context, client *http.Client, url string, timeout time.Duration, data interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to build request to '%s': %w", url, err)
	}
	resp, err := client.Do(req)
	if resp != nil && resp.Body != nil {
		defer func() {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}()
	}
	if err != nil {
		return fmt.Errorf("failed to perform request to '%s': %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response from '%s': %s", url, resp.Status)
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(data); err != nil {
		return fmt.Errorf("failed to decode metrics reply from '%s': %w", url, err)
	}
	return nil
}

// scrape a cygnus server and write the metrics to the given writer
func scrape(ctx context.Context, client *http.Client, metricsURL, statsURL string, timeout time.Duration, w io.Writer) error {
	cancelCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()
	var (
		metrics    MetricsReply
		stats      StatsReply
		metricsErr error
		statsErr   error
		writeErr   error
		wg         sync.WaitGroup
	)
	metricsErr = doGet(cancelCtx, client, metricsURL, timeout, &metrics)
	// scrape stats in parallel, while sending metrics to client
	wg.Add(1)
	go func() {
		defer wg.Done()
		// start a new cancellation timer
		cancelCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer cancel()
		statsErr = doGet(cancelCtx, client, statsURL, timeout, &stats)
	}()
	if metricsErr == nil {
		// No error collecting metrics, try to send them
		writeErr = metrics.WritePrometheus(w)
	}
	wg.Wait()
	if statsErr == nil && writeErr == nil {
		// No error collecting stats, try to send them
		writeErr = stats.WritePrometheus(w)
	}
	return errors.Join(metricsErr, statsErr, writeErr)
}

// metrics http handler to export prometheus metrics
func metrics(metricsURL, statsURL string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer func() {
				io.Copy(ioutil.Discard, r.Body)
				r.Body.Close()
			}()
		}
		ctx := r.Context()
		client := http.DefaultClient
		if err := scrape(ctx, client, metricsURL, statsURL, READ_TIMEOUT, w); err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/metrics", http.StatusMovedPermanently)
}

const (
	READ_TIMEOUT  = 5 * time.Second
	WRITE_TIMEOUT = 2 * READ_TIMEOUT
	IDLE_TIMEOUT  = 5 * time.Second
	HEADER_SIZE   = 1 << 20
)

func main() {

	port := flag.Int("port", 8000, "port to listen on")
	metricsURL := flag.String("metrics", "http://localhost:1026/admin/metrics", "metrics URL")
	statsURL := flag.String("stats", "http://localhost:1026/statistics", "stats URL")
	flag.Parse()

	mux := &http.ServeMux{}
	mux.Handle("/metrics", metrics(*metricsURL, *statsURL))
	mux.Handle("/", http.HandlerFunc(home))

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", *port),
		Handler:        mux,
		ReadTimeout:    READ_TIMEOUT,
		WriteTimeout:   WRITE_TIMEOUT,
		IdleTimeout:    IDLE_TIMEOUT,
		MaxHeaderBytes: HEADER_SIZE,
	}
	log.Printf("starting service at port %d", *port)
	panic(server.ListenAndServe())
}
