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

func (s *ChatService) Chat(ctx context.Context, req *models.ChatRequest) (*models.ChatResponse, error) {
	return s.mlClient.Chat(ctx, req)
}
