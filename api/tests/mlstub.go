package tests

import (
	"context"

	"github.com/ajohnston1219/eatme/api/models"
)

type MLStub struct {
	Responses []models.SuggestChatResponse
	call      int
}

func (m *MLStub) SuggestChat(_ context.Context, _ *models.InternalSuggestChatRequest) (*models.SuggestChatResponse, error) {
	resp := m.Responses[m.call]
	m.call++
	return &models.SuggestChatResponse{
		ResponseText: resp.ResponseText,
		NewRecipe:    resp.NewRecipe,
	}, nil
}

func (m *MLStub) ModifyChat(_ context.Context, _ *models.InternalModifyChatRequest) (*models.ModifyChatResponse, error) {
	return nil, nil
}

func (m *MLStub) GeneralChat(_ context.Context, _ *models.InternalGeneralChatRequest) (*models.GeneralChatResponse, error) {
	return nil, nil
}
