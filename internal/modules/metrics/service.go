package metrics 

import (
	"context"
	"errors"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo : r}
}

func (s *Service) Ingest(ctx context.Context, m Metric) error {
	if m.MetricType == "" {
		return errors.New("metric_type is required")
	}
	if m.Value < 0 {
		return errors.New("metric value cannot be negative")
	}
    if m.RecordedAt.IsZero() {
		return errors.New("recorded_at is required")
	}

	return s.repo.Insert(ctx, m)
}

func (s *Service) LatestMetric(
	ctx context.Context,
	productID string,
	metricType string,
) (*Metric, error) {
	return s.repo.GetLatest(ctx, productID, metricType)
}

func (s *Service) MetricTrend(
	ctx context.Context,
	productID string,
	metricType string,
	from time.Time,
	to time.Time,
) ([]Metric, error) {
    
	if from.After(to) {
		return nil, errors.New("from time must be before to time: Invalid time range")
	}
	return s.repo.GetTimeSeries(ctx, productID, metricType, from, to)
}