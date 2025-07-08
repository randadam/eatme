package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/models"
	userService "github.com/ajohnston1219/eatme/api/internal/services/user"
)

type UserHandler struct {
	service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 409 {object} models.APIError "Email already exists"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /signup [post]
func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var input models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	user, err := h.service.CreateUser(r.Context(), input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, userService.ErrEmailExists):
			errorJSON(w, http.StatusConflict, models.ErrEmailExists)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
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
// @Failure 400 {object} models.APIError "Invalid input"
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /profile [put]
func (h *UserHandler) SaveProfile(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	var profile models.ProfileUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		errorJSON(w, http.StatusBadRequest, models.ErrBadRequest)
		return
	}

	result, err := h.service.SaveProfile(r.Context(), userID, profile)
	if err != nil {
		errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
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
// @Failure 401 {object} models.APIError "Unauthorized"
// @Failure 404 {object} models.APIError "Not found"
// @Failure 500 {object} models.APIError "Internal server error"
// @Router /profile [get]
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)
	if userID == "" {
		errorJSON(w, http.StatusUnauthorized, models.ErrUnauthorized)
		return
	}

	profile, err := h.service.GetProfile(r.Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, userService.ErrProfileNotFound):
			errorJSON(w, http.StatusNotFound, models.ErrProfileNotFound)
		default:
			errorJSON(w, http.StatusInternalServerError, models.ErrInternal)
		}
		return
	}

	writeJSON(w, http.StatusOK, profile)
}
