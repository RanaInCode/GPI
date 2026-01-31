CREATE TABLE product_metrics (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES products(id),
    metric_type TEXT NOT NULL,
    value NUMERIC NOT NULL,
    unit TEXT,
    source TEXT,
    confidence_score NUMERIC CHECK (confidence_score >= 0 AND confidence_score <= 1),
    estimation_method TEXT,
    recorded_at DATE NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE OR REPLACE FUNCTION prevent_metric_update()
RETURNS trigger AS $$
BEGIN
  RAISE EXCEPTION 'Metrics are append-only';
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER no_update_product_metrics
BEFORE UPDATE ON product_metrics
FOR EACH ROW
EXECUTE FUNCTION prevent_metric_update();
