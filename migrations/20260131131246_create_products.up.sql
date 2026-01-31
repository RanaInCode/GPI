CREATE TABLE products (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    company_id UUID REFERENCES companies(id),
    category TEXT NOT NULL,
    sub_category TEXT,
    platform TEXT,
    country_origin TEXT,
    launch_year INT,
    status TEXT DEFAULT 'active',
    description TEXT,
    created_at TIMESTAMP DEFAULT now()
);
