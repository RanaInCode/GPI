package analytics

import (
	"testing"
)

func TestTrendUp(t *testing.T) {
	values := []float64{10, 20, 30, 40}
	if ClassifyTrend(values) != TrendUp {
		t.Fail()
	}
}

func TestTrendVolatile(t *testing.T) {
	values := []float64{10, 30, 15, 40}
	if ClassifyTrend(values) != TrendVolatile {
		t.Fail()
	}
}
