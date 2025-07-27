package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/ajohnston1219/eatme/api/internal/api"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/ajohnston1219/eatme/api/internal/utils/logger"
	"go.uber.org/zap"
)

func AuthMiddleware(store db.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
				return
			}
			userID := strings.TrimPrefix(token, "Bearer ")
			_, err := store.GetUser(r.Context(), userID)
			if err != nil {
				zap.L().Debug("failed to get user", zap.Error(err))
				api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), api.UserIDKey{}, userID)
			ctx = logger.SetLogger(ctx, zap.L().With(zap.String("user_id", userID)))
			*r = *r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger := logger.Logger(r.Context())

		srw := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(srw, r)

		logger.Info("http_request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status", srw.status),
			zap.Duration("duration_ms", time.Since(start)),
			zap.String("remote_ip", r.RemoteAddr),
		)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}
