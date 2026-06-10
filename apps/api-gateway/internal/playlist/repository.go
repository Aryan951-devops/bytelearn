package playlist

import (
	"context"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// CreatePlaylist inserts a new playlist entry.
func CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error) {
	query := `
	INSERT INTO playlists
		(type, title, description, 
		user_id, course_id, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5,NOW(),NOW())
	RETURNING playlist_id, type, title,
	description, user_id, course_id, 
	created_at, updated_at
	`

	var newPlaylist models.Playlist
	err := database.DB.QueryRow(
		context.Background(),
		query,
		playlist.Type,
		playlist.Title,
		playlist.Description,
		playlist.UserID,
		playlist.CourseID,
	).Scan(
		&newPlaylist.ID,
		&newPlaylist.Type,
		&newPlaylist.Title,
		&newPlaylist.Description,
		&newPlaylist.UserID,
		&newPlaylist.CourseID,
		&newPlaylist.CreatedAt,
		&newPlaylist.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newPlaylist, err
}

// GetUserPlaylists returns all playlists matching the provided userID.
func GetUserPlaylists(userID uuid.UUID) ([]models.Playlist, error) {
	query := `
	SELECT playlist_id, type, title, description, 
	user_id, course_id, created_at, updated_at
	FROM playlists
	WHERE user_id=$1
	ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	playlists := []models.Playlist{}

	for rows.Next() {
		var p models.Playlist

		err := rows.Scan(
			&p.ID,
			&p.Type,
			&p.Title,
			&p.Description,
			&p.UserID,
			&p.CourseID,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		playlists = append(playlists, p)
	}

	return playlists, nil
}

// AddVideoToPlaylist links a video file to a target playlist.
func AddVideoToPlaylist(
	playlistID uuid.UUID,
	videoID uuid.UUID,
) error {

	query := `
	INSERT INTO playlist_videos
		(playlist_id, video_id, order_index, created_at)
	VALUES (
		$1,
		$2,
		COALESCE(
			(SELECT MAX(order_index)+1 FROM playlist_videos WHERE playlist_id=$1),
			1
		),
		NOW()
	)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		playlistID,
		videoID,
	)
	return err
}

// DeleteVideoFromPlaylist removes a video entry from a specific playlist.
func DeleteVideoFromPlaylist(
	playlistID uuid.UUID,
	videoID uuid.UUID,
) error {

	query := `
	DELETE FROM playlist_videos
	WHERE playlist_id=$1 AND video_id=$2
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		playlistID,
		videoID,
	)
	return err
}

// UpdatePlaylist modifies attributes of an existing playlist.
func UpdatePlaylist(playlist *models.Playlist) (*models.Playlist, error) {

	query := `
    UPDATE playlists
    SET
        title = $2,
        description = $3,
        updated_at = NOW()
    WHERE playlist_id = $1
    RETURNING playlist_id, type, title, description, user_id, course_id, created_at, updated_at
    `

	var updatedPlaylist models.Playlist

	err := database.DB.QueryRow(
		context.Background(),
		query,
		playlist.ID,
		playlist.Title,
		playlist.Description,
	).Scan(
		&updatedPlaylist.ID,
		&updatedPlaylist.Type,
		&updatedPlaylist.Title,
		&updatedPlaylist.Description,
		&updatedPlaylist.UserID,
		&updatedPlaylist.CourseID,
		&updatedPlaylist.CreatedAt,
		&updatedPlaylist.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &updatedPlaylist, nil
}

// GetUserPlaylistWithVideos retrieves playlist details along with its videos.
func GetUserPlaylistWithVideos(playlistID uuid.UUID,
) (*PlaylistResponse, error) {

	playlistQuery := `
		SELECT playlist_id, type, title, description, 
		user_id, course_id, created_at, updated_at
		FROM playlists
		WHERE playlist_id = $1
	`

	var response PlaylistResponse
	response.Videos = []VideoMetadata{}

	err := database.DB.QueryRow(
		context.Background(),
		playlistQuery,
		playlistID,
	).Scan(
		&response.ID,
		&response.Type,
		&response.Title,
		&response.Description,
		&response.UserID,
		&response.CourseID,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	videosQuery := `
		SELECT 
			v.video_id,
			v.title,
			v.thumbnail_url,
			v.views
		FROM playlist_videos pv
		INNER JOIN videos v ON v.video_id = pv.video_id
		WHERE pv.playlist_id = $1
		ORDER BY pv.order_index ASC
	`

	rows, err := database.DB.Query(
		context.Background(),
		videosQuery,
		playlistID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video VideoMetadata
		err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.ThumbnailURL,
			&video.Views,
		)
		if err != nil {
			return nil, err
		}
		response.Videos = append(response.Videos, video)
	}

	return &response, nil
}

// GetCoursePlaylist fetches structured playlist details tailored for structural courses.
func GetCoursePlaylist(playlistID uuid.UUID,
) (*CoursePlaylistResponse, error) {

	playlistQuery := `
		SELECT p.playlist_id, p.type, p.title, p.description, 
		p.user_id, p.course_id, p.created_at, p.updated_at,
		u.user_id, u.username, u.name, u.profile_pic_url
		FROM playlists as p INNER JOIN users u
		ON p.user_id = u.user_id
		WHERE p.playlist_id = $1
	`

	var response CoursePlaylistResponse
	response.Videos = []VideoMetadata{}

	err := database.DB.QueryRow(
		context.Background(),
		playlistQuery,
		playlistID,
	).Scan(
		&response.ID,
		&response.Type,
		&response.Title,
		&response.Description,
		&response.UserID,
		&response.CourseID,
		&response.CreatedAt,
		&response.UpdatedAt,
		&response.EducatorID,
		&response.EducatorUsername,
		&response.EducatorName,
		&response.EducatorProfilePic,
	)
	if err != nil {
		return nil, err
	}
	videosQuery := `
		SELECT 
			v.video_id,
			v.title,
			v.thumbnail_url,
			v.views
		FROM playlist_videos pv
		INNER JOIN videos v ON v.video_id = pv.video_id
		WHERE pv.playlist_id = $1
		ORDER BY pv.order_index ASC
	`

	rows, err := database.DB.Query(
		context.Background(),
		videosQuery,
		playlistID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var video VideoMetadata
		err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.ThumbnailURL,
			&video.Views,
		)
		if err != nil {
			return nil, err
		}
		response.Videos = append(response.Videos, video)
	}

	return &response, nil
}

// VerifyPlaylistOwnership checks if the user possesses ownership to the playlist.
func VerifyPlaylistOwnership(
	playlistID uuid.UUID,
	userID uuid.UUID,
) error {

	query := `
	SELECT playlist_id
	FROM playlists
	WHERE playlist_id=$1 AND user_id=$2
	`

	var id uuid.UUID

	return database.DB.QueryRow(
		context.Background(),
		query,
		playlistID,
		userID,
	).Scan(&id)
}

// GetPlaylistByID retrieves a standard playlist object by its ID.
func GetPlaylistByID(playlistID uuid.UUID) (*models.Playlist, error) {
	query := `
		SELECT 
		playlist_id, type, title, description, 
		user_id, course_id, created_at, updated_at
		FROM playlists 
		WHERE playlist_id = $1
	`

	var playlist models.Playlist

	err := database.DB.QueryRow(
		context.Background(),
		query,
		playlistID,
	).Scan(
		&playlist.ID,
		&playlist.Type,
		&playlist.Title,
		&playlist.Description,
		&playlist.UserID,
		&playlist.CourseID,
		&playlist.CreatedAt,
		&playlist.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}
