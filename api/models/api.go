package models

// BadRequestResponse represents a bad request error response
// @Description Bad request error response
type BadRequestResponse struct {
	// Error message
	Error string `json:"error" example:"Invalid input" binding:"required"`
}

// UnauthorizedResponse represents an unauthorized error response
// @Description Unauthorized error response
type UnauthorizedResponse struct {
	// Error message
	Error string `json:"error" example:"Unauthorized" binding:"required"`
}

// InternalServerErrorResponse represents an internal server error response
// @Description Internal server error response
type InternalServerErrorResponse struct {
	// Error message
	Error string `json:"error" example:"Internal server error" binding:"required"`
}