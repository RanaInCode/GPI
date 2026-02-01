package insights

import (
	"testing"
	"context"
	
	"github.com/ranaincode/gpi-backend/internal/modules/analytics"
)

func TestGrowthInsightGenerated(t *testing.T) {
	mockAnalytics := &mockAnalytics{
		mom:   25,
		trend: analytics.TrendUp,
	}

	svc := NewService(mockAnalytics)

	insights, _ := svc.Generate(context.Background(), "p1")

	if len(insights) == 0 {
		t.Fail()
	}
}
