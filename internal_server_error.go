package ego

import "net/http"

type InternalServerError struct {
	message string
}

func (e InternalServerError) Error() string {
	return e.message
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		message: message,
	}
}

func (e InternalServerError) ToGqlErrorCode() string {
	return "INTERNAL_SERVER_ERROR"
}

func (e InternalServerError) ToHttpStatusCode() int {
	return http.StatusInternalServerError
}
