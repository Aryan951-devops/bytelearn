package utils

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse(message string, data any) APIResponse {
	return APIResponse{
		Message: message,
		Data:    data,
	}
}
