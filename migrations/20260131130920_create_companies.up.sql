CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    website TEXT,
    founded_year INT,
    headquarters_country TEXT,
    total_funding_usd BIGINT,
    funding_stage TEXT,
    created_at TIMESTAMP DEFAULT now()
);