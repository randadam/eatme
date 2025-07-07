package chat

import (
	"context"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/models"
)

type ChatService struct {
	mlClient *clients.MLClient
}

func NewChatService(mlClient *clients.MLClient) *ChatService {
	return &ChatService{
		mlClient: mlClient,
	}
}

func (s *ChatService) SuggestChat(ctx context.Context, req *models.InternalSuggestChatRequest) (*models.SuggestChatResponse, error) {
	return s.mlClient.SuggestChat(ctx, req)
}

func (s *ChatService) ModifyChat(ctx context.Context, req *models.InternalModifyChatRequest) (*models.ModifyChatResponse, error) {
	return s.mlClient.ModifyChat(ctx, req)
}

func (s *ChatService) GeneralChat(ctx context.Context, req *models.InternalGeneralChatRequest) (*models.GeneralChatResponse, error) {
	return s.mlClient.GeneralChat(ctx, req)
}
