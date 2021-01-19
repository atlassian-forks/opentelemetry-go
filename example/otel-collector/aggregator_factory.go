package main

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/number"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
)

type histogramAggregatorFactory struct {
	metric.AggregatorFactory

	boundaries []float64
	descriptor metric.Descriptor
}

func NewHistogramAggregatorFactory(boundaries []float64) *histogramAggregatorFactory {
	agg := &histogramAggregatorFactory{
		boundaries: boundaries,
		descriptor: metric.NewDescriptor("dummy", metric.ValueObserverInstrumentKind, number.Float64Kind),
	}
	return agg
}

// NewInstance returns one histogram aggregator with the defined explicit boundaries and descriptor
func (h *histogramAggregatorFactory) NewInstance() metric.Aggregator {
	return &histogram.New(1, &h.descriptor, histogram.WithExplicitBoundaries(h.boundaries))[0]
}
