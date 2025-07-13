package thread

import (
	"encoding/json"
	"fmt"

	"github.com/ajohnston1219/eatme/api/internal/models"
	"go.uber.org/zap"
)

func ReduceThreadEvents(threadID string, events []models.ThreadEvent) (*models.ThreadState, error) {
	thread := &models.ThreadState{
		ID:          threadID,
		Suggestions: []*models.RecipeSuggestion{},
		RecipeID:    nil,
	}
	zap.L().Debug("reducing thread events", zap.Int("event_count", len(events)))
	for _, event := range events {
		zap.L().Debug("reducing thread event", zap.String("event_type", event.Type), zap.String("event_payload", string(event.Payload)))
		thread.UpdatedAt = event.Timestamp
		switch event.Type {
		case string(models.ThreadEventTypePromptSet):
			var p models.PromptSetEvent
			err := json.Unmarshal(event.Payload, &p)
			if err != nil {
				zap.L().Error("failed to unmarshal prompt set event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			thread.OriginalPrompt = p.Prompt
			thread.CurrentPrompt = p.Prompt
			thread.CreatedAt = event.Timestamp
		case string(models.ThreadEventTypePromptEdited):
			var p models.PromptEditedEvent
			err := json.Unmarshal(event.Payload, &p)
			if err != nil {
				zap.L().Error("failed to unmarshal prompt edited event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			thread.CurrentPrompt = p.Prompt
		case string(models.ThreadEventTypeSuggestionGenerated):
			suggestionEvent := models.SuggestionGeneratedEvent{}
			err := json.Unmarshal(event.Payload, &suggestionEvent)
			if err != nil {
				zap.L().Error("failed to unmarshal suggestion generated event", zap.Error(err))
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
		case string(models.ThreadEventTypeSuggestionAccepted):
			suggestionEvent := models.SuggestionAcceptedEvent{}
			err := json.Unmarshal(event.Payload, &suggestionEvent)
			if err != nil {
				zap.L().Error("failed to unmarshal suggestion accepted event", zap.Error(err))
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
		case string(models.ThreadEventTypeSuggestionRejected):
			suggestionEvent := models.SuggestionRejectedEvent{}
			err := json.Unmarshal(event.Payload, &suggestionEvent)
			if err != nil {
				zap.L().Error("failed to unmarshal suggestion rejected event", zap.Error(err))
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
		case string(models.ThreadEventTypeRecipeModified):
			var recipeEvent models.RecipeModifiedEvent
			err := json.Unmarshal(event.Payload, &recipeEvent)
			if err != nil {
				zap.L().Error("failed to unmarshal recipe modified event", zap.Error(err))
				return nil, ErrInvalidThreadEventPayload
			}
			thread.CurrentRecipe = &recipeEvent.Recipe
		case string(models.ThreadEventTypeQuestionAnswered):
			questionEvent := models.QuestionAnsweredEvent{}
			err := json.Unmarshal(event.Payload, &questionEvent)
			if err != nil {
				zap.L().Error("failed to unmarshal question answered event", zap.Error(err))
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
			zap.L().Error("failed to reduce thread events", zap.Error(err))
			return nil, err
		}
	}
	return thread, nil
}
