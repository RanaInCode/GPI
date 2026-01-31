package analytics

type Trend string

const (
	TrendUp       Trend = "up"
	TrendDown     Trend = "down"
	TrendFlat     Trend = "flat"
	TrendVolatile Trend = "volatile"
)

func ClassifyTrend(values []float64) Trend {
	if len(values) < 3 {
		return TrendFlat
	}
	up, down := 0, 0
	
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			up++
		} else if values[i] < values[i-1] {
			down++
		}
	}

	if up > down*2 {
		return TrendUp
	}
	if down > up*2 {
		return TrendDown
	}
	if up > 0 && down > 0 {
		return TrendVolatile
	}

	return TrendFlat
}