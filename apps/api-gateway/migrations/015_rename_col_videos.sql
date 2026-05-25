-- =========================================
-- 015_rename_col_videos.sql
-- =========================================

-- +goose Up

ALTER TABLE videos RENAME column video_file_url TO videofile_url;

-- +goose Down