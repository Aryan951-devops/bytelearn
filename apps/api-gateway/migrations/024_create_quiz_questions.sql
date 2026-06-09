-- =========================================
-- 024_create_quiz_questions.sql
-- =========================================

-- +goose Up

CREATE TYPE question_type AS ENUM ('mcq', 'multiple', 'one_word', 'true_false');

CREATE TABLE quiz_questions (
    question_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    quiz_id UUID NOT NULL REFERENCES quizzes(quiz_id) ON DELETE CASCADE,
    type question_type NOT NULL,
    question TEXT NOT NULL,
    options TEXT[] DEFAULT '{}'::TEXT[], -- Stored natively as an array of strings
    correct_options INT[] DEFAULT '{}'::INT[], -- Stored natively as an array of integers
    correct_answer TEXT,
    marks INT NOT NULL,
    negative_marks INT DEFAULT 0 NOT NULL,
    explanation TEXT
);

CREATE INDEX IF NOT EXISTS idx_quiz_questions_quiz_id ON quiz_questions(quiz_id);

-- +goose Down



DROP INDEX IF EXISTS idx_quiz_questions_quiz_id;

DROP TABLE quiz_questions;

DROP TYPE IF EXISTS question_type;