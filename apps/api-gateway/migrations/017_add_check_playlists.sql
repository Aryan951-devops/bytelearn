-- =========================================
-- 017_add_check_playlists.sql
-- =========================================

-- +goose Up

ALTER TABLE playlists 
ADD CONSTRAINT chk_playlist_type CHECK (type IN ('course', 'user'));


-- +goose Down

ALTER TABLE playlists 
DROP CONSTRAINT chk_playlist_type;
