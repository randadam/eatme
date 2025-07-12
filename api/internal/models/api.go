package models

// @Description APIError represents an API error response
type APIError struct {
	Code    string `json:"code" binding:"required"`
	Message string `json:"message" binding:"required"`
	Details string `json:"details,omitempty"`
	Field   string `json:"field,omitempty"`
}

type APIErrorOps = func(e *APIError)

func NewAPIError(code, message string, ops ...APIErrorOps) APIError {
	e := APIError{
		Code:    code,
		Message: message,
	}
	for _, op := range ops {
		op(&e)
	}
	return e
}

func WithDetails(details string) APIErrorOps {
	return func(e *APIError) {
		e.Details = details
	}
}

func WithField(field string) APIErrorOps {
	return func(e *APIError) {
		e.Field = field
	}
}

func (e APIError) Error() string {
	return e.Message
}

var (
	// Common
	ApiErrUnauthorized = NewAPIError("UNAUTHORIZED", "You must be logged in")
	ApiErrBadRequest   = NewAPIError("BAD_REQUEST", "Invalid request")
	ApiErrInternal     = NewAPIError("INTERNAL_SERVER_ERROR", "Internal server error")

	// User
	ApiErrEmailExists     = NewAPIError("EMAIL_EXISTS", "Email already exists", WithField("email"))
	ApiErrUserNotFound    = NewAPIError("USER_NOT_FOUND", "User not found")
	ApiErrProfileNotFound = NewAPIError("PROFILE_NOT_FOUND", "Profile not found")

	// Recipe
	ApiErrRecipeNotFound = NewAPIError("RECIPE_NOT_FOUND", "Recipe not found")

	// Thread
	ApiErrThreadNotFound = NewAPIError("THREAD_NOT_FOUND", "Thread not found")
)
