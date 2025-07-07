package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/services/meal"
	"github.com/ajohnston1219/eatme/api/internal/services/user"
	"github.com/ajohnston1219/eatme/api/models"
)

type ChatHandler struct {
	mlClient    *clients.MLClient
	userService *user.UserService
	mealService *meal.MealService
}

func NewChatHandler(mlClient *clients.MLClient, userService *user.UserService, mealService *meal.MealService) *ChatHandler {
	return &ChatHandler{
		mlClient:    mlClient,
		userService: userService,
		mealService: mealService,
	}
}

// @Summary Handle chat request
// @Description Handle chat request
// @ID chat
// @Tags Chat
// @Accept json
// @Produce json
// @Param request body models.ChatRequest true "Chat request"
// @Success 200 {object} models.ChatResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat [post]
func (h *ChatHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	requestBody := models.ChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	profile, err := h.userService.GetProfile(userID)

	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	mealPlan, err := h.mealService.GetMealPlan(userID, requestBody.MealPlanID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			errorJSON(w, errors.New("meal plan not found"), http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	req := &models.InternalChatRequest{
		UserID:   userID,
		Message:  requestBody.Message,
		MealPlan: mealPlan,
		Profile:  profile,
	}
	resp, err := h.mlClient.Chat(ctx, req)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if resp.NewMealPlan != nil {
		resp.NewMealPlan.ID = requestBody.MealPlanID
		err = h.mealService.SaveMealPlan(userID, *resp.NewMealPlan)
		if err != nil {
			errorJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	writeJSON(w, http.StatusOK, resp)
}
