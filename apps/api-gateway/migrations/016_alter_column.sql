-- =========================================
-- 016_alter_column.sql
-- =========================================

-- +goose Up

ALTER TABLE users
RENAME column profile_pic TO profile_pic_url;

ALTER TABLE users
ADD COLUMN profile_pic_public_id TEXT;

ALTER TABLE videos 
ADD COLUMN videofile_public_id TEXT NOT NULL;

ALTER TABLE videos
ADD COLUMN thumbnail_public_id TEXT;

-- +goose Down

ALTER TABLE users
DROP COLUMN profile_pic_public_id;

ALTER TABLE videos
DROP COLUMN videofile_public_id;

ALTER TABLE videos
DROP COLUMN thumbnail_public_id;