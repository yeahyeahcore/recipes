CREATE TABLE IF NOT EXISTS (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    photo_url TEXT,
    ingredients TEXT[],
    steps JSONB[],
    total_steps_time INTERVAL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
    created_at TIMESTAMP WITHOUT TIME ZONE,
    deleted BOOLEAN DEFAULT false
);