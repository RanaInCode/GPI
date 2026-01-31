package analytics

import (
	"context"
	"time"
	"testing"

	"github.com/ranaincode/gpi-backend/internal/modules/metrics"
)

type mockReader struct {
	data []metrics.Metric
}

func (m *mockReader) GetTimeSeries(
	ctx context.Context,
	productID string,
	metricType string,
	from time.Time,
	to time.Time,
) ([]metrics.Metric, error) {
	return m.data, nil
}

func TestMoMService(t *testing.T) {
	mock := &mockReader{
		data: []metrics.Metric{
			{Value: 100},
			{Value: 120},
		},
	}

	svc := NewService(mock)

	result, err := svc.MoM(context.Background(), "p1", "revenue")
	if err != nil || result != 20 {
		t.Fail()
	}
}

