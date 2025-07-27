package main

import (
	"context"
	"net/http"
	"os"

	_ "github.com/ajohnston1219/eatme/api/docs"
	"github.com/ajohnston1219/eatme/api/internal/clients"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/router"
	"github.com/ajohnston1219/eatme/api/internal/utils/logger"
	"github.com/ajohnston1219/eatme/api/internal/utils/telemetry"
	"go.uber.org/zap"
)

// @title EatMe API
// @version 1.0
// @description API for the EatMe recipe generation service
// @host localhost:8080
// @BasePath /
func main() {
	baseLogger, err := zap.NewDevelopment(zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(baseLogger)

	ctx := context.Background()
	shutdown := telemetry.InitTracer(ctx, "backend-api")
	defer shutdown(ctx)

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

	logger.Logger(ctx).Info("🚀 Server running on port %s", zap.String("port", port))
	logger.Logger(ctx).Info("📚 API documentation available at http://localhost:%s/swagger/index.html", zap.String("port", port))
	http.ListenAndServe(":"+port, router)
}
