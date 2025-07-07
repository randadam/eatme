package models

// SuggestChatRequest represents a chat request to the ML backend to suggest a recipe
// @Description A chat request to the ML backend to suggest a recipe
type SuggestChatRequest struct {
	Message string `json:"message" binding:"required"`
}

type InternalSuggestChatRequest struct {
	Message string  `json:"message" binding:"required"`
	Profile Profile `json:"profile" binding:"required"`
}

//	SuggestChatResponse represents a chat response to the ML backend to suggest a recipe
//
// @Description A chat response to the ML backend to suggest a recipe
type SuggestChatResponse struct {
	ResponseText string  `json:"response_text" binding:"required"`
	NewRecipe    *Recipe `json:"new_recipe"`
}

// ModifyChatRequest represents a chat request to the ML backend to modify a recipe
// @Description A chat request to the ML backend to modify a recipe
type ModifyChatRequest struct {
	Message string `json:"message" binding:"required"`
}

type InternalModifyChatRequest struct {
	Message string  `json:"message" binding:"required"`
	Recipe  Recipe  `json:"recipe" binding:"required"`
	Profile Profile `json:"profile" binding:"required"`
}

// ModifyChatResponse represents a chat response to the ML backend to modify a recipe
// @Description A chat response to the ML backend to modify a recipe
type ModifyChatResponse struct {
	ResponseText       string  `json:"response_text" binding:"required"`
	NewRecipe          *Recipe `json:"new_recipe"`
	NeedsClarification bool    `json:"needs_clarification"`
}

// GeneralChatRequest represents a chat request to the ML backend to answer a question
// @Description A chat request to the ML backend to answer a question
type GeneralChatRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type InternalGeneralChatRequest struct {
	UserID   string   `json:"user_id" binding:"required"`
	Message  string   `json:"message" binding:"required"`
	MealPlan MealPlan `json:"meal_plan" binding:"required"`
	Profile  Profile  `json:"profile" binding:"required"`
}

// GeneralChatResponse represents a chat response to the ML backend to answer a question
// @Description A chat response to the ML backend to answer a question
type GeneralChatResponse struct {
	ResponseText string `json:"response_text" binding:"required"`
}
