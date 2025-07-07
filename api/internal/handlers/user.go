package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/services/user"
	"github.com/ajohnston1219/eatme/api/models"
)

type UserHandler struct {
	service *user.UserService
}

func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Create a new user account
// @Description Register a new user account
// @ID signup
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
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(input.Email, input.Password)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, models.SignupResponse{Token: user.ID})
}

// @Summary Save user profile
// @Description Save a user's profile
// @ID saveProfile
// @Tags users
// @Accept json
// @Produce json
// @Param profile body models.ProfileUpdateRequest true "User profile"
// @Success 200 {object} models.Profile
// @Failure 400 {object} models.BadRequestResponse "Invalid input"
// @Failure 401 {object} models.UnauthorizedResponse "Unauthorized"
// @Failure 500 {object} models.InternalServerErrorResponse "Internal server error"
// @Router /profile [put]
func (h *UserHandler) SaveProfile(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	var profile models.ProfileUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	result, err := h.service.SaveProfile(userID, profile)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, result)
}

// @Summary Get user profile
// @Description Gets the profile for a user
// @ID getProfile
// @Tags users
// @Produce json
// @Success 200 {object} models.Profile
// @Failure 401 {object} models.UnauthorizedResponse "Unauthorized"
// @Failure 500 {object} models.InternalServerErrorResponse "Internal server error"
// @Router /profile [get]
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, errors.New("missing user ID"), http.StatusUnauthorized)
		return
	}

	profile, err := h.service.GetProfile(userID)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, profile)
}
