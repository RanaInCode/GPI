package metrics

import "context"

type Repository interface {
	Insert(ctx context.Context, m Metric) error
}