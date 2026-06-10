package course

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines the HTTP routes for the course package.
func RegisterRoutes(router *gin.RouterGroup) {

	course := router.Group("/course")

	course.POST("/", auth.Middleware(), CreateCourseHandler)
	course.GET("/all", GetAllCoursesHandler)
	course.GET("/:courseId", GetCourseWithPlaylistsHandler)
	course.PATCH("/:courseId", UpdateCourseHandler)

}
