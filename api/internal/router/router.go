package router

import (
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/chat"
	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/middleware"
	"github.com/ajohnston1219/eatme/api/internal/recipe"
	"github.com/ajohnston1219/eatme/api/internal/thread"
	"github.com/ajohnston1219/eatme/api/internal/user"
	"github.com/ajohnston1219/eatme/api/internal/utils"
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
	utils.InitLogger()

	r := chi.NewRouter()

	r.Use(middleware.RequestLogger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
	))

	// User
	userService := user.NewUserService(app.store)
	userHandler := user.NewUserHandler(userService)
	r.Post("/signup", userHandler.Signup)
	r.Route("/profile", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(app.store))
		r.Put("/", userHandler.SaveProfile)
		r.Get("/", userHandler.GetProfile)
	})

	// Recipe
	recipeService := recipe.NewRecipeService(app.store)
	recipeHandler := recipe.NewRecipeHandler(recipeService)
	r.Route("/recipes", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(app.store))
		r.Get("/{recipe_id}", recipeHandler.GetRecipe)
		r.Get("/", recipeHandler.GetAllRecipes)
		r.Delete("/{recipe_id}", recipeHandler.DeleteRecipe)
	})

	// Thread
	chatService := chat.NewChatService(app.mlClient)
	threadService := thread.NewThreadService(app.store, recipeService, chatService)
	threadHandler := thread.NewThreadHandler(threadService)
	r.Route("/thread", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(app.store))
		r.Post("/suggest", threadHandler.StartSuggestionThread)
		r.Post("/{threadId}/suggest", threadHandler.GetNewSuggestions)
		r.Post("/{threadId}/accept/{suggestionId}", threadHandler.AcceptSuggestion)
		r.Post("/{threadId}/modify/chat", threadHandler.ModifyRecipeViaChat)
		r.Post("/{threadId}/question", threadHandler.AnswerCookingQuestion)
		r.Get("/{threadId}", threadHandler.GetThread)
	})

	return r
}
