package main

import (
	"fmt"
	"io"
)

// StatsReply replicates the format of the reply from the stats endpoint
type StatsReply struct {
	Counters struct {
		JsonRequests      float64 `json:"jsonRequests"`
		NoPayloadRequests float64 `json:"noPayloadRequests"`
		Requests          map[string]struct {
			Get    float64 `json:"GET"`
			Post   float64 `json:"POST"`
			Put    float64 `json:"PUT"`
			Patch  float64 `json:"PATCH"`
			Delete float64 `json:"DELETE"`
			Head   float64 `json:"HEAD"`
		} `json:"requests"`
		RequestsLegacy map[string]struct {
			Get    float64 `json:"GET"`
			Post   float64 `json:"POST"`
			Put    float64 `json:"PUT"`
			Patch  float64 `json:"PATCH"`
			Delete float64 `json:"DELETE"`
			Head   float64 `json:"HEAD"`
		} `json:"requestsLegacy"`
		InvalidRequests   float64 `json:"invalidRequests"`
		NotificationsSent float64 `json:"notificationsSent"`
	} `json:"counters"`
	Semwait struct {
		Request           float64 `json:"request"`
		DbConnectionPool  float64 `json:"dbConnectionPool"`
		Transaction       float64 `json:"transaction"`
		SubCache          float64 `json:"subCache"`
		ConnectionContext float64 `json:"connectionContext"`
		TimeStat          float64 `json:"timeStat"`
		Metrics           float64 `json:"metrics"`
	} `json:"semwait"`
	Timing struct {
		Accumulated struct {
			JsonV1Parse      float64 `json:"jsonV1Parse"`
			JsonV2Parse      float64 `json:"jsonV2Parse"`
			MongoBackend     float64 `json:"mongoBackend"`
			MongoReadWait    float64 `json:"mongoReadWait"`
			MongoWriteWait   float64 `json:"mongoWriteWait"`
			MongoCommandWait float64 `json:"mongoCommandWait"`
			Render           float64 `json:"render"`
			Total            float64 `json:"total"`
		} `json:"accumulated"`
		Last struct {
			JsonV2Parse    float64 `json:"jsonV2Parse"`
			MongoBackend   float64 `json:"mongoBackend"`
			MongoReadWait  float64 `json:"mongoReadWait"`
			MongoWriteWait float64 `json:"mongoWriteWait"`
			Render         float64 `json:"render"`
			Total          float64 `json:"total"`
		} `json:"last"`
	} `json:"timing"`
	NotifQueue struct {
		In             float64 `json:"in"`
		Out            float64 `json:"out"`
		Reject         float64 `json:"reject"`
		SentOk         float64 `json:"sentOk"`
		SentError      float64 `json:"sentError"`
		TimeInQueue    float64 `json:"timeInQueue"`
		AvgTimeInQueue float64 `json:"avgTimeInQueue"`
		Size           float64 `json:"size"`
	} `json:"notifQueue"`
	UptimeInSecs            float64 `json:"uptime_in_secs"`
	MeasuringIntervalInSecs float64 `json:"measuring_interval_in_secs"`
}

// Write stats data in plain text format
type statPrinter interface {
	write(w io.Writer) (int, error)
}

// plainStat is a statPrinter for a single value
type plainStat struct {
	text  string
	value float64
}

func (p plainStat) write(w io.Writer) (int, error) {
	return fmt.Fprintf(w, p.text, p.value)
}

// urlStat is a statPrinter for a single value with a url
type urlStat struct {
	text  string
	url   string
	value float64
}

func (p urlStat) write(w io.Writer) (int, error) {
	return fmt.Fprintf(w, p.text, p.url, p.value)
}

// Write stats data in prometheus format
func (stats StatsReply) WritePrometheus(w io.Writer) error {
	statslist := make([]statPrinter, 0, 128)
	for _, stat := range []plainStat{
		{text: "orion_stats_counters_json_requests %g\n", value: stats.Counters.JsonRequests},
		{text: "orion_stats_counters_no_payload_requests %g\n", value: stats.Counters.NoPayloadRequests},
		{text: "orion_stats_counters_invalid_requests %g\n", value: stats.Counters.InvalidRequests},
		{text: "orion_stats_counters_notifications_sent %g\n", value: stats.Counters.NotificationsSent},
		{text: "orion_stats_timing_accumulated{at=\"JsonV1Parse\"} %g\n", value: stats.Timing.Accumulated.JsonV1Parse},
		{text: "orion_stats_timing_accumulated{at=\"JsonV2Parse\"} %g\n", value: stats.Timing.Accumulated.JsonV2Parse},
		{text: "orion_stats_timing_accumulated{at=\"MongoBackend\"} %g\n", value: stats.Timing.Accumulated.MongoBackend},
		{text: "orion_stats_timing_accumulated{at=\"MongoReadWait\"} %g\n", value: stats.Timing.Accumulated.MongoReadWait},
		{text: "orion_stats_timing_accumulated{at=\"MongoWriteWait\"} %g\n", value: stats.Timing.Accumulated.MongoWriteWait},
		{text: "orion_stats_timing_accumulated{at=\"MongoCommandWait\"} %g\n", value: stats.Timing.Accumulated.MongoCommandWait},
		{text: "orion_stats_timing_accumulated{at=\"Render\"} %g\n", value: stats.Timing.Accumulated.Render},
		{text: "orion_stats_timing_accumulated_total %g\n", value: stats.Timing.Accumulated.Total},
		{text: "orion_stats_timing_last{at=\"JsonV2Parse\"} %g\n", value: stats.Timing.Last.JsonV2Parse},
		{text: "orion_stats_timing_last{at=\"MongoBackend\"} %g\n", value: stats.Timing.Last.MongoBackend},
		{text: "orion_stats_timing_last{at=\"MongoReadWait\"} %g\n", value: stats.Timing.Last.MongoReadWait},
		{text: "orion_stats_timing_last{at=\"MongoWriteWait\"} %g\n", value: stats.Timing.Last.MongoWriteWait},
		{text: "orion_stats_timing_last{at=\"Render\"} %g\n", value: stats.Timing.Last.Render},
		{text: "orion_stats_timing_last_total %g\n", value: stats.Timing.Last.Total},
		{text: "orion_stats_semwait_request %g\n", value: stats.Semwait.Request},
		{text: "orion_stats_semwait_db_connection_pool %g\n", value: stats.Semwait.DbConnectionPool},
		{text: "orion_stats_semwait_transaction %g\n", value: stats.Semwait.Transaction},
		{text: "orion_stats_semwait_sub_cache %g\n", value: stats.Semwait.SubCache},
		{text: "orion_stats_semwait_connection_context %g\n", value: stats.Semwait.ConnectionContext},
		{text: "orion_stats_semwait_timestat %g\n", value: stats.Semwait.TimeStat},
		{text: "orion_stats_semwait_metrics %g\n", value: stats.Semwait.Metrics},
		{text: "orion_stats_notif_queue_in %g\n", value: stats.NotifQueue.In},
		{text: "orion_stats_notif_queue_out %g\n", value: stats.NotifQueue.Out},
		{text: "orion_stats_notif_queue_reject %g\n", value: stats.NotifQueue.Reject},
		{text: "orion_stats_notif_queue_sent_ok %g\n", value: stats.NotifQueue.SentOk},
		{text: "orion_stats_notif_queue_sent_error %g\n", value: stats.NotifQueue.SentError},
		{text: "orion_stats_notif_queue_time_in_queue %g\n", value: stats.NotifQueue.AvgTimeInQueue},
		{text: "orion_stats_notif_queue_size %g\n", value: stats.NotifQueue.Size},
		{text: "orion_stats_notif_queue_in %g\n", value: stats.NotifQueue.In},
	} {
		statslist = append(statslist, stat)
	}
	for url, metrics := range stats.Counters.Requests {
		for _, stat := range []urlStat{
			{text: "orion_stats_counters_requests{url=\"%s\",method=\"GET\"} %g\n", url: url, value: metrics.Get},
			{text: "orion_stats_counters_requests{url=\"%s\",method=\"POST\"} %g\n", url: url, value: metrics.Post},
			{text: "orion_stats_counters_requests{url=\"%s\",method=\"PUT\"} %g\n", url: url, value: metrics.Put},
			{text: "orion_stats_counters_requests{url=\"%s\",method=\"PATCH\"} %g\n", url: url, value: metrics.Patch},
			{text: "orion_stats_counters_requests{url=\"%s\",method=\"DELETE\"} %g\n", url: url, value: metrics.Delete},
			{text: "orion_stats_counters_requests{url=\"%s\",method=\"HEAD\"} %g\n", url: url, value: metrics.Head},
		} {
			statslist = append(statslist, stat)
		}
	}
	for url, metrics := range stats.Counters.RequestsLegacy {
		for _, stat := range []urlStat{
			{text: "orion_stats_counters_requests_legacy{url=\"%s\",method=\"GET\"} %g\n", url: url, value: metrics.Get},
			{text: "orion_stats_counters_requests_legacy{url=\"%s\",method=\"POST\"} %g\n", url: url, value: metrics.Post},
			{text: "orion_stats_counters_requests_legacy{url=\"%s\",method=\"PUT\"} %g\n", url: url, value: metrics.Put},
			{text: "orion_stats_counters_requests_legacy{url=\"%s\",method=\"PATCH\"} %g\n", url: url, value: metrics.Patch},
			{text: "orion_stats_counters_requests_legacy{url=\"%s\",method=\"DELETE\"} %g\n", url: url, value: metrics.Delete},
			{text: "orion_stats_counters_requests_legacy{url=\"%s\",method=\"HEAD\"} %g\n", url: url, value: metrics.Head},
		} {
			statslist = append(statslist, stat)
		}
	}
	for _, stat := range statslist {
		if _, err := stat.write(w); err != nil {
			return err
		}
	}
	return nil
}
