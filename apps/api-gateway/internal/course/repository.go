package course

import (
	"context"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/database"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/google/uuid"
)

func CreateCourse(course *models.Course) (*models.Course, error) {

	query := `
	INSERT INTO courses
	(title, description, category, created_by, created_at, updated_at)
	VALUES ($1,$2,$3,$4,NOW(),NOW())
	RETURNING
	course_id, title, description, category, created_by,
	created_at, updated_at
	`

	var new_course models.Course

	err := database.DB.QueryRow(
		context.Background(),
		query,
		course.Title,
		course.Description,
		course.Category,
		course.CreatedBy,
	).Scan(
		&new_course.ID,
		&new_course.Title,
		&new_course.Description,
		&new_course.Category,
		&new_course.CreatedBy,
		&new_course.CreatedAt,
		&new_course.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &new_course, nil
}

func UpdateCourse(course *models.Course) (*models.Course, error) {

	query := `
	UPDATE courses
	SET
		title = $2,
		description = $3,
		category = $4,
		updated_at = NOW()
	WHERE course_id = $1
	RETURNING 
		course_id, title, description, category,
		created_by, created_at, updated_at
	`

	var updatedCourse models.Course

	err := database.DB.QueryRow(
		context.Background(),
		query,
		course.ID,
		course.Title,
		course.Description,
		course.Category,
	).Scan(
		&updatedCourse.ID,
		&updatedCourse.Title,
		&updatedCourse.Description,
		&updatedCourse.Category,
		&updatedCourse.CreatedBy,
		&updatedCourse.CreatedAt,
		&updatedCourse.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &updatedCourse, nil
}

func GetAllCourses() ([]models.Course, error) {

	query := `
	SELECT
		course_id,
		title,
		description,
		category,
		created_by,
		created_at,
		updated_at
	FROM courses
	ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []models.Course{}

	for rows.Next() {

		var course models.Course

		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.Category,
			&course.CreatedBy,
			&course.CreatedAt,
			&course.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func GetCourseById(course_id uuid.UUID) (*models.Course, error) {
	query := `
	SELECT
		course_id, title, description, 
		category, created_at, updated_at
	FROM courses
	WHERE c.course_id = $1
	`

	var course models.Course

	err := database.DB.QueryRow(
		context.Background(),
		query,
		course_id,
	).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.Category,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func GetCourseWithPlaylists(course_id uuid.UUID) (*CourseWithPlaylistResponse, error) {

	query := `
	SELECT
		course_id, title, description, 
		category, created_at, updated_at
	FROM courses
	WHERE course_id = $1
	`

	var response CourseWithPlaylistResponse
	response.Playlists = []PlaylistPreview{}
	response.Educators = []EducatorDetail{}

	err := database.DB.QueryRow(
		context.Background(),
		query,
		course_id,
	).Scan(
		&response.ID,
		&response.Title,
		&response.Description,
		&response.Category,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	playlistQuery := `
		SELECT 
			p.playlist_id, p.title, p.description,
			u.user_id, u.username, u.name, u.profile_pic_url
		FROM playlists p INNER JOIN users u
		ON p.user_id = u.user_id
		WHERE course_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		playlistQuery,
		course_id,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var playlist PlaylistPreview
		var educator EducatorDetail
		err := rows.Scan(
			&playlist.ID,
			&playlist.Title,
			&playlist.Description,
			&educator.ID,
			&educator.Username,
			&educator.Name,
			&educator.ProfilePic_Url,
		)

		if err != nil {
			return nil, err
		}

		response.Playlists = append(response.Playlists, playlist)
		response.Educators = append(response.Educators, educator)
	}

	return &response, nil
}
