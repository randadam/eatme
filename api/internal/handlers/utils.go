package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/ajohnston1219/eatme/api/internal/models"
)

func getUserID(r *http.Request) string {
	return r.Context().Value(userIDKey{}).(string)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func errorJSON(w http.ResponseWriter, status int, err models.APIError) {
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
