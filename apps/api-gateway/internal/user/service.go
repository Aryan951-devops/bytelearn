package user

import (
	"strings"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/config"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(new_password string, user_id uuid.UUID) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(new_password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	err = UpdatePassword(string(hashedPassword), user_id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserProfile(req UpdateAccountRequest, user_id uuid.UUID,
	tempPath *string) (*models.User, error) {

	var user models.User

	user.ID = user_id

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

	updated_user, err := UpdateUserByID(&user)

	if err != nil {
		return nil, err
	}

	return updated_user, nil
}
