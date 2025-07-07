package recipe

import (
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

func (s *RecipeService) NewRecipe(userID string, recipeBody models.RecipeBody) (models.UserRecipe, error) {
	recipeId := uuid.New().String()
	versionId := uuid.New().String()

	recipe := models.UserRecipe{
		ID:              recipeId,
		LatestVersionID: versionId,
		UserID:          userID,
		RecipeBody:      recipeBody,
	}
	err := s.store.SaveUserRecipe(recipe)
	if err != nil {
		return models.UserRecipe{}, err
	}

	recipeVersion := models.RecipeVersion{
		ID:           versionId,
		UserRecipeID: recipeId,
		RecipeBody:   recipeBody,
	}
	err = s.store.AddRecipeVersion(recipeVersion)
	if err != nil {
		return models.UserRecipe{}, err
	}
	return recipe, nil
}

func (s *RecipeService) UpdateRecipe(userID string, recipeID string, recipeBody models.RecipeBody) error {
	current, err := s.store.GetUserRecipe(userID, recipeID)
	if err != nil {
		return err
	}

	recipeVersion := models.RecipeVersion{
		ID:           uuid.New().String(),
		UserRecipeID: recipeID,
		ParentID:     &current.LatestVersionID,
		RecipeBody:   recipeBody,
	}

	err = s.store.AddRecipeVersion(recipeVersion)
	if err != nil {
		return err
	}
	err = s.store.UpdateUserRecipeVersion(userID, recipeID, recipeVersion.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipeService) GetUserRecipe(userID string, recipeID string) (models.UserRecipe, error) {
	return s.store.GetUserRecipe(userID, recipeID)
}

func (s *RecipeService) GetAllUserRecipes(userID string) ([]models.UserRecipe, error) {
	return s.store.GetAllUserRecipes(userID)
}
