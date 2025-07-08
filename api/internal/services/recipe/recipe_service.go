package recipe

import (
	"context"
	"fmt"

	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RecipeService struct {
	store db.Store
}

func NewRecipeService(store db.Store) *RecipeService {
	return &RecipeService{store: store}
}

func (s *RecipeService) NewSuggestionThread(ctx context.Context, userID string, thread models.SuggestionThread) error {
	err := s.store.WithTx(func(tx db.Store) error {
		err := tx.CreateSuggestionThread(ctx, userID, thread)
		if err != nil {
			return fmt.Errorf("failed to create suggestion thread in new suggestion thread: %w", err)
		}
		zap.L().Debug("created suggestion thread")
		if thread.Suggestions != nil {
			for i, suggestion := range thread.Suggestions {
				err = tx.AppendToSuggestionThread(ctx, thread.ID, suggestion)
				if err != nil {
					return fmt.Errorf("failed to append to suggestion thread in new suggestion thread: %w", err)
				}
				zap.L().Debug("appended to suggestion thread", zap.Int("index", i))
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to create suggestion thread in new suggestion thread: %w", err)
	}
	return nil
}

func (s *RecipeService) AppendToSuggestionThread(ctx context.Context, threadID string, suggestion models.RecipeSuggestion) error {
	err := s.store.AppendToSuggestionThread(ctx, threadID, suggestion)
	if err != nil {
		return fmt.Errorf("failed to append to suggestion thread in append to suggestion thread: %w", err)
	}
	return nil
}

func (s *RecipeService) AcceptSuggestion(ctx context.Context, userID string, threadID string, suggestion models.RecipeSuggestion) (models.UserRecipe, error) {
	var newRecipe models.UserRecipe
	err := s.store.WithTx(func(tx db.Store) error {
		err := tx.AcceptSuggestion(ctx, threadID, suggestion)
		if err != nil {
			return fmt.Errorf("failed to accept suggestion in accept suggestion: %w", err)
		}
		zap.L().Debug("accepted suggestion")

		newRecipe, err = createRecipe(ctx, tx, userID, suggestion.Suggestion)
		if err != nil {
			return fmt.Errorf("failed to create recipe in accept suggestion: %w", err)
		}
		zap.L().Debug("created recipe")
		return nil
	})
	return newRecipe, err
}

func (s *RecipeService) GetSuggestionThread(ctx context.Context, threadID string) (models.SuggestionThread, error) {
	thread, err := s.store.GetSuggestionThread(ctx, threadID)
	if err != nil {
		return models.SuggestionThread{}, fmt.Errorf("failed to get suggestion thread in get suggestion thread: %w", err)
	}
	return thread, nil
}

func (s *RecipeService) NewRecipe(ctx context.Context, userID string, recipeBody models.RecipeBody) (models.UserRecipe, error) {
	var recipe models.UserRecipe
	err := s.store.WithTx(func(tx db.Store) error {
		var err error
		recipe, err = createRecipe(ctx, tx, userID, recipeBody)
		if err != nil {
			return fmt.Errorf("failed to create recipe in new recipe: %w", err)
		}
		zap.L().Debug("created recipe")
		return nil
	})
	return recipe, err
}

func (s *RecipeService) UpdateRecipe(ctx context.Context, userID string, recipeID string, recipeBody models.RecipeBody) error {
	err := s.store.WithTx(func(tx db.Store) error {
		current, err := tx.GetUserRecipe(ctx, userID, recipeID)
		if err != nil {
			return fmt.Errorf("failed to get user recipe in update recipe: %w", err)
		}
		zap.L().Debug("got user recipe")

		recipeVersion := models.RecipeVersion{
			ID:           uuid.New().String(),
			UserRecipeID: recipeID,
			ParentID:     &current.LatestVersionID,
			RecipeBody:   recipeBody,
		}

		err = tx.AddRecipeVersion(ctx, recipeVersion)
		if err != nil {
			return fmt.Errorf("failed to add recipe version in update recipe: %w", err)
		}
		zap.L().Debug("added recipe version")
		err = tx.UpdateUserRecipeVersion(ctx, userID, recipeID, recipeVersion)
		if err != nil {
			return fmt.Errorf("failed to update user recipe version in update recipe: %w", err)
		}
		zap.L().Debug("updated user recipe version")
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update recipe in update recipe: %w", err)
	}
	zap.L().Debug("updated recipe")
	return nil
}

func (s *RecipeService) GetUserRecipe(ctx context.Context, userID string, recipeID string) (models.UserRecipe, error) {
	recipe, err := s.store.GetUserRecipe(ctx, userID, recipeID)
	if err != nil {
		return models.UserRecipe{}, fmt.Errorf("failed to get user recipe in get user recipe: %w", err)
	}
	return recipe, nil
}

func (s *RecipeService) GetAllUserRecipes(ctx context.Context, userID string) ([]models.UserRecipe, error) {
	recipes, err := s.store.GetAllUserRecipes(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user recipes in get all user recipes: %w", err)
	}
	return recipes, nil
}

func createRecipe(ctx context.Context, st db.Store, userID string,
	body models.RecipeBody) (models.UserRecipe, error) {

	recipeID := uuid.New().String()
	versionID := uuid.New().String()

	recipe := models.UserRecipe{
		ID:              recipeID,
		LatestVersionID: versionID,
		UserID:          userID,
		RecipeBody:      body,
	}
	if err := st.SaveUserRecipe(ctx, recipe); err != nil {
		return models.UserRecipe{}, fmt.Errorf("failed to save user recipe in create recipe: %w", err)
	}
	zap.L().Debug("saved user recipe")
	rv := models.RecipeVersion{
		ID:           versionID,
		UserRecipeID: recipeID,
		RecipeBody:   body,
	}
	if err := st.AddRecipeVersion(ctx, rv); err != nil {
		return models.UserRecipe{}, fmt.Errorf("failed to add recipe version in create recipe: %w", err)
	}
	zap.L().Debug("added recipe version")
	return recipe, nil
}
