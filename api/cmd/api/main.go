package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/ajohnston1219/eatme/api/docs"
	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/router"
)

// @title EatMe API
// @version 1.0
// @description API for the EatMe recipe generation service
// @host localhost:8080
// @BasePath /
func main() {
	dsn, ok := os.LookupEnv("DB_DSN")
	if !ok {
		dsn = "file:./.data/dev.db"
	}
	store, err := db.NewSQLiteStore(dsn)
	if err != nil {
		panic(err)
	}

	mlHost, ok := os.LookupEnv("ML_HOST")
	if !ok {
		mlHost = "http://ml-gateway:8000"
	}

	app := router.NewApp(store, clients.NewMLClient(mlHost))
	router := router.NewRouter(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)
	log.Printf("📚 API documentation available at http://localhost:%s/swagger/index.html", port)
	http.ListenAndServe(":"+port, router)
}
