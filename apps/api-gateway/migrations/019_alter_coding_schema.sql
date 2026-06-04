-- =========================================
-- 019_alter_coding_schema.sql
-- =========================================

-- +goose Up

ALTER TABLE coding_contests
RENAME TO coding_practice;

ALTER TABLE coding_practice
    -- 1. Drop the existing foreign key constraint
    DROP CONSTRAINT coding_contests_playlist_id_fkey,

    -- 2. Add the new constraint with ON DELETE CASCADE
    ADD CONSTRAINT coding_contests_playlist_id_fkey 
        FOREIGN KEY (playlist_id) 
        REFERENCES playlists(playlist_id) 
        ON DELETE CASCADE;

ALTER TABLE coding_questions
DROP COLUMN metadata;

ALTER TABLE coding_questions
ADD COLUMN title VARCHAR(255) NOT NULL,
ADD COLUMN difficulty VARCHAR(20) DEFAULT 'EASY',
ADD COLUMN statement TEXT NOT NULL,
ADD COLUMN constraints TEXT,
ADD COLUMN input_format TEXT,
ADD COLUMN output_format TEXT,
ADD COLUMN time_limit_ms INTEGER NOT NULL DEFAULT 1000,
ADD COLUMN memory_limit_mb INTEGER NOT NULL DEFAULT 256;


-- +goose Down

ALTER TABLE coding_questions
DROP COLUMN title,
DROP COLUMN difficulty,
DROP COLUMN statement,
DROP COLUMN constraints,
DROP COLUMN input_format,
DROP COLUMN output_format,
DROP COLUMN time_limit_ms,
DROP COLUMN memory_limit_mb;

ALTER TABLE coding_questions
ADD COLUMN metadata JSONB NOT NULL;

ALTER TABLE coding_practice
RENAME TO coding_contests;
