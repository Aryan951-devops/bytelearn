package user

import (
	"strings"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/config"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ChangePassword processes the request to change a password.
func ChangePassword(newPassword string, userID uuid.UUID) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	err = UpdatePassword(string(hashedPassword), userID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserProfile processes user detail changes.
func UpdateUserProfile(req UpdateAccountRequest, userID uuid.UUID,
	tempPath *string) (*models.User, error) {

	var user models.User

	user.ID = userID

	if strings.TrimSpace(req.PhoneNo) != "" {
		user.PhoneNo = &req.PhoneNo
	}

	if strings.TrimSpace(req.City) != "" {
		user.City = &req.City
	}

	if strings.TrimSpace(req.State) != "" {
		user.State = &req.State
	}

	if strings.TrimSpace(req.Pincode) != "" {
		user.Pincode = &req.Pincode
	}

	if tempPath != nil {
		imageData := config.AppConfig.MediaUploader.UploadProfilePic(*tempPath)

		if imageData.Err != nil {
			return nil, imageData.Err
		}

		user.ProfilePic_Url = &imageData.PublicURL
		user.ProfilePic_PublicID = &imageData.PublicID
	}

	updatedUser, err := UpdateUserByID(&user)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
