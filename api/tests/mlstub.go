package tests

import (
	"context"

	"github.com/ajohnston1219/eatme/api/internal/models"
)

type MLStub struct {
	SuggestResponses []models.SuggestChatResponse
	suggestCall      int
	ModifyResponses  []models.ModifyChatResponse
	modifyCall       int
}

func (m *MLStub) SuggestChat(_ context.Context, _ *models.InternalSuggestChatRequest) (*models.SuggestChatResponse, error) {
	resp := m.SuggestResponses[m.suggestCall]
	m.suggestCall++
	return &resp, nil
}

func (m *MLStub) ModifyChat(_ context.Context, _ *models.InternalModifyChatRequest) (*models.ModifyChatResponse, error) {
	resp := m.ModifyResponses[m.modifyCall]
	m.modifyCall++
	return &resp, nil
}

func (m *MLStub) GeneralChat(_ context.Context, _ *models.InternalGeneralChatRequest) (*models.GeneralChatResponse, error) {
	return nil, nil
}
