package recipe

import (
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/api"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/go-chi/chi/v5"
)

type RecipeHandler struct {
	recipeService *RecipeService
}

func NewRecipeHandler(recipeService *RecipeService) *RecipeHandler {
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
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipe_id")
	if recipeId == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	err := h.recipeService.db.WithTx(func(tx db.Store) error {
		ctx := db.ContextWithTx(r.Context(), tx)
		recipe, err := h.recipeService.GetUserRecipe(ctx, userID, recipeId)
		if err != nil {
			switch {
			case errors.Is(err, ErrRecipeNotFound):
				api.ErrorJSON(w, http.StatusNotFound, models.ApiErrRecipeNotFound)
			default:
				api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
			}
			return nil
		}
		api.WriteJSON(w, http.StatusOK, recipe)
		return nil
	})
	if err != nil {
		api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		return
	}
}

// @Summary Get all recipes for user
// @Description Get all recipes for user
// @ID getAllRecipes
// @Tags Recipe
// @Accept json
// @Produce json
// @Success 200 {array}  models.UserRecipe
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /recipes [get]
func (h *RecipeHandler) GetAllRecipes(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	err := h.recipeService.db.WithTx(func(tx db.Store) error {
		ctx := db.ContextWithTx(r.Context(), tx)
		recipes, err := h.recipeService.GetAllUserRecipes(ctx, userID)
		if err != nil {
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
			return nil
		}
		api.WriteJSON(w, http.StatusOK, recipes)
		return nil
	})
	if err != nil {
		api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		return
	}
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
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipe_id")
	if recipeId == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	err := h.recipeService.db.WithTx(func(tx db.Store) error {
		ctx := db.ContextWithTx(r.Context(), tx)
		return h.recipeService.DeleteUserRecipe(ctx, userID, recipeId)
	})
	if err != nil {
		switch {
		case errors.Is(err, ErrRecipeNotFound):
			api.ErrorJSON(w, http.StatusNotFound, models.ApiErrRecipeNotFound)
		default:
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		}
		return
	}
	api.WriteJSON(w, http.StatusOK, nil)
}
