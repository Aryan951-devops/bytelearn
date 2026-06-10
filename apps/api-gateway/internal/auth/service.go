package auth

import (
	"errors"

	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser processes logic for user registration.
func RegisterUser(req RegisterRequest) (*models.User, error) {

	existingUser, _ := GetUserByEmail(req.Email)

	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     req.Username,
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         req.Role,
	}

	err = CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser processes logic for user authentication.
func LoginUser(req LoginRequest) (string, error) {

	user, err := GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	accessToken, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
