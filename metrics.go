package main

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// convert CamelCase to snake_case
func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

type Metrics map[string]float64

// MetricsReply replicates the format of the reply from the metrics endpoint
type MetricsReply struct {
	Services map[string]struct {
		Subservices map[string]Metrics `json:"subservs"`
	} `json:"services"`
}

// write metrics data in Prometheus format
func (reply MetricsReply) WritePrometheus(w io.Writer) error {
	snakeCased := make(map[string]string)
	for serviceName, service := range reply.Services {
		for subserviceName, subservice := range service.Subservices {
			for metricName, metricValue := range subservice {
				// memorize conversions to snake case
				sn, ok := snakeCased[metricName]
				if !ok {
					sn = toSnakeCase(metricName)
					snakeCased[metricName] = sn
				}
				if _, err := fmt.Fprintf(w, "orion_%s {service=\"%s\", subservice=\"%s\"} %f\n", sn, serviceName, subserviceName, metricValue); err != nil {
					return fmt.Errorf("failed to write metric to output: %w", err)
				}
			}
		}
	}
	return nil
}
