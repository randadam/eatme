package chat

import (
	"context"
	"fmt"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"go.uber.org/zap"
)

type ChatService struct {
	mlClient clients.MLClient
}

func NewChatService(mlClient clients.MLClient) *ChatService {
	return &ChatService{
		mlClient: mlClient,
	}
}

func (s *ChatService) GenerateSuggestions(ctx context.Context, req *models.SuggestChatRequest) (*models.SuggestChatResponse, error) {
	internalReq := &models.InternalSuggestChatRequest{
		Message: req.Message,
		Profile: req.Profile,
		History: req.History,
	}
	resp, err := s.mlClient.SuggestChat(ctx, internalReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat response: %w", err)
	}
	zap.L().Debug("got chat response")
	return resp, nil
}

func (s *ChatService) ModifyRecipeViaChat(ctx context.Context, req *models.ModifyChatRequest) (*models.ModifyChatResponse, error) {
	internalReq := &models.InternalModifyChatRequest{
		Message: req.Message,
		Recipe:  req.Recipe,
		Profile: req.Profile,
	}
	resp, err := s.mlClient.ModifyChat(ctx, internalReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat response: %w", err)
	}
	zap.L().Debug("got chat response")
	return resp, nil
}

func (s *ChatService) AnswerCookingQuestion(ctx context.Context, req *models.GeneralChatRequest) (*models.GeneralChatResponse, error) {
	internalReq := &models.InternalGeneralChatRequest{
		Message: req.Message,
		Recipe:  req.Recipe,
		Profile: req.Profile,
	}
	resp, err := s.mlClient.GeneralChat(ctx, internalReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat response: %w", err)
	}
	zap.L().Debug("got chat response")
	return resp, nil
}
