package analytics

import "math"

func MoMGrowth(prev, curr float64) float64 {
	if prev == 0 {
		return 0
	}
	return ((curr - prev) / prev) * 100
}

func YoYGrowth(prevYear, curr float64) float64 {
	if prevYear == 0 {
		return 0
	}
	return ((curr - prevYear) / prevYear) * 100
}

func CAGR(start, end float64, years float64) float64 {
	if start <=0 || years <=0 {
		return 0
	}
	return (math.Pow(end/start, 1/years) -1) * 100
}