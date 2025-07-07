package chat

import (
	"context"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/services/recipe"
	"github.com/ajohnston1219/eatme/api/internal/services/user"
	"github.com/ajohnston1219/eatme/api/models"
)

type ChatService struct {
	mlClient      *clients.MLClient
	userService   *user.UserService
	recipeService *recipe.RecipeService
}

func NewChatService(mlClient *clients.MLClient, userService *user.UserService, recipeService *recipe.RecipeService) *ChatService {
	return &ChatService{
		mlClient:      mlClient,
		userService:   userService,
		recipeService: recipeService,
	}
}

func (s *ChatService) SuggestChat(ctx context.Context, userID string, req *models.SuggestChatRequest) (*models.SuggestChatResponse, error) {
	profile, err := s.userService.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	internalReq := &models.InternalSuggestChatRequest{
		Message: req.Message,
		Profile: profile,
	}
	resp, err := s.mlClient.SuggestChat(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	newRecipe, err := s.recipeService.NewRecipe(ctx, userID, resp.NewRecipe)
	if err != nil {
		return nil, err
	}

	resp.RecipeID = newRecipe.ID

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
