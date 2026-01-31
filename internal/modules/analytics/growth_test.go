package analytics

import (
	"testing"
	"math"
)

func TestMoMGrowth(t *testing.T) {
	result := MoMGrowth(100, 120)
	if result != 20 {
		t.Errorf("expected 20, got %f", result)
	}
}

func TestCAGR(t *testing.T) {
	result := CAGR(100, 200, 2)
	if math.Round(result) != 41 {
		t.Errorf("unexpected CAGR: %f", result)
	}
}
