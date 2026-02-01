package insights

import (
	"context"

	"github.com/ranaincode/gpi-backend/internal/modules/analytics"
)

type AnalyticsReader interface {
	GetMoMGrowth(ctx context.Context, productID, metric string) (float64, error)
	GetTrend(ctx context.Context, productID, metric string) (analytics.Trend, error)
}

type Service struct {
	analytics AnalyticsReader
}

func NewService(analytics AnalyticsReader) *Service {
	return &Service{analytics: analytics}
}

func (s *Service) Generate(
	ctx context.Context,
	productID string,
) ([]Insight, error) {

	var results []Insight

	mom, _ := s.analytics.GetMoMGrowth(ctx, productID, "revenue")
	trend, _ := s.analytics.GetTrend(ctx, productID, "revenue")

	if insight := GrowthInsight(mom, trend); insight != nil {
		results = append(results, *insight)
	}

	return results, nil
}

