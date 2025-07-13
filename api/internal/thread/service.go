package thread

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ajohnston1219/eatme/api/internal/chat"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/ajohnston1219/eatme/api/internal/recipe"
	"github.com/ajohnston1219/eatme/api/internal/user"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ThreadService struct {
	store         db.Store
	userService   *user.UserService
	recipeService *recipe.RecipeService
	chatService   *chat.ChatService
}

func NewThreadService(store db.Store, recipeService *recipe.RecipeService, chatService *chat.ChatService) *ThreadService {
	return &ThreadService{
		store:         store,
		recipeService: recipeService,
		chatService:   chatService,
	}
}

func (s *ThreadService) getStore(ctx context.Context) db.Store {
	if tx, ok := db.GetTx(ctx); ok {
		return tx
	}
	return s.store
}

func (s *ThreadService) StartSuggestionThread(ctx context.Context, userID string, prompt string) (*models.ThreadState, error) {
	var state *models.ThreadState
	err := s.store.WithTx(func(tx db.Store) error {
		ctx = db.ContextWithTx(ctx, tx)
		profile, err := s.userService.GetProfile(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get profile: %w", err)
		}
		zap.L().Debug("got profile")
		events := []models.ThreadEvent{
			{
				Type:      string(models.ThreadEventTypePromptSet),
				Payload:   []byte(`{"prompt":"` + prompt + `"}`),
				Timestamp: time.Now(),
			},
		}
		suggestionRequest := &models.SuggestChatRequest{
			Profile: *profile,
			Message: prompt,
			History: []string{},
		}
		suggestions, err := s.chatService.GenerateSuggestions(ctx, suggestionRequest)
		if err != nil {
			return fmt.Errorf("failed to generate recipe suggestions: %w", err)
		}
		zap.L().Debug("generated recipe suggestions")
		for _, suggestion := range suggestions.Suggestions {
			zap.L().Debug("generated recipe suggestion", zap.Any("suggestion", *suggestion))
			event := models.SuggestionGeneratedEvent{
				SuggestionID: uuid.New().String(),
				Recipe:       suggestion.Recipe,
				ResponseText: suggestion.ResponseText,
			}
			payload, err := json.Marshal(event)
			zap.L().Debug("generated recipe suggestion event", zap.String("payload", string(payload)))
			if err != nil {
				return ErrInvalidThreadEventPayload
			}
			events = append(events, models.ThreadEvent{
				Type:      string(models.ThreadEventTypeSuggestionGenerated),
				Payload:   payload,
				Timestamp: time.Now(),
			})
		}
		threadID := uuid.New().String()
		thread := models.Thread{
			ID:        threadID,
			Type:      models.ThreadTypeSuggestion,
			Events:    events,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := tx.CreateThread(ctx, userID, thread); err != nil {
			return fmt.Errorf("failed to save thread: %w", err)
		}
		zap.L().Debug("saved thread")
		state, err = ReduceThreadEvents(thread.ID, events)
		if err != nil {
			return fmt.Errorf("failed to reduce thread events: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start suggestion thread: %w", err)
	}
	return state, nil
}

func (s *ThreadService) GetNewSuggestions(ctx context.Context, userID string, threadID string, input models.GetNewSuggestionsRequest) ([]models.RecipeSuggestion, error) {
	var suggestions []models.RecipeSuggestion
	err := s.store.WithTx(func(tx db.Store) error {
		ctx = db.ContextWithTx(ctx, tx)
		profile, err := s.userService.GetProfile(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get profile: %w", err)
		}
		zap.L().Debug("got profile")
		thread, err := tx.GetThread(ctx, threadID)
		if err != nil {
			switch {
			case errors.Is(err, db.ErrNotFound):
				return ErrThreadNotFound
			default:
				return fmt.Errorf("failed to get thread: %w", err)
			}
		}
		state, err := ReduceThreadEvents(threadID, thread.Events)
		if err != nil {
			return fmt.Errorf("failed to reduce thread events: %w", err)
		}

		if input.Prompt != nil {
			promptEvent := models.PromptEditedEvent{
				Prompt: *input.Prompt,
			}
			payload, err := json.Marshal(promptEvent)
			if err != nil {
				return ErrInvalidThreadEventPayload
			}
			event := models.ThreadEvent{
				Type:      string(models.ThreadEventTypePromptEdited),
				Payload:   payload,
				Timestamp: time.Now(),
			}
			if err := s.AppendEventsToThread(ctx, threadID, []models.ThreadEvent{event}); err != nil {
				return fmt.Errorf("failed to append events to thread: %w", err)
			}
			thread.Events = append(thread.Events, event)
			state.CurrentPrompt = *input.Prompt
		}

		suggestionRequest := &models.SuggestChatRequest{
			Profile: *profile,
			Message: state.CurrentPrompt,
			History: []string{},
		}
		for _, suggestion := range state.Suggestions {
			suggestionRequest.History = append(suggestionRequest.History, suggestion.Suggestion.Title)
		}
		suggestionsResponse, err := s.chatService.GenerateSuggestions(ctx, suggestionRequest)
		if err != nil {
			return fmt.Errorf("failed to generate suggestions: %w", err)
		}
		zap.L().Debug("generated suggestions")
		suggestionEvents := make([]models.ThreadEvent, len(suggestionsResponse.Suggestions))
		for i, suggestion := range suggestionsResponse.Suggestions {
			event := models.SuggestionGeneratedEvent{
				SuggestionID: uuid.New().String(),
				Recipe:       suggestion.Recipe,
				ResponseText: suggestion.ResponseText,
			}
			payload, err := json.Marshal(event)
			if err != nil {
				return ErrInvalidThreadEventPayload
			}
			suggestionEvents[i] = models.ThreadEvent{
				Type:      string(models.ThreadEventTypeSuggestionGenerated),
				Payload:   payload,
				Timestamp: time.Now(),
			}
		}
		if err := s.AppendEventsToThread(ctx, threadID, suggestionEvents); err != nil {
			return fmt.Errorf("failed to append events to thread: %w", err)
		}
		zap.L().Debug("appended events to thread")

		suggestions = make([]models.RecipeSuggestion, len(suggestionsResponse.Suggestions))
		for i, suggestion := range suggestionsResponse.Suggestions {
			suggestions[i] = models.RecipeSuggestion{
				ID:           uuid.New().String(),
				ThreadID:     threadID,
				Suggestion:   suggestion.Recipe,
				ResponseText: suggestion.ResponseText,
				Accepted:     false,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
		}
		zap.L().Debug("created suggestions")
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get new suggestions: %w", err)
	}
	return suggestions, nil
}

func (s *ThreadService) AcceptSuggestion(ctx context.Context, userID string, threadID string, suggestionID string) (*models.UserRecipe, error) {
	var recipe *models.UserRecipe
	err := s.store.WithTx(func(tx db.Store) error {
		ctx = db.ContextWithTx(ctx, tx)
		thread, err := tx.GetThread(ctx, threadID)
		if err != nil {
			return fmt.Errorf("failed to get thread: %w", err)
		}
		found := false
		for _, event := range thread.Events {
			if event.Type == string(models.ThreadEventTypeSuggestionGenerated) {
				payload := event.Payload
				var suggestionEvent models.SuggestionGeneratedEvent
				if err := json.Unmarshal(payload, &suggestionEvent); err != nil {
					return fmt.Errorf("failed to unmarshal suggestion generated event: %w", err)
				}
				newRecipe, err := s.recipeService.NewRecipe(ctx, userID, threadID, suggestionEvent.Recipe)
				if err != nil {
					return fmt.Errorf("failed to create new recipe: %w", err)
				}
				recipe = newRecipe
				found = true
				break
			}
		}
		if !found {
			return ErrSuggestionNotFound
		}
		if err := tx.AssociateThreadWithRecipe(ctx, threadID, recipe.ID); err != nil {
			return fmt.Errorf("failed to associate thread with recipe: %w", err)
		}
		zap.L().Debug("associated thread with recipe")
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to accept suggestion: %w", err)
	}
	return recipe, nil
}

func (s *ThreadService) ModifyRecipeViaChat(ctx context.Context, userID string, recipeID string, prompt string) (*models.RecipeBody, error) {
	var recipeBody *models.RecipeBody
	err := s.store.WithTx(func(tx db.Store) error {
		ctx = db.ContextWithTx(ctx, tx)
		profile, err := s.userService.GetProfile(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get profile: %w", err)
		}
		recipe, err := s.recipeService.GetUserRecipe(ctx, userID, recipeID)
		if err != nil {
			return fmt.Errorf("failed to get recipe: %w", err)
		}
		thread, err := tx.GetThread(ctx, recipe.ThreadID)
		if err != nil {
			return fmt.Errorf("failed to get thread: %w", err)
		}
		modifyRequest := &models.ModifyChatRequest{
			Message: prompt,
			Recipe:  recipe.RecipeBody,
			Profile: *profile,
		}
		modifyResponse, err := s.chatService.ModifyRecipeViaChat(ctx, modifyRequest)
		if err != nil {
			return fmt.Errorf("failed to modify recipe: %w", err)
		}
		zap.L().Debug("modified recipe")
		recipeBody = &modifyResponse.NewRecipe
		modifyEvent := models.RecipeModifiedEvent{
			Recipe: *recipeBody,
		}
		payload, err := json.Marshal(modifyEvent)
		if err != nil {
			return ErrInvalidThreadEventPayload
		}
		event := models.ThreadEvent{
			Type:      string(models.ThreadEventTypeRecipeModified),
			Payload:   payload,
			Timestamp: time.Now(),
		}
		if err := s.AppendEventsToThread(ctx, thread.ID, []models.ThreadEvent{event}); err != nil {
			return fmt.Errorf("failed to append events to thread: %w", err)
		}
		if err := s.recipeService.UpdateRecipe(ctx, userID, recipeID, *recipeBody); err != nil {
			return fmt.Errorf("failed to update recipe version: %w", err)
		}
		zap.L().Debug("updated recipe version")
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to modify recipe via chat: %w", err)
	}
	return recipeBody, nil
}

func (s *ThreadService) AnswerCookingQuestion(ctx context.Context, userID string, threadID string, question string) (*models.AnswerCookingQuestionResponse, error) {
	var response *models.AnswerCookingQuestionResponse
	err := s.store.WithTx(func(tx db.Store) error {
		ctx = db.ContextWithTx(ctx, tx)
		profile, err := s.userService.GetProfile(ctx, userID)
		if err != nil {
			return fmt.Errorf("failed to get profile: %w", err)
		}
		thread, err := tx.GetThread(ctx, threadID)
		if err != nil {
			return fmt.Errorf("failed to get thread: %w", err)
		}
		recipeID := thread.RecipeID
		if recipeID == nil {
			return ErrThreadNotAssociatedWithRecipeVersion
		}
		recipe, err := s.recipeService.GetUserRecipe(ctx, userID, *recipeID)
		if err != nil {
			return fmt.Errorf("failed to get recipe: %w", err)
		}
		generalChatRequest := &models.GeneralChatRequest{
			Message: question,
			Recipe:  recipe.RecipeBody,
			Profile: *profile,
		}
		generalChatResponse, err := s.chatService.AnswerCookingQuestion(ctx, generalChatRequest)
		if err != nil {
			return fmt.Errorf("failed to answer cooking question: %w", err)
		}
		questionEvent := models.QuestionAnsweredEvent{
			Question: question,
			Answer:   generalChatResponse.ResponseText,
		}
		payload, err := json.Marshal(questionEvent)
		if err != nil {
			return ErrInvalidThreadEventPayload
		}
		event := models.ThreadEvent{
			Type:      string(models.ThreadEventTypeQuestionAnswered),
			Payload:   payload,
			Timestamp: time.Now(),
		}
		if err := s.AppendEventsToThread(ctx, threadID, []models.ThreadEvent{event}); err != nil {
			return fmt.Errorf("failed to append events to thread: %w", err)
		}
		response = &models.AnswerCookingQuestionResponse{
			Answer: generalChatResponse.ResponseText,
		}
		zap.L().Debug("answered cooking question")
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to answer cooking question: %w", err)
	}
	return response, nil
}

func (s *ThreadService) AppendEventsToThread(ctx context.Context, threadID string, events []models.ThreadEvent) error {
	store := s.getStore(ctx)
	for _, event := range events {
		if err := store.AppendToThread(ctx, threadID, []models.ThreadEvent{event}); err != nil {
			return fmt.Errorf("failed to append to thread: %w", err)
		}
	}
	zap.L().Debug("appended events to thread")
	return nil
}

func (s *ThreadService) GetThreadState(ctx context.Context, threadID string) (*models.ThreadState, error) {
	thread, err := s.store.GetThread(ctx, threadID)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrNotFound):
			return nil, ErrThreadNotFound
		default:
			return nil, fmt.Errorf("failed to get thread: %w", err)
		}
	}
	zap.L().Debug("got thread")
	return ReduceThreadEvents(threadID, thread.Events)
}
