package course

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCourseHandler exposes target frameworks for creation routines.
func CreateCourseHandler(c *gin.Context) {
	var req CreateCourseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"You are not authorized to create course",
			nil,
		))
		return
	}

	course, err := CreateCourseService(req, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"course created successfully",
		gin.H{
			"course": course,
		},
	))
}

// UpdateCourseHandler executes targeted operational updates.
func UpdateCourseHandler(c *gin.Context) {
	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid courseId",
			nil,
		))
		return
	}

	var req UpdateCourseRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	if user.Role != "admin" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"You are not authorized to update course",
			nil,
		))
		return
	}

	updateCourse, err := UpdateCourseService(req, courseID, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"course updated successfully",
		gin.H{
			"course": updateCourse,
		},
	))
}

// GetAllCoursesHandler fetches all courses.
func GetAllCoursesHandler(c *gin.Context) {
	courses, err := GetAllCoursesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"all courses fetched successfully",
		gin.H{
			"courses": courses,
		},
	))
}

// GetCourseWithPlaylistsHandler presents modular views of course sequences.
func GetCourseWithPlaylistsHandler(c *gin.Context) {

	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid courseId",
			nil,
		))
		return
	}

	course, err := GetCourseWithPlaylistsService(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"course fetched successfully",
		gin.H{
			"course": course,
		},
	))
}
