package api

import (
	"encoding/json"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/models"
	"go.uber.org/zap"
)

type UserIDKey struct{}

func GetUserID(r *http.Request) string {
	return r.Context().Value(UserIDKey{}).(string)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func ErrorJSON(w http.ResponseWriter, status int, err models.APIError) {
	if status >= 500 {
		zap.L().Error("server error", zap.Error(err))
	} else {
		zap.L().Warn("client error", zap.Error(err))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"error": err,
	})
}
