-- =========================================
-- 025_create_quiz_attempts.sql
-- =========================================

-- +goose Up
CREATE TYPE quiz_attempt_status AS ENUM ('in_progress', 'submitted', 'expired');

CREATE TABLE quiz_attempts (
    attempt_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    quiz_id UUID NOT NULL REFERENCES quizzes(quiz_id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    score INTEGER DEFAULT 0,
    total_marks INTEGER DEFAULT 0,
    status quiz_attempt_status DEFAULT 'in_progress'::quiz_attempt_status NOT NULL,
    submitted_answers JSONB DEFAULT '[]'::JSONB NOT NULL,
    started_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    submitted_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_quiz_attempts_quiz_id ON quiz_attempts(quiz_id);
CREATE INDEX IF NOT EXISTS idx_quiz_attempts_user_id ON quiz_attempts(user_id);

-- +goose Down

DROP INDEX IF EXISTS idx_quiz_attempts_user_id;
DROP INDEX IF EXISTS idx_quiz_attempts_quiz_id;

DROP TABLE quiz_attempts;

DROP TYPE IF EXISTS quiz_attempt_status;