package course

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	course := router.Group("/course")

	course.POST("/", auth.AuthMiddleware(), CreateCourseHandler)
	course.GET("/all", GetAllCoursesHandler)
	course.GET("/:courseId", GetCourseWithPlaylistsHandler)
	course.PATCH("/:courseId", UpdateCourseHandler)

}
