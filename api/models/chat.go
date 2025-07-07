package models

// Chat request
// @Description A chat request
type ChatRequest struct {
	MealPlanID string `json:"meal_plan_id" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

// InternalChatRequest represents a chat request to the ML backend
// @Description A chat request to the ML backend
type InternalChatRequest struct {
	UserID   string   `json:"user_id" binding:"required"`
	Message  string   `json:"message" binding:"required"`
	MealPlan MealPlan `json:"meal_plan" binding:"required"`
	Profile  Profile  `json:"profile" binding:"required"`
}

// ChatResponse represents a chat response
// @Description A chat response
type ChatResponse struct {
	Intent             string    `json:"intent" binding:"required"`
	ResponseText       string    `json:"response_text" binding:"required"`
	NewMealPlan        *MealPlan `json:"new_meal_plan"`
	NeedsClarification bool      `json:"needs_clarification"`
}
