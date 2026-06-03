package auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

func GetUserByEmail(email string) (*models.User, error) {

	query := `
		SELECT 
			user_id, username, name, email, phone_no, 
			profile_pic_url, profile_pic_public_id, password_hash,
			city, state, pincode, role, created_at, updated_at
		FROM users 
		WHERE email = $1
	`
	var user models.User
	err := database.DB.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.PhoneNo,
		&user.ProfilePic_Url,
		&user.ProfilePic_PublicID,
		&user.PasswordHash,
		&user.City,
		&user.State,
		&user.Pincode,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByUserId(id uuid.UUID) (*models.User, error) {
	query := `
		SELECT 
			user_id, username, name, email, phone_no, 
			profile_pic_url, profile_pic_public_id, password_hash,
			city, state, pincode, role, created_at, updated_at
		FROM users 
		WHERE user_id = $1
	`
	var user models.User
	err := database.DB.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Name,
		&user.Email,
		&user.PhoneNo,
		&user.ProfilePic_Url,
		&user.ProfilePic_PublicID,
		&user.PasswordHash,
		&user.City,
		&user.State,
		&user.Pincode,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *models.User) error {

	query := `
		INSERT INTO users (
			username, name, email, password_hash, role
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id
	`

	return database.DB.QueryRow(
		context.Background(),
		query,
		user.Username,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(&user.ID)
}
