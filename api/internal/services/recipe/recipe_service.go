package recipe

import (
	"context"

	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/models"
	"github.com/google/uuid"
)

type RecipeService struct {
	store db.Store
}

func NewRecipeService(store db.Store) *RecipeService {
	return &RecipeService{store: store}
}

func (s *RecipeService) NewRecipe(ctx context.Context, userID string, recipeBody models.RecipeBody) (models.UserRecipe, error) {
	recipeId := uuid.New().String()
	versionId := uuid.New().String()

	recipe := models.UserRecipe{
		ID:              recipeId,
		LatestVersionID: versionId,
		UserID:          userID,
		RecipeBody:      recipeBody,
	}

	err := s.store.WithTx(func(tx db.Store) error {
		err := tx.SaveUserRecipe(ctx, recipe)
		if err != nil {
			return err
		}

		recipeVersion := models.RecipeVersion{
			ID:           versionId,
			UserRecipeID: recipeId,
			RecipeBody:   recipeBody,
		}
		return tx.AddRecipeVersion(ctx, recipeVersion)
	})

	if err != nil {
		return models.UserRecipe{}, err
	}

	return recipe, nil
}

func (s *RecipeService) UpdateRecipe(ctx context.Context, userID string, recipeID string, recipeBody models.RecipeBody) error {
	return s.store.WithTx(func(tx db.Store) error {
		current, err := tx.GetUserRecipe(ctx, userID, recipeID)
		if err != nil {
			return err
		}

		recipeVersion := models.RecipeVersion{
			ID:           uuid.New().String(),
			UserRecipeID: recipeID,
			ParentID:     &current.LatestVersionID,
			RecipeBody:   recipeBody,
		}

		err = tx.AddRecipeVersion(ctx, recipeVersion)
		if err != nil {
			return err
		}
		return tx.UpdateUserRecipeVersion(ctx, userID, recipeID, recipeVersion.ID)
	})
}

func (s *RecipeService) GetUserRecipe(ctx context.Context, userID string, recipeID string) (models.UserRecipe, error) {
	return s.store.GetUserRecipe(ctx, userID, recipeID)
}

func (s *RecipeService) GetAllUserRecipes(ctx context.Context, userID string) ([]models.UserRecipe, error) {
	return s.store.GetAllUserRecipes(ctx, userID)
}
