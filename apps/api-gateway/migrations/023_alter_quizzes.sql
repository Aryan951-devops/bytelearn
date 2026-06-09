-- =========================================
-- 023_alter_quizzes.sql
-- =========================================

-- +goose Up

ALTER TABLE quizzes
DROP COLUMN questions,
ADD COLUMN duration_minutes INTEGER DEFAULT 50 CHECK (duration_minutes > 0);

ALTER TABLE quizzes
RENAME COLUMN id TO quiz_id;

-- +goose Down

ALTER TABLE quizzes
ADD COLUMN questions JSONB NOT NULL,
DROP COLUMN duration_minutes;

ALTER TABLE quizzes
RENAME COLUMN quiz_id TO id;