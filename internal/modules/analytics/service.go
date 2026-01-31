package analytics

import (
	"context"
	"time"

	"github.com/ranaincode/gpi-backend/internal/modules/metrics"
)

type MetricsReader interface {
	GetTimeSeries(
		ctx context.Context,
		productID string,
		metricType string,
		from time.Time,
		to time.Time,
	) ([]metrics.Metric, error)
}

type Service struct {
	reader MetricsReader
}

func NewService(reader MetricsReader) *Service {
	return &Service{reader: reader}
}

func (s *Service) MoM(
	ctx context.Context,
	productID string,
	metricType string,
) (float64, error) {

	now := time.Now()
	from := now.AddDate(0, -2, 0)

	data, err := s.reader.GetTimeSeries(ctx, productID, metricType, from, now)
	if err != nil || len(data) < 2 {
		return 0, err
	}

	prev := data[len(data)-2].Value
	curr := data[len(data)-1].Value

	return MoMGrowth(prev, curr), nil
}

