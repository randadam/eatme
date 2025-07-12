package recipe

import (
	"context"
	"errors"
	"fmt"

	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RecipeService struct {
	db db.Store
}

func NewRecipeService(db db.Store) *RecipeService {
	return &RecipeService{
		db: db,
	}
}

func (s *RecipeService) getStore(ctx context.Context) db.Store {
	if tx, ok := db.GetTx(ctx); ok {
		return tx
	}
	return s.db
}

func (s *RecipeService) NewRecipe(ctx context.Context, userID string, threadID string, recipeBody models.RecipeBody) (*models.UserRecipe, error) {
	store := s.getStore(ctx)
	recipeID := uuid.New().String()
	versionID := uuid.New().String()

	recipe := models.UserRecipe{
		ID:              recipeID,
		LatestVersionID: versionID,
		UserID:          userID,
		RecipeBody:      recipeBody,
		ThreadID:        threadID,
	}
	if err := store.SaveUserRecipe(ctx, recipe); err != nil {
		return nil, fmt.Errorf("failed to save user recipe: %w", err)
	}
	zap.L().Debug("saved user recipe")
	rv := models.RecipeVersion{
		ID:           versionID,
		UserRecipeID: recipeID,
		RecipeBody:   recipeBody,
	}
	if err := store.AddRecipeVersion(ctx, rv); err != nil {
		return nil, fmt.Errorf("failed to add recipe version: %w", err)
	}
	zap.L().Debug("added recipe version")
	return &recipe, nil
}

func (s *RecipeService) UpdateRecipe(ctx context.Context, userID string, recipeID string, recipeBody models.RecipeBody) error {
	store := s.getStore(ctx)
	current, err := store.GetUserRecipe(ctx, userID, recipeID)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrNotFound):
			return ErrRecipeNotFound
		default:
			return fmt.Errorf("failed to get user recipe: %w", err)
		}
	}
	zap.L().Debug("got user recipe")

	recipeVersion := models.RecipeVersion{
		ID:           uuid.New().String(),
		UserRecipeID: recipeID,
		ParentID:     &current.LatestVersionID,
		RecipeBody:   recipeBody,
	}

	err = store.AddRecipeVersion(ctx, recipeVersion)
	if err != nil {
		return fmt.Errorf("failed to add recipe version: %w", err)
	}
	zap.L().Debug("added recipe version")
	err = store.UpdateUserRecipeVersion(ctx, userID, recipeID, recipeVersion)
	if err != nil {
		return fmt.Errorf("failed to update user recipe version: %w", err)
	}
	zap.L().Debug("updated user recipe version")
	return nil
}

func (s *RecipeService) GetUserRecipe(ctx context.Context, userID string, recipeID string) (*models.UserRecipe, error) {
	store := s.getStore(ctx)
	recipe, err := store.GetUserRecipe(ctx, userID, recipeID)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrNotFound):
			return nil, ErrRecipeNotFound
		default:
			return nil, fmt.Errorf("failed to get user recipe: %w", err)
		}
	}
	return &recipe, nil
}

func (s *RecipeService) GetAllUserRecipes(ctx context.Context, userID string) ([]models.UserRecipe, error) {
	store := s.getStore(ctx)
	recipes, err := store.GetAllUserRecipes(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user recipes: %w", err)
	}
	return recipes, nil
}

func (s *RecipeService) DeleteUserRecipe(ctx context.Context, userID string, recipeID string) error {
	store := s.getStore(ctx)
	err := store.DeleteUserRecipe(ctx, userID, recipeID)
	if err != nil {
		return fmt.Errorf("failed to delete user recipe: %w", err)
	}
	return nil
}
