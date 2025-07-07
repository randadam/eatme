package tests

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/router"
	"github.com/ajohnston1219/eatme/api/models"
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

func makeFakeRecipe(title string) models.RecipeBody {
	return models.RecipeBody{
		Title:            title,
		Description:      "Description",
		TotalTimeMinutes: 60,
		Servings:         4,
		Ingredients:      []models.Ingredient{{Name: "Ingredient 1"}, {Name: "Ingredient 2"}},
		Steps:            []string{"Step 1", "Step 2"},
	}
}

func createUser(store *db.SQLiteStore, userID string) {
	store.CreateUser(context.Background(), userID, "test@example.com")

	defaultProfile := models.Profile{}
	store.SaveProfile(context.Background(), userID, defaultProfile)
}
