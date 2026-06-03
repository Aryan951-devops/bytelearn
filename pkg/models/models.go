package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID `json:"user_id"`
	Username            string    `json:"username"`
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	PhoneNo             *string   `json:"phone_no"`
	ProfilePic_Url      *string   `json:"profile_pic_url"`
	ProfilePic_PublicID *string   `json:"-"`
	PasswordHash        string    `json:"-"`
	City                *string   `json:"city"`
	State               *string   `json:"state"`
	Pincode             *string   `json:"pincode"`
	Role                string    `json:"role"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type Video struct {
	ID                 uuid.UUID `json:"video_id"`
	Title              string    `json:"title"`
	Description        *string   `json:"description"`
	Videofile_Url      string    `json:"videofile_url"`
	Videofile_PublicID string    `json:"-"`
	Thumbnail_Url      *string   `json:"thumbnail_url"`
	Thumbnail_PublicID *string   `json:"-"`
	DurationSeconds    int32     `json:"duration_seconds"`
	Views              int64     `json:"views"`
	UploadedBy         uuid.UUID `json:"uploaded_by"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uuid.UUID `json:"comment_id"`
	VideoID   uuid.UUID `json:"video_id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Course struct {
	ID          uuid.UUID `json:"course_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Category    *string   `json:"category"`
	CreatedBy   uuid.UUID `json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Playlist struct {
	ID          uuid.UUID  `json:"playlist_id"`
	Type        string     `json:"type"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	UserID      uuid.UUID  `json:"user_id"`
	CourseID    *uuid.UUID `json:"course_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type PlaylistVideo struct {
	PlaylistID uuid.UUID `json:"playlist_id"`
	VideoID    uuid.UUID `json:"video_id"`
	OrderIndex int32     `json:"order_index"` // Dictates sequence of streaming playback
	CreatedAt  time.Time `json:"created_at"`
}

type Enrollment struct {
	UserID    uuid.UUID `json:"user_id"`
	CourseID  uuid.UUID `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
}

type WatchHistory struct {
	UserID     uuid.UUID `json:"user_id"`
	VideoID    uuid.UUID `json:"video_id"`
	ResumeTime int32     `json:"resume_time"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type VideoLike struct {
	UserID    uuid.UUID `json:"user_id"`
	VideoID   uuid.UUID `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentLike struct {
	UserID    uuid.UUID `json:"user_id"`
	CommentID uuid.UUID `json:"comment_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CodingContest struct {
	ID          uuid.UUID `json:"contest_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PlaylistID  uuid.UUID `json:"playlist_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CodingQuestion struct {
	ID        uuid.UUID `json:"question_id"`
	Metadata  string    `json:"metadata"`
	ContestID uuid.UUID `json:"contest_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Submission struct {
	ID          uuid.UUID `json:"submission_id"`
	QuestionID  uuid.UUID `json:"question_id"`
	UserID      uuid.UUID `json:"user_id"`
	Code        string    `json:"code"`
	Language    string    `json:"language"`
	Verdict     string    `json:"verdict"`
	SubmittedAt time.Time `json:"submitted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// JSONBMap allows structured JSON slices to map directly into PostgreSQL native JSONB columns.
type JSONBMap []interface{}

// Value satisfies driver.Valuer interface, converting map to string byte layout for SQL insertion.
func (j JSONBMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan satisfies sql.Scanner interface, decoding raw database input strings straight into structural maps.
func (j *JSONBMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed inside JSONB conversion block")
	}
	return json.Unmarshal(bytes, j)
}

type Quiz struct {
	ID         uuid.UUID `json:"quiz_id"`
	Title      string    `json:"title"`
	PlaylistID uuid.UUID `json:"playlist_id"`                 // Binds quiz directly to course timeline modules
	Questions  JSONBMap  `json:"questions" gorm:"type:jsonb"` // Scalable schema-less database array storage
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
