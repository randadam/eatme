package router

import (
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/handlers"
	"github.com/ajohnston1219/eatme/api/internal/services/chat"
	"github.com/ajohnston1219/eatme/api/internal/services/recipe"
	"github.com/ajohnston1219/eatme/api/internal/services/user"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	store    db.Store
	mlClient clients.MLClient
}

func NewApp(store db.Store, mlClient clients.MLClient) *App {
	return &App{
		store:    store,
		mlClient: mlClient,
	}
}

func NewRouter(app *App) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
	))

	// User
	userService := user.NewUserService(app.store)
	userHandler := handlers.NewUserHandler(userService)
	r.Post("/signup", userHandler.Signup)
	r.Route("/profile", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware(app.store))
		r.Put("/", userHandler.SaveProfile)
		r.Get("/", userHandler.GetProfile)
	})

	// Recipe
	recipeService := recipe.NewRecipeService(app.store)
	recipeHandler := handlers.NewRecipeHandler(recipeService)
	r.Route("/recipes", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware(app.store))
		r.Get("/{recipe_id}", recipeHandler.GetRecipe)
		r.Get("/", recipeHandler.GetAllRecipes)
	})

	// Chat
	chatService := chat.NewChatService(app.mlClient, userService, recipeService)
	chatHandler := handlers.NewChatHandler(chatService)
	r.Route("/chat", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware(app.store))
		r.Post("/suggest", chatHandler.StartSuggestChat)
		r.Get("/suggest/{threadId}", chatHandler.GetSuggestionThread)
		r.Post("/suggest/{threadId}/next", chatHandler.NextRecipeSuggestion)
		r.Post("/suggest/{threadId}/accept/{suggestionId}", chatHandler.AcceptRecipeSuggestion)
		r.Put("/modify/recipes/{recipeId}", chatHandler.ModifyRecipeChat)
		r.Post("/question/recipes/{recipeId}", chatHandler.GeneralRecipeChat)
	})

	return r
}
