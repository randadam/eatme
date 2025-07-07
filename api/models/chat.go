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
	ResponseText string     `json:"response_text" binding:"required"`
	RecipeID     string     `json:"recipe_id" binding:"required"`
	NewRecipe    RecipeBody `json:"new_recipe" binding:"required"`
}

// ModifyChatRequest represents a chat request to the ML backend to modify a recipe
// @Description A chat request to the ML backend to modify a recipe
type ModifyChatRequest struct {
	Message string `json:"message" binding:"required"`
}

type InternalModifyChatRequest struct {
	Message string     `json:"message" binding:"required"`
	Recipe  RecipeBody `json:"recipe" binding:"required"`
	Profile Profile    `json:"profile" binding:"required"`
}

// ModifyChatResponse represents a chat response to the ML backend to modify a recipe
// @Description A chat response to the ML backend to modify a recipe
type ModifyChatResponse struct {
	ResponseText       string     `json:"response_text" binding:"required"`
	NewRecipe          RecipeBody `json:"new_recipe" binding:"required"`
	NeedsClarification bool       `json:"needs_clarification" binding:"required"`
}

// GeneralChatRequest represents a chat request to the ML backend to answer a question
// @Description A chat request to the ML backend to answer a question
type GeneralChatRequest struct {
	Message string `json:"message" binding:"required"`
}

type InternalGeneralChatRequest struct {
	Message string     `json:"message" binding:"required"`
	Recipe  RecipeBody `json:"recipe" binding:"required"`
	Profile Profile    `json:"profile" binding:"required"`
}

// GeneralChatResponse represents a chat response to the ML backend to answer a question
// @Description A chat response to the ML backend to answer a question
type GeneralChatResponse struct {
	ResponseText string `json:"response_text" binding:"required"`
}
