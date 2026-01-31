package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ranaincode/gpi-backend/internal/modules/metrics"
)

type MetricsRepository struct {
	db *pgxpool.Pool
}

func NewMetricRepository(db *pgxpool.Pool) *MetricsRepository {
	return &MetricsRepository{db: db}
}

func (r *MetricsRepository) Insert(ctx context.Context, m metrics.Metric) error {
	query := `
	    INSERT INTO product_metrics (
		   id, product_id, metric_type, value,
		   unit, source, confidence_score,
		   estimation_method, recorded_at
	    )
		VALUES (
		    gen_random_uuid(), $1, $2, $3,
			$4, $5, $6, $7, $8
		)
	`
	_, err := r.db.Exec(ctx, query,
	    m.ProductID,
	    m.MetricType,
	    m.Value,
	    m.Unit,
	    m.Source,
	    m.ConfidenceScore,
	    m.EstimationMethod,
	    m.RecordedAt,
	)
	return err
}