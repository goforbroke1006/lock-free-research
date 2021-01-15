package internal

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	OperationCountProcessed    prometheus.Counter
	OperationCountFailed       prometheus.Counter
	OperationDurationProcessed prometheus.Histogram
)

func RegisterNumberCounterMetrics(serviceName, componentName string) {
	OperationCountProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_%s_processed_operations_processed", serviceName, componentName),
		Help: "The number of processed operations ADD",
	})
	OperationCountFailed = promauto.NewCounter(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_%s_processed_operations_failed", serviceName, componentName),
		Help: "The number of failed operations ADD",
	})
	OperationDurationProcessed = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: fmt.Sprintf("%s_%s_processed_operations_duration", serviceName, componentName),
		Help: "The duration of one processed operation ADD",
		Buckets: []float64{
			0, 1, 5, 10, 50, 100,
			1e3, 10e3, 100e3,
			1e6, 5e6, 10e6, 100e6,
			1e9, 10e9},
	})
}
