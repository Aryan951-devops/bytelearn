package utils

// ResponseFromUpload holds the details returned from a media upload.
type ResponseFromUpload struct {
	Err       error
	PublicURL string
	PublicID  string
}
