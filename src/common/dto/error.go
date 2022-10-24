package common_dto

func NewErrorResponse(status int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{Error: errorMessage, StatusCode: status}
}

func NewErrorResponse400(status int, errorMessage string) *ErrorResponse400 {
	return &ErrorResponse400{Error: errorMessage, StatusCode: status}
}

func NewErrorResponse500(status int, errorMessage string) *ErrorResponse500 {
	return &ErrorResponse500{Error: errorMessage, StatusCode: status}
}

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"status"`
}

type ErrorResponse400 ErrorResponse
type ErrorResponse500 ErrorResponse
