package ego

import "net/http"

type AuthenticationError struct {
	message string
}

func (e AuthenticationError) Error() string {
	return e.message
}

func NewAuthenticationError(message string) AuthenticationError {
	return AuthenticationError{
		message: message,
	}
}

func (e AuthenticationError) ToGqlErrorCode() string {
	return "UNAUTHENTICATED"
}

func (e AuthenticationError) ToHttpStatusCode() int {
	return http.StatusUnauthorized
}
