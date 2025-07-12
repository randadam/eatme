package models

import (
	"encoding/json"
	"time"
)

type ThreadType string

const (
	ThreadTypeSuggestion ThreadType = "Suggestion"
)

// @Description Thread represents a thread of events that occurred as part of a suggestion thread
type Thread struct {
	ID        string        `json:"id" binding:"required"`
	Type      ThreadType    `json:"type" binding:"required"`
	RecipeID  *string       `json:"recipe_id"`
	Events    []ThreadEvent `json:"events" binding:"required"`
	CreatedAt time.Time     `json:"created_at" binding:"required"`
	UpdatedAt time.Time     `json:"updated_at" binding:"required"`
}

// @Description ThreadEvent represents an event that occurred as part of a suggestion thread
type ThreadEvent struct {
	Type      string          `json:"type" binding:"required"`
	Payload   json.RawMessage `json:"payload" binding:"required"`
	Timestamp time.Time       `json:"timestamp" binding:"required"`
}

type ThreadEventType string

const (
	ThreadEventTypePromptSet           ThreadEventType = "PromptSet"
	ThreadEventTypePromptEdited        ThreadEventType = "PromptEdited"
	ThreadEventTypeSuggestionGenerated ThreadEventType = "SuggestionGenerated"
	ThreadEventTypeSuggestionAccepted  ThreadEventType = "SuggestionAccepted"
	ThreadEventTypeRecipeModified      ThreadEventType = "RecipeModified"
	ThreadEventTypeQuestionAnswered    ThreadEventType = "QuestionAnswered"
)

// @Description PromptSetEvent represents setting the inital prompt
type PromptSetEvent struct {
	Prompt string `json:"prompt" binding:"required"`
}

// @Description PromptEditedEvent represents editing the prompt
type PromptEditedEvent struct {
	Prompt string `json:"prompt" binding:"required"`
}

// @Description SuggestionGeneratedEvent represents generating a new recipe suggestion
type SuggestionGeneratedEvent struct {
	SuggestionID string     `json:"suggestion_id" binding:"required"`
	Recipe       RecipeBody `json:"recipe" binding:"required"`
	ResponseText string     `json:"response_text" binding:"required"`
}

// @Description SuggestionAcceptedEvent represents accepting a recipe suggestion
type SuggestionAcceptedEvent struct {
	SuggestionID string `json:"suggestion_id" binding:"required"`
	RecipeID     string `json:"recipe_id" binding:"required"`
}

// @Description SuggestionRejectedEvent represents rejecting a recipe suggestion
type SuggestionRejectedEvent struct {
	SuggestionID string `json:"suggestion_id" binding:"required"`
}

// @Description RecipeSuggestion represents a suggestion for a recipe
type RecipeSuggestion struct {
	ID           string     `json:"id" binding:"required"`
	ThreadID     string     `json:"thread_id" binding:"required"`
	Suggestion   RecipeBody `json:"suggestion" binding:"required"`
	ResponseText string     `json:"response_text" binding:"required"`
	Accepted     bool       `json:"accepted" binding:"required"`
	CreatedAt    time.Time  `json:"created_at" binding:"required"`
	UpdatedAt    time.Time  `json:"updated_at" binding:"required"`
}

// @Description RecipeModifiedEvent represents modifying the recipe
type RecipeModifiedEvent struct {
	Recipe RecipeBody `json:"recipe" binding:"required"`
}

// @Description QuestionAnsweredEvent represents answering a question
type QuestionAnsweredEvent struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

// @Description A thread of suggestions for a recipe
type ThreadState struct {
	ID             string              `json:"id" binding:"required"`
	RecipeID       *string             `json:"recipe_id"`
	OriginalPrompt string              `json:"original_prompt" binding:"required"`
	CurrentPrompt  string              `json:"current_prompt" binding:"required"`
	Suggestions    []*RecipeSuggestion `json:"suggestions" binding:"required"`
	CurrentRecipe  *RecipeBody         `json:"current_recipe"`
	CreatedAt      time.Time           `json:"created_at" binding:"required"`
	UpdatedAt      time.Time           `json:"updated_at" binding:"required"`
}

// @Description StartSuggestionThreadRequest represents a request to start a suggestion thread
type StartSuggestionThreadRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

// @Description GetNewSuggestionsRequest represents a request to get new suggestions
type GetNewSuggestionsRequest struct {
	Prompt *string `json:"prompt"`
}

// @Description ModifyRecipeViaChatRequest represents a request to modify a recipe via chat
type ModifyRecipeViaChatRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

// @Description AnswerCookingQuestionRequest represents a request to answer a cooking question
type AnswerCookingQuestionRequest struct {
	Question string `json:"question" binding:"required"`
}

// @Description AnswerCookingQuestionResponse represents a response to an answer cooking question
type AnswerCookingQuestionResponse struct {
	Answer string `json:"answer" binding:"required"`
}
