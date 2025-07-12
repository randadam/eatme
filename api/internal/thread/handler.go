package thread

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/api"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ThreadHandler struct {
	threadService *ThreadService
}

func NewThreadHandler(threadService *ThreadService) *ThreadHandler {
	return &ThreadHandler{
		threadService: threadService,
	}
}

// @Summary Start a new suggestion thread
// @Description Start a new suggestion thread
// @ID startSuggestionThread
// @Tags thread
// @Accept json
// @Produce json
// @Param request body models.StartSuggestionThreadRequest true "Suggestion thread request"
// @Success 200 {object} models.ThreadState
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /thread/suggest [post]
func (h *ThreadHandler) StartSuggestionThread(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrBadRequest)
		return
	}

	var input models.StartSuggestionThreadRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		zap.L().Error("failed to decode start suggestion thread request", zap.Error(err))
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	threadState, err := h.threadService.StartSuggestionThread(r.Context(), userID, input.Prompt)
	if err != nil {
		zap.L().Error("failed to start suggestion thread", zap.Error(err))
		switch {
		case errors.Is(err, ErrThreadNotFound):
			api.ErrorJSON(w, http.StatusNotFound, models.ApiErrThreadNotFound)
		default:
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		}
		return
	}
	api.WriteJSON(w, http.StatusOK, threadState)
}

// @Summary Get new suggestions
// @Description Get new recipe suggestions
// @ID getNewSuggestions
// @Tags thread
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Param request body models.GetNewSuggestionsRequest true "Get new suggestions request"
// @Success 200 {object} []models.RecipeSuggestion
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Thread not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /thread/{threadId}/suggest [post]
func (h *ThreadHandler) GetNewSuggestions(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrBadRequest)
		return
	}

	threadID := chi.URLParam(r, "threadId")
	if threadID == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	var input models.GetNewSuggestionsRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		zap.L().Error("failed to decode get new suggestions request", zap.Error(err))
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	suggestions, err := h.threadService.GetNewSuggestions(r.Context(), userID, threadID, input)
	if err != nil {
		zap.L().Error("failed to get new suggestions", zap.Error(err))
		switch {
		case errors.Is(err, ErrThreadNotFound):
			api.ErrorJSON(w, http.StatusNotFound, models.ApiErrThreadNotFound)
		default:
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		}
		return
	}
	api.WriteJSON(w, http.StatusOK, suggestions)
}

// @Summary Accept a suggestion
// @Description Accept a suggestion
// @ID acceptSuggestion
// @Tags thread
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Param suggestionId path string true "Suggestion ID"
// @Success 200 {object} models.UserRecipe
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Thread not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /thread/{threadId}/accept/{suggestionId} [post]
func (h *ThreadHandler) AcceptSuggestion(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrBadRequest)
		return
	}

	threadID := chi.URLParam(r, "threadId")
	if threadID == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	suggestionID := chi.URLParam(r, "suggestionId")
	if suggestionID == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	recipe, err := h.threadService.AcceptSuggestion(r.Context(), userID, threadID, suggestionID)
	if err != nil {
		zap.L().Error("failed to accept suggestion", zap.Error(err))
		switch {
		case errors.Is(err, ErrThreadNotFound):
			api.ErrorJSON(w, http.StatusNotFound, models.ApiErrThreadNotFound)
		default:
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		}
		return
	}
	api.WriteJSON(w, http.StatusOK, recipe)
}

// @Summary Modify a recipe via chat
// @Description Modify a recipe via chat
// @ID modifyRecipe
// @Tags thread
// @Accept json
// @Produce json
// @Param recipeId path string true "Recipe ID"
// @Param request body models.ModifyRecipeViaChatRequest true "Modify recipe via chat request"
// @Success 200 {object} models.UserRecipe
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Thread not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /recipes/{recipeId}/modify/chat [post]
func (h *ThreadHandler) ModifyRecipeViaChat(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrBadRequest)
		return
	}

	recipeID := chi.URLParam(r, "recipeId")
	if recipeID == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	var input models.ModifyRecipeViaChatRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		zap.L().Error("failed to decode modify recipe via chat request", zap.Error(err))
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	recipe, err := h.threadService.ModifyRecipeViaChat(r.Context(), userID, recipeID, input.Prompt)
	if err != nil {
		zap.L().Error("failed to modify recipe via chat", zap.Error(err))
		switch {
		case errors.Is(err, ErrThreadNotFound):
			api.ErrorJSON(w, http.StatusNotFound, models.ApiErrThreadNotFound)
		default:
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		}
		return
	}
	api.WriteJSON(w, http.StatusOK, recipe)
}

// @Summary Answer a cooking question
// @Description Answer a cooking question
// @ID answerCookingQuestion
// @Tags thread
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Param request body models.AnswerCookingQuestionRequest true "Answer cooking question request"
// @Success 200 {object} models.AnswerCookingQuestionResponse
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /thread/{threadId}/question [post]
func (h *ThreadHandler) AnswerCookingQuestion(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrBadRequest)
		return
	}

	var input models.AnswerCookingQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		zap.L().Error("failed to decode answer cooking question request", zap.Error(err))
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	response, err := h.threadService.AnswerCookingQuestion(r.Context(), userID, input.Question)
	if err != nil {
		zap.L().Error("failed to answer cooking question", zap.Error(err))
		switch {
		case errors.Is(err, ErrThreadNotFound):
			api.ErrorJSON(w, http.StatusNotFound, models.ApiErrThreadNotFound)
		default:
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		}
		return
	}
	api.WriteJSON(w, http.StatusOK, response)
}

// @Summary Get a thread
// @Description Get a thread
// @ID getThread
// @Tags thread
// @Accept json
// @Produce json
// @Param threadId path string true "Thread ID"
// @Success 200 {object} models.ThreadState
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Thread not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /thread/{threadId} [get]
func (h *ThreadHandler) GetThread(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	threadID := chi.URLParam(r, "threadId")
	if threadID == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	threadState, err := h.threadService.GetThreadState(r.Context(), threadID)
	if err != nil {
		zap.L().Error("failed to get thread state", zap.Error(err))
		switch {
		case errors.Is(err, ErrThreadNotFound):
			api.ErrorJSON(w, http.StatusNotFound, models.ApiErrThreadNotFound)
		default:
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		}
		return
	}
	api.WriteJSON(w, http.StatusOK, threadState)
}
