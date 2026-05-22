-- =========================================
-- 014_create_quizzes.sql
-- =========================================

-- +goose Up

CREATE TABLE quizzes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    playlist_id UUID NOT NULL REFERENCES playlists(playlist_id) ON DELETE CASCADE,
    questions JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_quizzes_playlist_id
ON quizzes(playlist_id);

CREATE INDEX idx_quizzes_questions
ON quizzes USING GIN(questions);

-- +goose Down

DROP TABLE quizzes;