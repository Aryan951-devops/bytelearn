package auth

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	user, err := RegisterUser(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"user registered successfully",
		gin.H{
			"user": user,
		},
	))
}

func LoginHandler(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	token, err := LoginUser(req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.SetCookie("access_token", token, 3600*24*7, "/", "", false, true) // secure=false for localhost

	c.JSON(http.StatusOK, utils.NewResponse(
		"login successful",
		nil,
	))
}

func LogoutHandler(c *gin.Context) {

	c.SetCookie("access_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, utils.NewResponse(
		"logged out successfully",
		nil,
	))
}

func VerifyAccessToken(c *gin.Context) {
	_, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"invalid access token",
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"access token is valid",
		nil,
	))
}
