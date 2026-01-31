package metrics

import "time"

type Metric struct {
	ID               string
	ProductID        string
	MetricType       string
	Value            float64
	Unit             string
	Source           string
	ConfidenceScore  float64
	EstimationMethod string
	RecordedAt       time.Time
}