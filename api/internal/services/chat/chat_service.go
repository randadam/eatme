package chat

import (
	"context"
	"errors"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/services/recipe"
	"github.com/ajohnston1219/eatme/api/internal/services/user"
	"github.com/ajohnston1219/eatme/api/models"
	"github.com/google/uuid"
)

type ChatService struct {
	mlClient      clients.MLClient
	userService   *user.UserService
	recipeService *recipe.RecipeService
}

func NewChatService(mlClient clients.MLClient, userService *user.UserService, recipeService *recipe.RecipeService) *ChatService {
	return &ChatService{
		mlClient:      mlClient,
		userService:   userService,
		recipeService: recipeService,
	}
}

func (s *ChatService) StartSuggestChat(ctx context.Context, userID string, req *models.SuggestChatRequest) (*models.SuggestChatResponse, error) {
	profile, err := s.userService.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	internalReq := &models.InternalSuggestChatRequest{
		Message: req.Message,
		Profile: profile,
		History: []string{},
	}
	resp, err := s.mlClient.SuggestChat(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	suggestionThread := models.SuggestionThread{
		ID:             uuid.New().String(),
		OriginalPrompt: req.Message,
		Suggestions:    []models.RecipeSuggestion{},
	}
	firstSuggestion := models.RecipeSuggestion{
		ID:           uuid.New().String(),
		ThreadID:     suggestionThread.ID,
		Suggestion:   resp.NewRecipe,
		ResponseText: resp.ResponseText,
		Accepted:     false,
	}
	suggestionThread.Suggestions = append(suggestionThread.Suggestions, firstSuggestion)
	resp.ThreadID = suggestionThread.ID

	err = s.recipeService.NewSuggestionThread(ctx, userID, suggestionThread)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ChatService) AcceptRecipeSuggestion(ctx context.Context, userID string, threadID string) (models.UserRecipe, error) {
	currentThread, err := s.recipeService.GetSuggestionThread(ctx, threadID)
	if err != nil {
		return models.UserRecipe{}, err
	}

	if len(currentThread.Suggestions) == 0 {
		return models.UserRecipe{}, errors.New("no suggestions found")
	}

	currentSuggestion := currentThread.Suggestions[len(currentThread.Suggestions)-1]

	return s.recipeService.AcceptSuggestion(ctx, userID, threadID, currentSuggestion)
}

func (s *ChatService) GetNextSuggestion(ctx context.Context, userID string, threadID string) (*models.SuggestChatResponse, error) {
	profile, err := s.userService.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	currentThread, err := s.recipeService.GetSuggestionThread(ctx, threadID)
	if err != nil {
		return nil, err
	}

	history := make([]string, len(currentThread.Suggestions))
	for i, suggestion := range currentThread.Suggestions {
		history[i] = suggestion.Suggestion.Title
	}

	req := models.InternalSuggestChatRequest{
		Message: currentThread.OriginalPrompt,
		Profile: profile,
		History: history,
	}
	resp, err := s.mlClient.SuggestChat(ctx, &req)
	if err != nil {
		return nil, err
	}

	nextSuggestion := models.RecipeSuggestion{
		ID:           uuid.New().String(),
		ThreadID:     threadID,
		Suggestion:   resp.NewRecipe,
		ResponseText: resp.ResponseText,
		Accepted:     false,
	}

	err = s.recipeService.AppendToSuggestionThread(ctx, threadID, nextSuggestion)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ChatService) ModifyChat(ctx context.Context, userID string, recipeID string, req *models.ModifyChatRequest) (*models.ModifyChatResponse, error) {
	profile, err := s.userService.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	recipe, err := s.recipeService.GetUserRecipe(ctx, userID, recipeID)
	if err != nil {
		return nil, err
	}

	internalReq := &models.InternalModifyChatRequest{
		Message: req.Message,
		Recipe:  recipe.RecipeBody,
		Profile: profile,
	}
	resp, err := s.mlClient.ModifyChat(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	err = s.recipeService.UpdateRecipe(ctx, userID, recipeID, resp.NewRecipe)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ChatService) GeneralChat(ctx context.Context, userID string, recipeID string, req *models.GeneralChatRequest) (*models.GeneralChatResponse, error) {
	profile, err := s.userService.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	recipe, err := s.recipeService.GetUserRecipe(ctx, userID, recipeID)
	if err != nil {
		return nil, err
	}

	internalReq := &models.InternalGeneralChatRequest{
		Message: req.Message,
		Profile: profile,
		Recipe:  recipe.RecipeBody,
	}
	return s.mlClient.GeneralChat(ctx, internalReq)
}
