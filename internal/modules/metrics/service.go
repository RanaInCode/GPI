package metrics 

import (
	"context"
	"errors"
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