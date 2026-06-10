package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// UpdatePassword changes the password string in the database.
func UpdatePassword(hashedPassword string, userID uuid.UUID) error {

	query := `
		UPDATE users
		SET password_hash = $1,
			updated_at = NOW()
		WHERE user_id = $2
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		hashedPassword,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

// UpdateUserByID updates profile details in the database.
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
	var updatedUser models.User
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
		&updatedUser.ID,
		&updatedUser.Username,
		&updatedUser.Name,
		&updatedUser.Email,
		&updatedUser.PhoneNo,
		&updatedUser.ProfilePic_Url,
		&updatedUser.ProfilePic_PublicID,
		&updatedUser.PasswordHash,
		&updatedUser.City,
		&updatedUser.State,
		&updatedUser.Pincode,
		&updatedUser.Role,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &updatedUser, nil
}
