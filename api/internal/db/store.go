package db

import (
	"context"

	"github.com/ajohnston1219/eatme/api/internal/models"
)

type Store interface {
	CreateUser(ctx context.Context, email, password string) (models.User, error)
	GetUser(ctx context.Context, userID string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	CheckPassword(ctx context.Context, userID string, password string) error

	GetProfile(ctx context.Context, userID string) (models.Profile, error)
	SaveProfile(ctx context.Context, userID string, profile models.Profile) error

	GetGlobalRecipe(ctx context.Context, id string) (models.GlobalRecipe, error)
	SaveGlobalRecipe(ctx context.Context, recipe models.GlobalRecipe) error

	CreateThread(ctx context.Context, userID string, thread models.Thread) error
	GetThread(ctx context.Context, threadID string) (models.Thread, error)
	AppendToThread(ctx context.Context, threadID string, events []models.ThreadEvent) error
	AssociateThreadWithRecipe(ctx context.Context, threadID string, recipeID string) error

	GetUserRecipe(ctx context.Context, userID string, recipeID string) (models.UserRecipe, error)
	GetAllUserRecipes(ctx context.Context, userID string) ([]models.UserRecipe, error)
	SaveUserRecipe(ctx context.Context, recipe models.UserRecipe) error
	UpdateUserRecipeVersion(ctx context.Context, userID string, recipeID string, version models.RecipeVersion) error
	DeleteUserRecipe(ctx context.Context, userID string, recipeID string) error

	GetRecipeVersion(ctx context.Context, recipeVersionID string) (models.RecipeVersion, error)
	AddRecipeVersion(ctx context.Context, recipeVersion models.RecipeVersion) error

	GetAllPlans(ctx context.Context, userID string) ([]models.MealPlan, error)
	GetMealPlan(ctx context.Context, userID string, mealPlanID string) (models.MealPlan, error)
	SaveMealPlan(ctx context.Context, userID string, mealPlan models.MealPlan) error

	WithTx(fn func(tx Store) error) error
}
