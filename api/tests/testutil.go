package tests

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/ajohnston1219/eatme/api/internal/recipe"
	"github.com/ajohnston1219/eatme/api/internal/router"
	"github.com/ajohnston1219/eatme/api/internal/user"
)

func NewTestServer(t *testing.T, mlStub clients.MLClient) (*httptest.Server, *db.SQLiteStore) {
	t.Helper()

	// ➊ in-memory SQLite
	sqlDB, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	store, err := db.NewSQLiteStoreWithDB(sqlDB)
	if err != nil {
		t.Fatal(err)
	}

	app := router.NewApp(store, mlStub)
	r := router.NewRouter(app)

	ts := httptest.NewServer(r)
	return ts, store
}

type FakeRecipeOpt func(*models.RecipeBody)

func WithDescription(desc string) FakeRecipeOpt {
	return func(r *models.RecipeBody) {
		r.Description = desc
	}
}

func WithTotalTimeMinutes(minutes int) FakeRecipeOpt {
	return func(r *models.RecipeBody) {
		r.TotalTimeMinutes = minutes
	}
}

func WithServings(servings int) FakeRecipeOpt {
	return func(r *models.RecipeBody) {
		r.Servings = servings
	}
}

func WithIngredients(ingredients []models.Ingredient) FakeRecipeOpt {
	return func(r *models.RecipeBody) {
		r.Ingredients = ingredients
	}
}

func WithSteps(steps []string) FakeRecipeOpt {
	return func(r *models.RecipeBody) {
		r.Steps = steps
	}
}

func makeFakeRecipe(title string, opts ...FakeRecipeOpt) models.RecipeBody {
	recipe := models.RecipeBody{
		Title:            title,
		Description:      "Description",
		TotalTimeMinutes: 60,
		Servings:         4,
		Ingredients:      []models.Ingredient{{Name: "Ingredient 1"}, {Name: "Ingredient 2"}},
		Steps:            []string{"Step 1", "Step 2"},
	}
	for _, opt := range opts {
		opt(&recipe)
	}
	return recipe
}

func createUser(store *db.SQLiteStore, email string) (*models.User, error) {
	svc := user.NewUserService(store)
	user, err := svc.CreateUser(context.Background(), email, "password")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func createRecipe(store *db.SQLiteStore, userID string, recipeBody models.RecipeBody) (*models.UserRecipe, error) {
	svc := recipe.NewRecipeService(store)
	newRecipe, err := svc.NewRecipe(context.Background(), userID, recipeBody)
	return newRecipe, err
}
