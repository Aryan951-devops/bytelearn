-- =========================================
-- 022_create_submission_results.sql
-- =========================================

-- +goose Up

CREATE TABLE submission_results (

    submission_id UUID NOT NULL
        REFERENCES submissions(submission_id)
        ON DELETE CASCADE,

    testcase_id UUID NOT NULL
        REFERENCES testcases(testcase_id)
        ON DELETE CASCADE,

    actual_output TEXT,

    error_output TEXT,

    is_passed BOOLEAN NOT NULL,

    verdict VARCHAR(100),

    runtime_ms INTEGER,

    memory_kb INTEGER,

    created_at TIMESTAMPTZ DEFAULT NOW(),

    PRIMARY KEY (submission_id, testcase_id)
);

CREATE INDEX idx_submission_results_submission_id
ON submission_results(submission_id);

CREATE INDEX idx_submission_results_testcase_id
ON submission_results(testcase_id);

-- +goose Down

DROP TABLE submission_results;