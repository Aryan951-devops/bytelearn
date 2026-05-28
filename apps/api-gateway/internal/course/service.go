package course

import (
	"errors"
	"strings"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/google/uuid"
)

func CreateCourseService(
	req CreateCourseRequest,
	user_id uuid.UUID,
) (*models.Course, error) {

	course := &models.Course{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		CreatedBy:   user_id,
	}

	return CreateCourse(course)
}

func UpdateCourseService(
	req UpdateCourseRequest,
	course_id uuid.UUID,
	user_id uuid.UUID,
) (*models.Course, error) {

	existingCourse, err := GetCourseById(course_id)
	if err != nil {
		return nil, err
	}

	if existingCourse.CreatedBy != user_id {
		return nil, errors.New("you are not authorized to update this playlist")
	}

	if req.Title != nil && strings.TrimSpace(*req.Title) != "" {
		existingCourse.Title = *req.Title
	}

	if req.Description != nil && strings.TrimSpace(*req.Description) != "" {
		existingCourse.Description = req.Description
	}

	if req.Category != nil && strings.TrimSpace(*req.Category) != "" {
		existingCourse.Category = req.Category
	}

	return UpdateCourse(existingCourse)
}

func GetAllCoursesService() ([]models.Course, error) {
	return GetAllCourses()
}

func GetCourseWithPlaylistsService(
	courseID uuid.UUID,
) (*CourseWithPlaylistResponse, error) {

	return GetCourseWithPlaylists(courseID)
}
