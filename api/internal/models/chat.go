package models

// @Description SuggestChatRequest represents a chat request to the ML backend to suggest a recipe
type SuggestChatRequest struct {
	Message string   `json:"message" binding:"required"`
	Profile Profile  `json:"profile" binding:"required"`
	History []string `json:"history" binding:"required"`
}

// @Description NextSuggestionRequest represents a chat request to the ML backend to get the next recipe suggestion
type NextSuggestionRequest struct {
	ThreadID string `json:"thread_id" binding:"required"`
}

type InternalSuggestChatRequest struct {
	Message string   `json:"message" binding:"required"`
	Profile Profile  `json:"profile" binding:"required"`
	History []string `json:"history" binding:"required"`
}

// @Description RecipeSuggestion represents a recipe suggestion from the ML backend
type Suggestion struct {
	Recipe       RecipeBody `json:"recipe" binding:"required"`
	ResponseText string     `json:"response_text" binding:"required"`
}

// @Description SuggestChatResponse represents a chat response to the ML backend to suggest a recipe
type SuggestChatResponse struct {
	ThreadID    string        `json:"thread_id" binding:"required"`
	Suggestions []*Suggestion `json:"suggestions" binding:"required"`
}

// @Description ModifyChatRequest represents a chat request to the ML backend to modify a recipe
type ModifyChatRequest struct {
	Message string     `json:"message" binding:"required"`
	Recipe  RecipeBody `json:"recipe" binding:"required"`
	Profile Profile    `json:"profile" binding:"required"`
}

type InternalModifyChatRequest struct {
	Message string     `json:"message" binding:"required"`
	Recipe  RecipeBody `json:"recipe" binding:"required"`
	Profile Profile    `json:"profile" binding:"required"`
}

// @Description ModifyChatResponse represents a chat response to the ML backend to modify a recipe
type ModifyChatResponse struct {
	ResponseText       string     `json:"response_text" binding:"required"`
	NewRecipe          RecipeBody `json:"new_recipe" binding:"required"`
	NeedsClarification bool       `json:"needs_clarification" binding:"required"`
}

// @Description GeneralChatRequest represents a chat request to the ML backend to answer a question
type GeneralChatRequest struct {
	Message string     `json:"message" binding:"required"`
	Recipe  RecipeBody `json:"recipe" binding:"required"`
	Profile Profile    `json:"profile" binding:"required"`
}

type InternalGeneralChatRequest struct {
	Message string     `json:"message" binding:"required"`
	Recipe  RecipeBody `json:"recipe" binding:"required"`
	Profile Profile    `json:"profile" binding:"required"`
}

// @Description GeneralChatResponse represents a chat response to the ML backend to answer a question
type GeneralChatResponse struct {
	ResponseText string `json:"response_text" binding:"required"`
}
