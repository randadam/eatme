package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ajohnston1219/eatme/api/models"
	"github.com/ajohnston1219/eatme/api/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Create a new user account
// @Description Register a new user with name and email
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.SignupRequest true "User signup information"
// @Success 200 {object} models.SignupResponse
// @Failure 400 {object} models.BadRequestResponse "Invalid input"
// @Failure 500 {object} models.InternalServerErrorResponse "Internal server error"
// @Router /signup [post]
func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var input models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(input.FirstName, input.LastName, input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.SignupResponse{UserID: user.ID})
}

// @Summary Set user preferences
// @Description Update or create preferences for a user
// @Tags users
// @Accept json
// @Produce json
// @Param X-User-ID header string true "User ID" example(usr_123456789)
// @Param preferences body models.Preferences true "User preferences"
// @Success 200 {object} models.Preferences
// @Failure 400 {object} models.BadRequestResponse "Invalid input"
// @Failure 401 {object} models.UnauthorizedResponse "Missing user ID"
// @Failure 500 {object} models.InternalServerErrorResponse "Internal server error"
// @Router /preferences [put]
func (h *UserHandler) SetPreferences(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "missing user ID", http.StatusUnauthorized)
		return
	}

	var prefs models.Preferences
	if err := json.NewDecoder(r.Body).Decode(&prefs); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	prefs.UserID = userID
	if err := h.service.SavePreferences(prefs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prefs)
}

// @Summary Get user preferences
// @Description Retrieve preferences for a user
// @Tags users
// @Produce json
// @Param X-User-ID header string true "User ID" example(usr_123456789)
// @Success 200 {object} models.Preferences
// @Failure 401 {object} models.UnauthorizedResponse "Missing user ID"
// @Failure 500 {object} models.InternalServerErrorResponse "Internal server error"
// @Router /preferences [get]
func (h *UserHandler) GetPreferences(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		http.Error(w, "missing user ID", http.StatusUnauthorized)
		return
	}

	prefs, err := h.service.GetPreferences(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prefs)
}
