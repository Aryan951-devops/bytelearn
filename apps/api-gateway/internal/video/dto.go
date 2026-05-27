package video

type UploadVideoRequest struct {
	Title              string `json:"title" binding:"required"`
	Description        string `json:"description"`
	Videofile_Url      string `json:"videofile_url" binding:"required"`
	Videofile_PublicID string `json:"videofile_public_id" binding:"required"`
	DurationSeconds    int32  `json:"duration_seconds"`
}

type UpdateVideoRequest struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}
