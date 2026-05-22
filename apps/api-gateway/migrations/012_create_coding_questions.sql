-- =========================================
-- 012_create_coding_questions.sql
-- =========================================

-- +goose Up

CREATE TABLE coding_questions (
    question_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    metadata JSONB NOT NULL,
    contest_id UUID NOT NULL REFERENCES coding_contests(contest_id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_coding_questions_contest_id
ON coding_questions(contest_id);

CREATE INDEX idx_coding_questions_metadata
ON coding_questions USING GIN(metadata);

-- +goose Down

DROP TABLE coding_questions;