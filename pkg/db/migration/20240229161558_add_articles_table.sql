-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- Table: articles
CREATE TABLE IF NOT EXISTS articles (
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
	tags JSONB DEFAULT '[]'::JSONB,
	created_at timestamptz not null default clock_timestamp(),
  	updated_at timestamptz not null default clock_timestamp()
);
-- Indexes
CREATE UNIQUE INDEX articles_title_uniq_idx ON articles(LOWER(title));

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- Table: articles
DROP TABLE IF EXISTS articles;
-- Indexes
DROP INDEX IF EXISTS articles_title_uniq_idx;