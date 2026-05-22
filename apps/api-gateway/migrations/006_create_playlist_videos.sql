-- =========================================
-- 006_create_playlist_videos.sql
-- =========================================

-- +goose Up

CREATE TABLE playlist_videos (
    playlist_id UUID NOT NULL REFERENCES playlists(playlist_id) ON DELETE CASCADE,
    video_id UUID NOT NULL REFERENCES videos(video_id) ON DELETE CASCADE,
    order_index INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (playlist_id, video_id)
);

CREATE INDEX idx_playlist_videos_order ON playlist_videos(order_index);

-- +goose Down

DROP TABLE playlist_videos;