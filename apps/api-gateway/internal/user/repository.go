package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/database"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/google/uuid"
)

func UpdatePassword(hashed_password string, user_id uuid.UUID) error {

	query := `
		UPDATE users
		SET password_hash = $1,
			updated_at = NOW()
		WHERE user_id = $2
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		hashed_password,
		user_id,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUserByID(user *models.User) (*models.User, error) {

	query := `
		UPDATE users
		SET
			phone_no = COALESCE($1, phone_no),
			city = COALESCE($2, city),
			state = COALESCE($3, state),
			pincode = COALESCE($4, pincode),
			profile_pic_url = COALESCE($5, profile_pic_url),
			profile_pic_public_id = COALESCE($6, profile_pic_public_id),
			updated_at = NOW()
		WHERE user_id = $7
		RETURNING	
			user_id, username, name, email, phone_no, 
			profile_pic_url, profile_pic_public_id, password_hash,
			city, state, pincode, role, created_at, updated_at
	`
	var updated_user models.User
	err := database.DB.QueryRow(context.Background(),
		query,
		user.PhoneNo,
		user.City,
		user.State,
		user.Pincode,
		user.ProfilePic_Url,
		user.ProfilePic_PublicID,
		user.ID,
	).Scan(
		&updated_user.ID,
		&updated_user.Username,
		&updated_user.Name,
		&updated_user.Email,
		&updated_user.PhoneNo,
		&updated_user.ProfilePic_Url,
		&updated_user.ProfilePic_PublicID,
		&updated_user.PasswordHash,
		&updated_user.City,
		&updated_user.State,
		&updated_user.Pincode,
		&updated_user.Role,
		&updated_user.CreatedAt,
		&updated_user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &updated_user, nil
}
