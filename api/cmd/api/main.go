package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/ajohnston1219/eatme/api/docs"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/handlers"
	"github.com/ajohnston1219/eatme/api/internal/services/user"
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

	dsn, ok := os.LookupEnv("DB_DSN")
	if !ok {
		dsn = "file:./.data/dev.db"
	}
	store, err := db.NewSQLiteStore(dsn)
	if err != nil {
		panic(err)
	}

	corsHandler := cors.AllowAll().Handler(r)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
	))

	// User
	service := user.NewUserService(store)
	handler := handlers.NewUserHandler(service)
	r.Post("/signup", handler.Signup)
	r.Route("/profile", func(r chi.Router) {
		r.Use(handlers.AuthMiddleware(store))
		r.Put("/", handler.SaveProfile)
		r.Get("/", handler.GetProfile)
	})

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
