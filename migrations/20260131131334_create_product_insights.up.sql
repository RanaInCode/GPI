CREATE TABLE product_insights (
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES products(id),
    insight_type TEXT NOT NULL,
    content JSONB NOT NULL,
    generated_by TEXT,
    confidence_score NUMERIC,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
