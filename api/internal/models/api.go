package models

// APIError represents an API error response
// @Description API error response
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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
	ErrUnauthorized = NewAPIError("UNAUTHORIZED", "You must be logged in")
	ErrBadRequest   = NewAPIError("BAD_REQUEST", "Invalid request")
	ErrInternal     = NewAPIError("INTERNAL_SERVER_ERROR", "Internal server error")

	// User
	ErrEmailExists     = NewAPIError("EMAIL_EXISTS", "Email already exists", WithField("email"))
	ErrUserNotFound    = NewAPIError("USER_NOT_FOUND", "User not found")
	ErrProfileNotFound = NewAPIError("PROFILE_NOT_FOUND", "Profile not found")

	// Recipe
	ErrRecipeNotFound = NewAPIError("RECIPE_NOT_FOUND", "Recipe not found")

	// Chat
	ErrSuggestionThreadNotFound = NewAPIError("SUGGESTION_THREAD_NOT_FOUND", "Suggestion thread not found")
	ErrSuggestionNotFound       = NewAPIError("SUGGESTION_NOT_FOUND", "Suggestion not found")
)
