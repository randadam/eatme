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

// @Summary Handle recipe suggestion chat request
// @Description Handle recipe suggestion chat request
// @ID suggestRecipe
// @Tags Chat
// @Accept json
// @Produce json
// @Param request body models.SuggestChatRequest true "Suggest chat request"
// @Success 200 {object} models.SuggestChatResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/recipes [post]
func (h *ChatHandler) SuggestChat(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.chatService.SuggestChat(r.Context(), userID, &requestBody)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// @Summary Handle recipe modification chat request
// @Description Handle recipe modification chat request
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
