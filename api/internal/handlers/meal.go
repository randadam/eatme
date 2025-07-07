package handlers

import (
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/services/meal"
	"github.com/go-chi/chi/v5"
)

type MealHandler struct {
	mealService *meal.MealService
}

func NewMealHandler(mealService *meal.MealService) *MealHandler {
	return &MealHandler{
		mealService: mealService,
	}
}

// @Summary Create new meal plan
// @Description Create new meal plan
// @ID createMealPlan
// @Tags Meal
// @Accept json
// @Produce json
// @Success 200 {object} models.MealPlan
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /meal/plan [post]
func (h *MealHandler) CreateMealPlan(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlan, err := h.mealService.NewMealPlan(userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, mealPlan)
}

// @Summary Get meal plan by ID
// @Description Get meal plan by ID
// @ID getMealPlan
// @Tags Meal
// @Accept json
// @Produce json
// @Param meal_plan_id path string true "Meal plan ID"
// @Success 200 {object} models.MealPlan
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /meal/plan/{meal_plan_id} [get]
func (h *MealHandler) GetMealPlan(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlanId := chi.URLParam(r, "meal_plan_id")
	if mealPlanId == "" {
		errorJSON(w, errors.New("missing meal plan ID"), http.StatusBadRequest)
		return
	}

	mealPlan, err := h.mealService.GetMealPlan(userID, mealPlanId)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, mealPlan)
}

// @Summary Get all meal plans for user
// @Description Get all meal plans for user
// @ID getAllMealPlans
// @Tags Meal
// @Accept json
// @Produce json
// @Success 200 {array} models.MealPlan
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /meal/plans [get]
func (h *MealHandler) GetAllMealPlans(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	mealPlans, err := h.mealService.GetAllPlans(userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, mealPlans)
}
