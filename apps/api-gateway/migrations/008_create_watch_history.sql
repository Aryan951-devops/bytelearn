-- =========================================
-- 008_create_watch_history.sql
-- =========================================

-- +goose Up

CREATE TABLE watch_history (
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    video_id UUID NOT NULL REFERENCES videos(video_id) ON DELETE CASCADE,
    resume_time INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id, video_id)
);

-- +goose Down

DROP TABLE watch_history;