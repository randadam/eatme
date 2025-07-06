package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/services/recipe"
	"github.com/ajohnston1219/eatme/api/internal/services/user"
	"github.com/ajohnston1219/eatme/api/models"
)

type ChatHandler struct {
	mlClient      *clients.MLClient
	userService   *user.UserService
	recipeService *recipe.RecipeService
}

func NewChatHandler(mlClient *clients.MLClient, userService *user.UserService, recipeService *recipe.RecipeService) *ChatHandler {
	return &ChatHandler{
		mlClient:      mlClient,
		userService:   userService,
		recipeService: recipeService,
	}
}

// @Summary Handle chat request
// @Description Handle chat request
// @ID chat
// @Tags Chat
// @Accept json
// @Produce json
// @Param meal_plan_id query string true "Meal plan ID"
// @Param message query string true "Message"
// @Success 200 {object} models.ChatResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /chat [post]
func (h *ChatHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlanId := r.URL.Query().Get("meal_plan_id")
	if mealPlanId == "" {
		errorJSON(w, errors.New("missing meal plan ID"), http.StatusBadRequest)
		return
	}

	profile, err := h.userService.GetProfile(userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	mealPlan, err := h.recipeService.GetMealPlan(userID, mealPlanId)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	req := &models.ChatRequest{
		UserID:   userID,
		Message:  r.URL.Query().Get("message"),
		MealPlan: mealPlan,
		Profile:  profile,
	}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.mlClient.Chat(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
