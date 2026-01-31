package database

import (
	"context"
	"time"

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

func (r *MetricsRepository) GetLatest(
	ctx context.Context,
	productID string,
	metricType string,
) (*metrics.Metric, error) {

	 query := `
	    SELECT product_id, metric_type, value, unit,
		       source, confidence_score, estimation_method, recorded_at
		FROM product_metrics
		WHERE product_id = $1 AND metric_type = $2
		ORDER BY recorded_at DESC
		LIMIT 1
	`
	row := r.db.QueryRow(ctx, query, productID, metricType)

	var m metrics.Metric
	err := row.Scan(
	    &m.ProductID,
	    &m.MetricType,
	    &m.Value,
	    &m.Unit,
	    &m.Source,
	    &m.ConfidenceScore,
	    &m.EstimationMethod,
	    &m.RecordedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *MetricsRepository) GetTimeSeries(
	ctx context.Context,
	productID string,
	metricType string,
	from time.Time,
	to time.Time,
) ([]metrics.Metric, error) {

	query := `
		SELECT product_id, metric_type, value, unit,
		       source, confidence_score, estimation_method, recorded_at
		FROM product_metrics
		WHERE product_id = $1
		  AND metric_type = $2
		  AND recorded_at BETWEEN $3 AND $4
		ORDER BY recorded_at ASC
	`

	rows, err := r.db.Query(ctx, query, productID, metricType, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []metrics.Metric
	for rows.Next() {
		var m metrics.Metric
		err := rows.Scan(
			&m.ProductID,
			&m.MetricType,
			&m.Value,
			&m.Unit,
			&m.Source,
			&m.ConfidenceScore,
			&m.EstimationMethod,
			&m.RecordedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, m)
	}

	return result, nil
}
