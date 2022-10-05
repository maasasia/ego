package ego

import "net/http"

type ForbiddenError struct {
	message string
}

func (e ForbiddenError) Error() string {
	return e.message
}

func NewForbiddenError(message string) ForbiddenError {
	return ForbiddenError{
		message: message,
	}
}

func (e ForbiddenError) ToGqlErrorCode() string {
	return "FORBIDDEN"
}

func (e ForbiddenError) ToHttpStatusCode() int {
	return http.StatusForbidden
}
