-- =========================================
-- 009_create_video_likes.sql
-- =========================================

-- +goose Up

CREATE TABLE video_likes (
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    video_id UUID NOT NULL REFERENCES videos(video_id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id, video_id)
);

-- +goose Down

DROP TABLE video_likes;