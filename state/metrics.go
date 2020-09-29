package state

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

const (
	// MetricsSubsystem is a subsystem shared by all metrics exposed by this
	// package.
	MetricsSubsystem = "state"
)

// Metrics contains metrics exposed by this package.
type Metrics struct {
	// Time between BeginBlock and EndBlock.
	BlockProcessingTime metrics.Histogram

	// Time between BeginBlock and EndBlock, not as a histogram
	BlockProcessingTimeSingle metrics.Gauge

	// Time to validate block in ms
	ValidationTime metrics.Gauge

	// Time to verify signatures in ms
	VerifyCommitTime metrics.Gauge

	// Time to validate validator updates
	ValidateUpdateTime metrics.Gauge

	// Time to commit state and update mempool
	StateCommitMempoolTime metrics.Gauge

	// Total time to execute block
	TotalBlockExecTime metrics.Gauge
}

// PrometheusMetrics returns Metrics build using Prometheus client library.
// Optionally, labels can be provided along with their values ("foo",
// "fooValue").
func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		BlockProcessingTime: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_processing_time",
			Help:      "Time between BeginBlock and EndBlock in ms.",
			Buckets:   stdprometheus.LinearBuckets(1, 10, 10),
		}, labels).With(labelsAndValues...),
		BlockProcessingTimeSingle: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "block_processing_time_single",
			Help:      "Time to execute the current block in milliseconds",
		}, labels).With(labelsAndValues...),
		ValidationTime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validation_time",
			Help:      "Time to validate block in ms.",
		}, labels).With(labelsAndValues...),
		VerifyCommitTime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "verify_commit_time",
			Help:      "Time to verify signatures in ms.",
		}, labels).With(labelsAndValues...),
		ValidateUpdateTime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "validate_update_time",
			Help:      "Time to validate validator updates in ms.",
		}, labels).With(labelsAndValues...),
		StateCommitMempoolTime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "state_commit_mempool_update_time",
			Help:      "Time to commit state and update mempool in ms.",
		}, labels).With(labelsAndValues...),
		TotalBlockExecTime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "total_block_exec_time",
			Help:      "Total time to execute block in ms.",
		}, labels).With(labelsAndValues...),
	}
}

// NopMetrics returns no-op Metrics.
func NopMetrics() *Metrics {
	return &Metrics{
		BlockProcessingTime:       discard.NewHistogram(),
		BlockProcessingTimeSingle: discard.NewGauge(),
		ValidationTime:            discard.NewGauge(),
		VerifyCommitTime:          discard.NewGauge(),
		ValidateUpdateTime:        discard.NewGauge(),
		StateCommitMempoolTime:    discard.NewGauge(),
		TotalBlockExecTime:        discard.NewGauge(),
	}
}
