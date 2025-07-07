package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/services/chat"
	"github.com/ajohnston1219/eatme/api/models"
	"github.com/go-chi/chi/v5"
)

type ChatHandler struct {
	chatService *chat.ChatService
}

func NewChatHandler(chatService *chat.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

// @Summary Handle starting a recipe suggestion chat
// @Description Handle starting a recipe suggestion chat
// @ID suggestRecipe
// @Tags Chat
// @Accept json
// @Produce json
// @Param request body models.SuggestChatRequest true "Suggest chat request"
// @Success 200 {object} models.SuggestChatResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/suggest [post]
func (h *ChatHandler) StartSuggestChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	requestBody := models.SuggestChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp, err := h.chatService.StartSuggestChat(r.Context(), userID, &requestBody)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// @Summary Handle getting next recipe suggestion
// @Description Handle getting next recipe suggestion
// @ID nextRecipeSuggestion
// @Tags Chat
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} models.RecipeSuggestion
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/suggest/{threadId}/next [post]
func (h *ChatHandler) NextRecipeSuggestion(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	threadId := chi.URLParam(r, "threadId")
	if threadId == "" {
		errorJSON(w, errors.New("missing thread ID"), http.StatusBadRequest)
		return
	}

	resp, err := h.chatService.GetNextSuggestion(r.Context(), userID, threadId)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// @Summary Handle accepting a recipe suggestion
// @Description Handle accepting a recipe suggestion
// @ID acceptRecipeSuggestion
// @Tags Chat
// @Produce json
// @Param threadId path string true "Thread ID"
// @Param suggestionId path string true "Suggestion ID"
// @Success 200 {object} models.UserRecipe
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/suggest/{threadId}/accept/{suggestionId} [post]
func (h *ChatHandler) AcceptRecipeSuggestion(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	threadId := chi.URLParam(r, "threadId")
	if threadId == "" {
		errorJSON(w, errors.New("missing thread ID"), http.StatusBadRequest)
		return
	}

	suggestionId := chi.URLParam(r, "suggestionId")
	if suggestionId == "" {
		errorJSON(w, errors.New("missing suggestion ID"), http.StatusBadRequest)
		return
	}

	resp, err := h.chatService.AcceptRecipeSuggestion(r.Context(), userID, threadId, suggestionId)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// @Summary Get suggestion thread
// @Description Get suggestion thread
// @ID getSuggestionThread
// @Tags Chat
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} models.SuggestionThread
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/suggest/{threadId} [get]
func (h *ChatHandler) GetSuggestionThread(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	threadId := chi.URLParam(r, "threadId")
	if threadId == "" {
		errorJSON(w, errors.New("missing thread ID"), http.StatusBadRequest)
		return
	}

	resp, err := h.chatService.GetSuggestionThread(r.Context(), userID, threadId)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// @Summary Handle modifying a recipe
// @Description Handle modifying a recipe
// @ID modifyRecipe
// @Tags Chat
// @Accept json
// @Produce json
// @Param recipeId path string true "Recipe ID"
// @Param request body models.ModifyChatRequest true "Modify chat request"
// @Success 200 {object} models.ModifyChatResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/recipes/{recipeId} [put]
func (h *ChatHandler) ModifyChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipeId")
	if recipeId == "" {
		errorJSON(w, errors.New("missing recipe ID"), http.StatusBadRequest)
		return
	}

	requestBody := models.ModifyChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp, err := h.chatService.ModifyChat(r.Context(), userID, recipeId, &requestBody)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// @Summary Handle general chat request
// @Description Handle general chat request
// @ID generalChat
// @Tags Chat
// @Accept json
// @Produce json
// @Param recipeId path string true "Recipe ID"
// @Param request body models.GeneralChatRequest true "General chat request"
// @Success 200 {object} models.GeneralChatResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/recipes/{recipeId}/question [post]
func (h *ChatHandler) GeneralChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipeId")
	if recipeId == "" {
		errorJSON(w, errors.New("missing recipe ID"), http.StatusBadRequest)
		return
	}

	requestBody := models.GeneralChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp, err := h.chatService.GeneralChat(r.Context(), userID, recipeId, &requestBody)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
