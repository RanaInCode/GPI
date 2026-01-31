package metrics

import (
    "context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, m Metric) error

	GetLatest(
		ctx context.Context,
		productID string,
		metricType string,
	) (*Metric, error)

	GetTimeSeries(
		ctx context.Context,
		productID string,
		metricType string,
		from time.Time,
		to time.Time,
	) ([]Metric, error)
}