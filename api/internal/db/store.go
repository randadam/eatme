package db

import "github.com/ajohnston1219/eatme/api/models"

type Store interface {
	CreateUser(email, password string) (models.User, error)

	GetProfile(userID string) (models.Profile, error)
	SaveProfile(userID string, profile models.Profile) error

	GetGlobalRecipe(id string) (models.GlobalRecipe, error)
	SaveGlobalRecipe(recipe models.GlobalRecipe) error

	GetUserRecipe(userID string, recipeID string) (models.UserRecipe, error)
	GetAllUserRecipes(userID string) ([]models.UserRecipe, error)
	SaveUserRecipe(recipe models.UserRecipe) error
	UpdateUserRecipeVersion(userID string, recipeID string, versionID string) error

	GetRecipeVersion(recipeVersionID string) (models.RecipeVersion, error)
	AddRecipeVersion(recipeVersion models.RecipeVersion) error

	GetAllPlans(userID string) ([]models.MealPlan, error)
	GetMealPlan(userID string, mealPlanID string) (models.MealPlan, error)
	SaveMealPlan(userID string, mealPlan models.MealPlan) error
}
