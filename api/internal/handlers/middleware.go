package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ajohnston1219/eatme/api/internal/db"
)

func AuthMiddleware(store db.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				errorJSON(w, errors.New("missing auth"), http.StatusUnauthorized)
				return
			}
			userID := strings.TrimPrefix(token, "Bearer ")
			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
