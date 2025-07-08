package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/ajohnston1219/eatme/api/internal/services/chat"
	"github.com/ajohnston1219/eatme/api/internal/services/recipe"
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /chat/suggest [post]
func (h *ChatHandler) StartSuggestChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	requestBody := models.SuggestChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	resp, err := h.chatService.StartSuggestChat(r.Context(), userID, &requestBody)
	if err != nil {
		switch {
		case errors.Is(err, recipe.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Suggestion thread not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /chat/suggest/{threadId}/next [post]
func (h *ChatHandler) NextRecipeSuggestion(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	threadId := chi.URLParam(r, "threadId")
	if threadId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	resp, err := h.chatService.GetNextSuggestion(r.Context(), userID, threadId)
	if err != nil {
		switch {
		case errors.Is(err, recipe.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		case errors.Is(err, recipe.ErrSuggestionThreadNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrSuggestionThreadNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Suggestion thread not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /chat/suggest/{threadId}/accept/{suggestionId} [post]
func (h *ChatHandler) AcceptRecipeSuggestion(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	threadId := chi.URLParam(r, "threadId")
	if threadId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	suggestionId := chi.URLParam(r, "suggestionId")
	if suggestionId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	resp, err := h.chatService.AcceptRecipeSuggestion(r.Context(), userID, threadId, suggestionId)
	if err != nil {
		switch {
		case errors.Is(err, recipe.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		case errors.Is(err, recipe.ErrSuggestionThreadNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrSuggestionThreadNotFound)
		case errors.Is(err, recipe.ErrSuggestionNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrSuggestionNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Suggestion thread not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /chat/suggest/{threadId} [get]
func (h *ChatHandler) GetSuggestionThread(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	threadId := chi.URLParam(r, "threadId")
	if threadId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	resp, err := h.chatService.GetSuggestionThread(r.Context(), userID, threadId)
	if err != nil {
		switch {
		case errors.Is(err, recipe.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		case errors.Is(err, recipe.ErrSuggestionThreadNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrSuggestionThreadNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Recipe not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /chat/modify/recipes/{recipeId} [put]
func (h *ChatHandler) ModifyRecipeChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipeId")
	if recipeId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	requestBody := models.ModifyChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	resp, err := h.chatService.ModifyRecipeChat(r.Context(), userID, recipeId, &requestBody)
	if err != nil {
		switch {
		case errors.Is(err, recipe.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Recipe not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /chat/question/recipes/{recipeId} [post]
func (h *ChatHandler) GeneralRecipeChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipeId")
	if recipeId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	requestBody := models.GeneralChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	resp, err := h.chatService.GeneralRecipeChat(r.Context(), userID, recipeId, &requestBody)
	if err != nil {
		switch {
		case errors.Is(err, recipe.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
