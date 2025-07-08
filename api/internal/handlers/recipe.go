package handlers

import (
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/models"
	recipeService "github.com/ajohnston1219/eatme/api/internal/services/recipe"
	"github.com/go-chi/chi/v5"
)

type RecipeHandler struct {
	recipeService *recipeService.RecipeService
}

func NewRecipeHandler(recipeService *recipeService.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		recipeService: recipeService,
	}
}

// @Summary Get recipe by ID
// @Description Get recipe by ID
// @ID getRecipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipe_id path string true "Recipe ID"
// @Success 200 {object} models.UserRecipe
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Recipe not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /recipes/{recipe_id} [get]
func (h *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipe_id")
	if recipeId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	recipe, err := h.recipeService.GetUserRecipe(r.Context(), userID, recipeId)
	if err != nil {
		switch {
		case errors.Is(err, recipeService.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
		return
	}
	writeJSON(w, http.StatusOK, recipe)
}

// @Summary Get all recipes for user
// @Description Get all recipes for user
// @ID getAllRecipes
// @Tags Recipe
// @Accept json
// @Produce json
// @Success 200 {array} models.UserRecipe
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /recipes [get]
func (h *RecipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	recipes, err := h.recipeService.GetAllUserRecipes(r.Context(), userID)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		return
	}
	writeJSON(w, http.StatusOK, recipes)
}

// @Summary Delete recipe
// @Description Delete recipe
// @ID deleteRecipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipe_id path string true "Recipe ID"
// @Success 200 {object} models.UserRecipe
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Recipe not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /recipes/{recipe_id} [delete]
func (h *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipe_id")
	if recipeId == "" {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	err := h.recipeService.DeleteUserRecipe(r.Context(), userID, recipeId)
	if err != nil {
		switch {
		case errors.Is(err, recipeService.ErrRecipeNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrRecipeNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
		return
	}
	writeJSON(w, http.StatusOK, nil)
}
