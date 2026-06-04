-- =========================================
-- 020_alter_submission.sql
-- =========================================

-- +goose Up

ALTER TABLE submissions
RENAME COLUMN id TO submission_id;

ALTER TABLE submissions
ADD COLUMN status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
ADD COLUMN passed_cases INTEGER DEFAULT 0,
ADD COLUMN total_cases INTEGER DEFAULT 0,
ADD COLUMN started_at TIMESTAMPTZ,
ADD COLUMN finished_at TIMESTAMPTZ,
DROP COLUMN verdict;


DROP INDEX IF EXISTS idx_submissions_user_id;
DROP INDEX IF EXISTS idx_submissions_question_id;

CREATE INDEX idx_submissions_user_id
ON submissions(user_id);

CREATE INDEX idx_submissions_question_id
ON submissions(question_id);

CREATE INDEX idx_submissions_status
ON submissions(status);

-- +goose Down

ALTER TABLE submissions
RENAME COLUMN submission_id TO id;

ALTER TABLE submissions
DROP COLUMN status,
DROP COLUMN passed_cases,
DROP COLUMN total_cases,
DROP COLUMN started_at,
DROP COLUMN finished_at,
ADD COLUMN verdict VARCHAR(50);
