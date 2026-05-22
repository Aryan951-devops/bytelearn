-- =========================================
-- 010_create_comment_likes.sql
-- =========================================

-- +goose Up

CREATE TABLE comment_likes (
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    comment_id UUID NOT NULL REFERENCES comments(comment_id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id, comment_id)
);

-- +goose Down

DROP TABLE comment_likes;