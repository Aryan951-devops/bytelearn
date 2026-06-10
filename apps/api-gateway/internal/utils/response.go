package utils

// APIResponse standardizes structured REST API outputs.
type APIResponse struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

// NewResponse initializes an APIResponse structural model instance.
func NewResponse(message string, data any) APIResponse {
	return APIResponse{
		Message: message,
		Data:    data,
	}
}
