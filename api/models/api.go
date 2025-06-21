package models

// SignupRequest represents the user signup request payload
// @Description User signup request
type SignupRequest struct {
	// User's first name
	FirstName string `json:"first_name" example:"John"`

	// User's last name
	LastName string `json:"last_name" example:"Doe"`

	// User's email address
	Email string `json:"email" example:"john.doe@example.com"`
}

// SignupResponse represents the user signup response
// @Description User signup response containing the new user's ID
type SignupResponse struct {
	// Unique identifier for the created user
	UserID string `json:"user_id" example:"usr_123456789"`
}

// BadRequestResponse represents a bad request error response
// @Description Bad request error response
type BadRequestResponse struct {
	// Error message
	Error string `json:"error" example:"Invalid input"`
}

// UnauthorizedResponse represents an unauthorized error response
// @Description Unauthorized error response
type UnauthorizedResponse struct {
	// Error message
	Error string `json:"error" example:"Unauthorized"`
}

// InternalServerErrorResponse represents an internal server error response
// @Description Internal server error response
type InternalServerErrorResponse struct {
	// Error message
	Error string `json:"error" example:"Internal server error"`
}
