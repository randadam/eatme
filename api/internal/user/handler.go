package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ajohnston1219/eatme/api/internal/api"
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
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
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	err := h.service.store.WithTx(func(tx db.Store) error {
		ctx := db.ContextWithTx(r.Context(), tx)
		user, err := h.service.CreateUser(ctx, input.Email, input.Password)
		if err != nil {
			switch {
			case errors.Is(err, ErrEmailExists):
				api.ErrorJSON(w, http.StatusConflict, models.ApiErrEmailExists)
			default:
				api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
			}
			return nil
		}
		api.WriteJSON(w, http.StatusCreated, models.SignupResponse{Token: user.ID})
		return nil
	})
	if err != nil {
		api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		return
	}
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
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	var profile models.ProfileUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		api.ErrorJSON(w, http.StatusBadRequest, models.ApiErrBadRequest)
		return
	}

	var result *models.Profile
	err := h.service.store.WithTx(func(tx db.Store) error {
		var err error
		ctx := db.ContextWithTx(r.Context(), tx)
		result, err = h.service.SaveProfile(ctx, userID, profile)
		if err != nil {
			api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
			return nil
		}
		api.WriteJSON(w, http.StatusOK, result)
		return nil
	})
	if err != nil {
		api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		return
	}

	api.WriteJSON(w, http.StatusOK, result)
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
	userID := api.GetUserID(r)
	if userID == "" {
		api.ErrorJSON(w, http.StatusUnauthorized, models.ApiErrUnauthorized)
		return
	}

	var profile *models.Profile
	err := h.service.store.WithTx(func(tx db.Store) error {
		var err error
		ctx := db.ContextWithTx(r.Context(), tx)
		profile, err = h.service.GetProfile(ctx, userID)
		if err != nil {
			switch {
			case errors.Is(err, ErrProfileNotFound):
				api.ErrorJSON(w, http.StatusNotFound, models.ApiErrProfileNotFound)
			default:
				api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
			}
			return nil
		}
		api.WriteJSON(w, http.StatusOK, profile)
		return nil
	})
	if err != nil {
		api.ErrorJSON(w, http.StatusInternalServerError, models.ApiErrInternal)
		return
	}

	api.WriteJSON(w, http.StatusOK, profile)
}
