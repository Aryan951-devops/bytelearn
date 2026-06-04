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

type CodingPractice struct {
	ID          uuid.UUID `json:"contest_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PlaylistID  uuid.UUID `json:"playlist_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CodingQuestion struct {
	ID            uuid.UUID `json:"question_id"`
	ContestID     uuid.UUID `json:"contest_id"`
	Title         string    `json:"title"`
	Difficulty    string    `json:"difficulty"`
	Statement     string    `json:"statement"`
	Constraints   *string   `json:"constraints"`
	InputFormat   *string   `json:"input_format"`
	OutputFormat  *string   `json:"output_format"`
	TimeLimitMS   int32     `json:"time_limit_ms"`
	MemoryLimitMB int32     `json:"memory_limit_mb"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TestCase struct {
	ID             uuid.UUID `json:"testcase_id"`
	QuestionID     uuid.UUID `json:"question_id"`
	Input          string    `json:"input"`
	ExpectedOutput string    `json:"expected_output"`
	IsHidden       bool      `json:"is_hidden"`
	CreatedAt      time.Time `json:"created_at"`
}

type Submission struct {
	ID          uuid.UUID `json:"submission_id"`
	QuestionID  uuid.UUID `json:"question_id"`
	UserID      uuid.UUID `json:"user_id"`
	Code        string    `json:"code"`
	Language    string    `json:"language"`
	Status      string    `json:"status"`
	PassedCases int32     `json:"passed_cases"`
	TotalCases  int32     `json:"total_cases"`
	Started_At  time.Time `json:"started_at"`
	Finished_At time.Time `json:"finished_at"`
	SubmittedAt time.Time `json:"submitted_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubmissionResults struct {
	SubmissionID uuid.UUID `json:"submission_id"`
	TestCaseID   uuid.UUID `json:"testcase_id"`
	ActualOutput *string   `json:"actual_output"`
	ErrorOutput  *string   `json:"error_output"`
	IsPassed     bool      `json:"is_passed"`
	Verdict      string    `json:"verdict"`
	RuntimeMS    *int32    `json:"runtime_ms"`
	MemoryKB     *int32    `json:"memory_kb"`
	CreatedAt    time.Time `json:"created_at"`
}

type SubmissionJob struct {
	SubmissionID uuid.UUID `json:"submission_id"`
	IsHidden     bool      `json:"is_hidden"`
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
