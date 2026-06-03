package user

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ChangePasswordHandler(c *gin.Context) {

	var req ChangePasswordRequest
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

	user, ok := ctxUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			"Internal context type mismatch",
			nil,
		))
		return
	}

	err := ChangePassword(req.NewPassword, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"password updated successfully",
		nil,
	))
}

func GetUserHandler(c *gin.Context) {

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user, ok := ctxUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			"Internal context type mismatch",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"user data fetched successfully",
		gin.H{
			"user": user,
		},
	))
}

func UpdateAccountHandler(c *gin.Context) {

	var req UpdateAccountRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
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

	user, ok := ctxUser.(*models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			"Internal context type mismatch",
			nil,
		))
		return
	}

	file, err := c.FormFile("profile_pic")

	var updated_user *models.User

	if err == nil {
		if err := utils.IsImageAllowed(file); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				err.Error(),
				nil,
			))
			return
		}

		os.MkdirAll("./temp", os.ModePerm)

		// unique file name
		tempPath := fmt.Sprintf(
			"./temp/%s%s",
			uuid.New().String(),
			filepath.Ext(file.Filename),
		)

		if err := c.SaveUploadedFile(file, tempPath); err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewResponse(
				"failed to save uploaded file",
				nil,
			))
			return
		}
		defer os.Remove(tempPath)

		updated_user, err = UpdateUserProfile(req, user.ID, &tempPath)
	} else {
		updated_user, err = UpdateUserProfile(req, user.ID, nil)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			"failed to update account",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"user data updated successfully",
		gin.H{
			"user": updated_user,
		},
	))
}
