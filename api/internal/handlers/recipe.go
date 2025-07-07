package handlers

import (
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

// @Summary Get recipe by ID
// @Description Get recipe by ID
// @ID getRecipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipe_id path string true "Recipe ID"
// @Success 200 {object} models.UserRecipe
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /recipes/{recipe_id} [get]
func (h *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipe_id")
	if recipeId == "" {
		errorJSON(w, errors.New("missing recipe ID"), http.StatusBadRequest)
		return
	}

	recipe, err := h.recipeService.GetUserRecipe(r.Context(), userID, recipeId)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
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
// @Failure 400 {object} models.BadRequestResponse
// @Failure 500 {object} models.InternalServerErrorResponse
// @Router /recipes [get]
func (h *RecipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	recipes, err := h.recipeService.GetAllUserRecipes(r.Context(), userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, recipes)
}
