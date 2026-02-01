package insights

func ComputeConfidence(dataPoints int, volatility bool) float64 {
	score := 0.5

	if dataPoints > 12 {
		score += 0.3
	}
	if !volatility {
		score += 0.2
	}

	if score > 1 {
		score = 1
	}
	return score
}
