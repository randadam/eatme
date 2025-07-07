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

func (s *RecipeService) NewSuggestionThread(ctx context.Context, userID string, thread models.SuggestionThread) error {
	return s.store.WithTx(func(tx db.Store) error {
		err := tx.CreateSuggestionThread(ctx, userID, thread)
		if err != nil {
			return err
		}
		if thread.Suggestions != nil {
			for _, suggestion := range thread.Suggestions {
				err = tx.AppendToSuggestionThread(ctx, thread.ID, suggestion)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *RecipeService) AppendToSuggestionThread(ctx context.Context, threadID string, suggestion models.RecipeSuggestion) error {
	return s.store.AppendToSuggestionThread(ctx, threadID, suggestion)
}

func (s *RecipeService) AcceptSuggestion(ctx context.Context, userID string, threadID string, suggestion models.RecipeSuggestion) (models.UserRecipe, error) {
	var newRecipe models.UserRecipe
	err := s.store.WithTx(func(tx db.Store) error {
		err := tx.AcceptSuggestion(ctx, threadID, suggestion)
		if err != nil {
			return err
		}

		newRecipe, err = createRecipe(ctx, tx, userID, suggestion.Suggestion)
		return err
	})
	return newRecipe, err
}

func (s *RecipeService) GetSuggestionThread(ctx context.Context, threadID string) (models.SuggestionThread, error) {
	return s.store.GetSuggestionThread(ctx, threadID)
}

func (s *RecipeService) NewRecipe(ctx context.Context, userID string, recipeBody models.RecipeBody) (models.UserRecipe, error) {
	var recipe models.UserRecipe
	err := s.store.WithTx(func(tx db.Store) error {
		var err error
		recipe, err = createRecipe(ctx, tx, userID, recipeBody)
		return err
	})
	return recipe, err
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
		return models.UserRecipe{}, err
	}
	rv := models.RecipeVersion{
		ID:           versionID,
		UserRecipeID: recipeID,
		RecipeBody:   body,
	}
	if err := st.AddRecipeVersion(ctx, rv); err != nil {
		return models.UserRecipe{}, err
	}
	return recipe, nil
}
