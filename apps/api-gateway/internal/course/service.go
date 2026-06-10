package course

import (
	"errors"
	"strings"

	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// CreateCourseService coordinates validations for staging new courses.
func CreateCourseService(
	req CreateCourseRequest,
	userID uuid.UUID,
) (*models.Course, error) {

	course := &models.Course{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		CreatedBy:   userID,
	}

	return CreateCourse(course)
}

// UpdateCourseService processes validation structures requesting delta modifications.
func UpdateCourseService(
	req UpdateCourseRequest,
	courseID uuid.UUID,
	userID uuid.UUID,
) (*models.Course, error) {

	existingCourse, err := GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	if existingCourse.CreatedBy != userID {
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

// GetAllCoursesService exposes lookup lists for all validated configurations.
func GetAllCoursesService() ([]models.Course, error) {
	return GetAllCourses()
}

// GetCourseWithPlaylistsService fetches course with its associated playlists.
func GetCourseWithPlaylistsService(
	courseID uuid.UUID,
) (*CourseWithPlaylistResponse, error) {

	return GetCourseWithPlaylists(courseID)
}
