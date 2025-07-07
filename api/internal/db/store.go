package db

import (
	"context"

	"github.com/ajohnston1219/eatme/api/models"
)

type Store interface {
	CreateUser(ctx context.Context, email, password string) (models.User, error)

	GetProfile(ctx context.Context, userID string) (models.Profile, error)
	SaveProfile(ctx context.Context, userID string, profile models.Profile) error

	GetGlobalRecipe(ctx context.Context, id string) (models.GlobalRecipe, error)
	SaveGlobalRecipe(ctx context.Context, recipe models.GlobalRecipe) error

	GetUserRecipe(ctx context.Context, userID string, recipeID string) (models.UserRecipe, error)
	GetAllUserRecipes(ctx context.Context, userID string) ([]models.UserRecipe, error)
	SaveUserRecipe(ctx context.Context, recipe models.UserRecipe) error
	UpdateUserRecipeVersion(ctx context.Context, userID string, recipeID string, versionID string) error

	GetRecipeVersion(ctx context.Context, recipeVersionID string) (models.RecipeVersion, error)
	AddRecipeVersion(ctx context.Context, recipeVersion models.RecipeVersion) error

	GetAllPlans(ctx context.Context, userID string) ([]models.MealPlan, error)
	GetMealPlan(ctx context.Context, userID string, mealPlanID string) (models.MealPlan, error)
	SaveMealPlan(ctx context.Context, userID string, mealPlan models.MealPlan) error

	WithTx(fn func(tx Store) error) error
}
