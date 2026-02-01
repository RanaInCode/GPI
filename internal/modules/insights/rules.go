package insights

import "github.com/ranaincode/gpi-backend/internal/modules/analytics"

func GrowthInsight(mom float64, trend analytics.Trend) *Insight {
	if mom > 20 && trend == analytics.TrendUp {
		return &Insight{
			Title:   "Strong Organic Growth",
			Summary: "The product shows sustained month-over-month growth, indicating strong market pull.",
			Evidence: []string{
				"MoM growth above 20%",
				"Upward trend over multiple months",
			},
			Confidence: 0.85,
			Category:   "growth",
		}
	}
	return nil
}
