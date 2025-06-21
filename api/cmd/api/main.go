package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ajohnston1219/eatme/api/db"
	_ "github.com/ajohnston1219/eatme/api/docs"
	"github.com/ajohnston1219/eatme/api/handlers"
	"github.com/ajohnston1219/eatme/api/services"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title EatMe API
// @version 1.0
// @description API for the EatMe recipe generation service
// @host localhost:8080
// @BasePath /
func main() {
	r := chi.NewRouter()

	corsHandler := cors.AllowAll().Handler(r)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
	))

	// User
	store := db.NewMemoryStore()
	service := services.NewUserService(store)
	handler := handlers.NewUserHandler(service)
	r.Post("/signup", handler.Signup)
	r.Post("/preferences", handler.SetPreferences)
	r.Get("/preferences", handler.GetPreferences)

	// LLM
	r.Post("/generate", handlers.GenerateRecipeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)
	log.Printf("📚 API documentation available at http://localhost:%s/swagger/index.html", port)
	http.ListenAndServe(":"+port, corsHandler)
}
