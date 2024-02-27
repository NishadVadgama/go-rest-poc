-- Articles table schema
CREATE TABLE IF NOT EXISTS articles (
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
	tags JSONB DEFAULT '[]'::JSONB
)
