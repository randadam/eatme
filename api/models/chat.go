package models

// ChatRequest represents a chat request
// @Description A chat request
type ChatRequest struct {
	UserID   string   `json:"user_id" binding:"required"`
	Message  string   `json:"message" binding:"required"`
	MealPlan MealPlan `json:"meal_plan" binding:"required"`
	Profile  Profile  `json:"profile" binding:"required"`
}

// ChatResponse represents a chat response
// @Description A chat response
type ChatResponse struct {
	Intent             string                 `json:"intent" binding:"required"`
	ResponseText       string                 `json:"response_text" binding:"required"`
	UIPayload          map[string]interface{} `json:"ui_payload" binding:"required"`
	NewMealPlan        *MealPlan              `json:"new_meal_plan"`
	NeedsClarification bool                   `json:"needs_clarification"`
}
