-- =========================================
-- 011_create_coding_contests.sql
-- =========================================

-- +goose Up

CREATE TABLE coding_contests (
    contest_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    playlist_id UUID REFERENCES playlists(playlist_id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down

DROP TABLE coding_contests;