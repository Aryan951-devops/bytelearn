package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"user_id"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PhoneNo      *string   `json:"phone_no"`
	ProfilePic   *string   `json:"profile_pic"`
	PasswordHash string    `json:"-"`
	City         *string   `json:"city"`
	State        *string   `json:"state"`
	Pincode      *string   `json:"pincode"`
	Role         string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Video struct {
	ID              string    `json:"video_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	VideofileUrl    string    `json:"videofile_url"`
	ThumbnailUrl    string    `json:"thumbnail_url"`
	DurationSeconds int32     `json:"duration_seconds"`
	Views           int64     `json:"views"`
	UploadedBy      string    `json:"uploaded_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Comment struct {
	ID        string    `json:"comment_id"`
	VideoID   string    `json:"video_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Course struct {
	ID          string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Playlist struct {
	ID          string    `json:"playlist_id"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	CourseID    *string   `json:"course_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PlaylistVideo struct {
	PlaylistID string    `json:"playlist_id"`
	VideoID    string    `json:"video_id"`
	OrderIndex int32     `json:"order_index"` // Dictates sequence of streaming playback
	CreatedAt  time.Time `json:"created_at"`
}

type Enrollment struct {
	UserID    string    `json:"user_id"`
	CourseID  string    `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
}

type WatchHistory struct {
	UserID     string    `json:"user_id"`
	VideoID    string    `json:"video_id"`
	ResumeTime int32     `json:"resume_time"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type VideoLike struct {
	UserID    string    `json:"user_id"`
	VideoID   string    `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentLike struct {
	UserID    string    `json:"user_id"`
	CommentID string    `json:"comment_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CodingContest struct {
	ID          string    `json:"contest_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PlaylistID  string    `json:"playlist_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CodingQuestion struct {
	ID        string    `json:"question_id"`
	Metadata  string    `json:"metadata"`
	ContestID string    `json:"contest_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Submission struct {
	ID          string    `json:"submission_id"`
	QuestionID  string    `json:"question_id"`
	UserID      string    `json:"user_id"`
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
	ID         string    `json:"quiz_id"`
	Title      string    `json:"title"`
	PlaylistID string    `json:"playlist_id"`                 // Binds quiz directly to course timeline modules
	Questions  JSONBMap  `json:"questions" gorm:"type:jsonb"` // Scalable schema-less database array storage
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
