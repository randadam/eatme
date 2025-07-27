package thread

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/ajohnston1219/eatme/api/internal/utils/logger"
	"go.uber.org/zap"
)

func ReduceThreadEvents(ctx context.Context, threadID string, events []models.ThreadEvent, originalState *models.ThreadState) (*models.ThreadState, error) {
	thread := &models.ThreadState{
		ID:          threadID,
		Suggestions: []*models.RecipeSuggestion{},
		RecipeID:    nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if originalState != nil {
		thread.OriginalPrompt = originalState.OriginalPrompt
		thread.CurrentPrompt = originalState.CurrentPrompt
		thread.Suggestions = originalState.Suggestions
		thread.ChatHistory = originalState.ChatHistory
		thread.CurrentRecipe = originalState.CurrentRecipe
		thread.ModifiedRecipe = originalState.ModifiedRecipe
		thread.CreatedAt = originalState.CreatedAt
		thread.UpdatedAt = originalState.UpdatedAt
	}
	logger.Logger(ctx).Debug("reducing thread events", zap.Int("event_count", len(events)))
	for _, event := range events {
		logger.Logger(ctx).Debug("reducing thread event", zap.String("event_type", string(event.Type)), zap.Any("event", event))
		thread.UpdatedAt = event.Timestamp
		switch event.Type {
		case models.ThreadEventTypePromptSet:
			var p models.PromptSetEvent
			err := json.Unmarshal(event.Payload, &p)
			if err != nil {
				logger.Logger(ctx).Error("failed to unmarshal prompt set event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			thread.OriginalPrompt = p.Prompt
			thread.CurrentPrompt = p.Prompt
			thread.CreatedAt = event.Timestamp
		case models.ThreadEventTypePromptEdited:
			var p models.PromptEditedEvent
			err := json.Unmarshal(event.Payload, &p)
			if err != nil {
				logger.Logger(ctx).Error("failed to unmarshal prompt edited event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			thread.CurrentPrompt = p.Prompt
		case models.ThreadEventTypeSuggestionGenerated:
			suggestionEvent := models.SuggestionGeneratedEvent{}
			err := json.Unmarshal(event.Payload, &suggestionEvent)
			if err != nil {
				logger.Logger(ctx).Error("failed to unmarshal suggestion generated event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			suggestion := &models.RecipeSuggestion{
				ID:           suggestionEvent.SuggestionID,
				ThreadID:     threadID,
				Suggestion:   suggestionEvent.Recipe,
				ResponseText: suggestionEvent.ResponseText,
				Accepted:     false,
			}
			thread.Suggestions = append(thread.Suggestions, suggestion)
		case models.ThreadEventTypeSuggestionAccepted:
			suggestionEvent := models.SuggestionAcceptedEvent{}
			err := json.Unmarshal(event.Payload, &suggestionEvent)
			if err != nil {
				logger.Logger(ctx).Error("failed to unmarshal suggestion accepted event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			found := false
			for i, suggestion := range thread.Suggestions {
				if suggestion.ID == suggestionEvent.SuggestionID {
					thread.Suggestions[i].Accepted = true
					thread.Suggestions[i].UpdatedAt = event.Timestamp
					thread.CurrentRecipe = &suggestion.Suggestion
					found = true
					break
				}
			}
			if !found {
				return nil, ErrSuggestionNotFound
			}
		case models.ThreadEventTypeSuggestionRejected:
			suggestionEvent := models.SuggestionRejectedEvent{}
			err := json.Unmarshal(event.Payload, &suggestionEvent)
			if err != nil {
				logger.Logger(ctx).Error("failed to unmarshal suggestion rejected event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			found := false
			for i, suggestion := range thread.Suggestions {
				if suggestion.ID == suggestionEvent.SuggestionID {
					thread.Suggestions[i].Rejected = true
					thread.Suggestions[i].UpdatedAt = event.Timestamp
					found = true
					break
				}
			}
			if !found {
				return nil, ErrSuggestionNotFound
			}
		case models.ThreadEventTypeRecipeModified:
			var recipeEvent models.RecipeModifiedEvent
			err := json.Unmarshal(event.Payload, &recipeEvent)
			if err != nil {
				logger.Logger(ctx).Error("failed to unmarshal recipe modified event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			thread.ModifiedRecipe = &recipeEvent.Recipe
		case models.ThreadEventTypeRecipeModificationAccepted:
			thread.CurrentRecipe = thread.ModifiedRecipe
			thread.ModifiedRecipe = nil
		case models.ThreadEventTypeRecipeModificationRejected:
			thread.ModifiedRecipe = nil
		case models.ThreadEventTypeQuestionAnswered:
			questionEvent := models.QuestionAnsweredEvent{}
			err := json.Unmarshal(event.Payload, &questionEvent)
			if err != nil {
				logger.Logger(ctx).Error("failed to unmarshal question answered event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			thread.ChatHistory = append(thread.ChatHistory, &models.ChatMessage{
				Source:  "user",
				Message: questionEvent.Question,
			})
			thread.ChatHistory = append(thread.ChatHistory, &models.ChatMessage{
				Source:  "assistant",
				Message: questionEvent.Answer,
			})
		default:
			err := fmt.Errorf("%w: invalid thread event type: %s", ErrInvalidThreadEventType, event.Type)
			logger.Logger(ctx).Error("failed to reduce thread events", zap.Error(err))
			return nil, err
		}
	}
	return thread, nil
}
