package custom_errors

import (
	"fmt"
	"net/http"
)

const (
	CodeErrorInternalServer = "internal_error"
	CodeErrorInvalidBody    = "invalid_body"

	MessageErrorInternalServer = "internal server error on connect with %s"
	MessageErrorParseError     = "error to parse response"
)

type ApiError struct {
	StatusCode int `json:"statusCode"`

	ErrorCode string `json:"errorCode"`

	Message string `json:"Message"`

	Detail string `json:"details,omitempty"`
}

func (apiErr ApiError) Error() string {
	return fmt.Sprintf(`{"status":%d, "code":"%s", "message":"%s", "detail":%+v}`,
		apiErr.StatusCode, apiErr.ErrorCode, apiErr.Message, apiErr.Detail)
}

func NewInternalServerError(code string, message string, detail string) error {
	return &ApiError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  code,
		Message:    message,
		Detail:     detail,
	}
}

func NewValidationError(code string, message string, detail string) error {
	return &ApiError{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  code,
		Message:    message,
		Detail:     detail,
	}
}
