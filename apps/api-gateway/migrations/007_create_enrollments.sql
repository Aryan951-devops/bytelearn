-- =========================================
-- 007_create_enrollments.sql
-- =========================================

-- +goose Up

CREATE TABLE enrollments (
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    course_id UUID NOT NULL REFERENCES courses(course_id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id, course_id)
);

-- +goose Down

DROP TABLE enrollments;