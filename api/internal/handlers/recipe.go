package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/services/recipe"
	"github.com/go-chi/chi/v5"
)

type RecipeHandler struct {
	recipeService *recipe.RecipeService
}

func NewRecipeHandler(recipeService *recipe.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		recipeService: recipeService,
	}
}

// @Summary Get meal plan by ID
// @Description Get meal plan by ID
// @ID getMealPlan
// @Tags Recipe
// @Accept json
// @Produce json
// @Param meal_plan_id path string true "Meal plan ID"
// @Success 200 {object} models.MealPlan
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
func (h *RecipeHandler) GetMealPlan(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlanId := chi.URLParam(r, "meal_plan_id")
	if mealPlanId == "" {
		errorJSON(w, errors.New("missing meal plan ID"), http.StatusBadRequest)
		return
	}

	mealPlan, err := h.recipeService.GetMealPlan(userID, mealPlanId)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(mealPlan)
}
