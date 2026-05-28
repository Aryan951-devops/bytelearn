-- =========================================
-- 018_alter_column_course.sql
-- =========================================

-- +goose Up

ALTER TABLE courses
ADD COLUMN created_by UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE;


-- +goose Down

ALTER TABLE courses
DROP COLUMN created_by;
