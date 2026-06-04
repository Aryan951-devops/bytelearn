-- =========================================
-- 021_create_testcases.sql
-- =========================================

-- +goose Up

CREATE TABLE testcases (
    testcase_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    question_id UUID NOT NULL
        REFERENCES coding_questions(question_id)
        ON DELETE CASCADE,

    input TEXT NOT NULL,

    expected_output TEXT NOT NULL,

    is_hidden BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_testcases_question_id
ON testcases(question_id);

CREATE INDEX idx_testcases_question_hidden
ON testcases(question_id, is_hidden);

-- +goose Down

DROP TABLE testcases;