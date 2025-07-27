package recipe

import (
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/api"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/ajohnston1219/eatme/api/internal/utils/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
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
// @Param recipeId path string true "Recipe ID"
// @Success 200 {object} models.UserRecipe
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Recipe not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /recipes/{recipeId} [get]
func (h *RecipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipeId")
	if recipeId == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	var recipe *models.UserRecipe
	err := h.recipeService.db.WithTx(func(tx db.Store) error {
		var err error
		ctx := db.ContextWithTx(r.Context(), tx)
		recipe, err = h.recipeService.GetUserRecipe(ctx, userID, recipeId)
		if err != nil {
			logger.Logger(r.Context()).Error("failed to get user recipe", zap.Error(err))
			return err
		}
		return nil
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
	api.WriteJSON(w, http.StatusOK, recipe)
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

	var recipes []models.UserRecipe
	err := h.recipeService.db.WithTx(func(tx db.Store) error {
		var err error
		ctx := db.ContextWithTx(r.Context(), tx)
		recipes, err = h.recipeService.GetAllUserRecipes(ctx, userID)
		if err != nil {
			logger.Logger(r.Context()).Error("failed to get all user recipes", zap.Error(err))
			return err
		}
		return nil
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
	api.WriteJSON(w, http.StatusOK, recipes)
}

// @Summary Delete recipe
// @Description Delete recipe
// @ID deleteRecipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param recipeId path string true "Recipe ID"
// @Success 200 {object} models.UserRecipe
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Recipe not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /recipes/{recipeId} [delete]
func (h *RecipeHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	recipeId := chi.URLParam(r, "recipeId")
	if recipeId == "" {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	err := h.recipeService.db.WithTx(func(tx db.Store) error {
		ctx := db.ContextWithTx(r.Context(), tx)
		return h.recipeService.DeleteUserRecipe(ctx, userID, recipeId)
	})
	if err != nil {
		logger.Logger(r.Context()).Error("failed to delete user recipe", zap.Error(err))
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
