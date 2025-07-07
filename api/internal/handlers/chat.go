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
	"github.com/go-chi/chi/v5"
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
// @Router /chat/plan/:planId/recipe [post]
func (h *ChatHandler) SuggestChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlanId := chi.URLParam(r, "mealPlanId")
	if mealPlanId == "" {
		errorJSON(w, errors.New("missing meal plan ID"), http.StatusBadRequest)
		return
	}

	requestBody := models.SuggestChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	profile, err := h.userService.GetProfile(userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	mealPlan, err := h.mealService.GetMealPlan(userID, mealPlanId)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			errorJSON(w, errors.New("meal plan not found"), http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	req := &models.InternalSuggestChatRequest{
		Message: requestBody.Message,
		Profile: profile,
	}
	resp, err := h.mlClient.SuggestChat(ctx, req)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	mealPlan.Recipes = append(mealPlan.Recipes, *resp.NewRecipe)
	if err := h.mealService.SaveMealPlan(userID, mealPlan); err != nil {
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
// @Param request body models.ModifyChatRequest true "Modify chat request"
// @Success 200 {object} models.ModifyChatResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/plan/:planId/recipe/:recipeId [post]
func (h *ChatHandler) ModifyChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlanId := chi.URLParam(r, "mealPlanId")
	if mealPlanId == "" {
		errorJSON(w, errors.New("missing meal plan ID"), http.StatusBadRequest)
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

	profile, err := h.userService.GetProfile(userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	mealPlan, err := h.mealService.GetMealPlan(userID, mealPlanId)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			errorJSON(w, errors.New("meal plan not found"), http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	currentRecipe := &models.Recipe{}
	found := false
	for _, recipe := range mealPlan.Recipes {
		if recipe.ID == recipeId {
			currentRecipe = &recipe
			found = true
			break
		}
	}
	if !found {
		errorJSON(w, errors.New("recipe not found"), http.StatusNotFound)
		return
	}

	ctx := context.Background()
	req := &models.InternalModifyChatRequest{
		Message: requestBody.Message,
		Recipe:  *currentRecipe,
		Profile: profile,
	}
	resp, err := h.mlClient.ModifyChat(ctx, req)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	*currentRecipe = *resp.NewRecipe
	if err := h.mealService.SaveMealPlan(userID, mealPlan); err != nil {
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
// @Param request body models.GeneralChatRequest true "General chat request"
// @Success 200 {object} models.GeneralChatResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /chat/plan/:planId [post]
func (h *ChatHandler) GeneralChat(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlanId := chi.URLParam(r, "mealPlanId")
	if mealPlanId == "" {
		errorJSON(w, errors.New("missing meal plan ID"), http.StatusBadRequest)
		return
	}

	requestBody := models.GeneralChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	profile, err := h.userService.GetProfile(userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	mealPlan, err := h.mealService.GetMealPlan(userID, mealPlanId)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			errorJSON(w, errors.New("meal plan not found"), http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	req := &models.InternalGeneralChatRequest{
		Message:  requestBody.Message,
		MealPlan: mealPlan,
		Profile:  profile,
	}
	resp, err := h.mlClient.GeneralChat(ctx, req)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}
