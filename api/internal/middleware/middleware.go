package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/ajohnston1219/eatme/api/internal/api"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
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
			ctx := context.WithValue(r.Context(), api.UserIDKey{}, userID)
			*r = *r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		srw := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(srw, r)

		userID, ok := r.Context().Value(api.UserIDKey{}).(string)
		if !ok {
			userID = ""
		}

		zap.L().Info("http_request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status", srw.status),
			zap.Duration("duration_ms", time.Since(start)),
			zap.String("remote_ip", r.RemoteAddr),
			zap.String("user_id", userID),
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
