package user

import (
	"errors"

	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/models"
)

type UserService struct {
	store db.Store
}

func NewUserService(store db.Store) *UserService {
	return &UserService{store: store}
}

func (s *UserService) CreateUser(email string, password string) (models.User, error) {
	user, err := s.store.CreateUser(email, password)
	if err != nil {
		return models.User{}, err
	}
	err = s.store.SaveProfile(user.ID, models.Profile{SetupStep: models.SetupStepProfile})
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) SaveProfile(userID string, profile models.Profile) error {
	if userID == "" {
		return errors.New("user ID is required")
	}

	err := s.store.SaveProfile(userID, profile)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetProfile(userID string) (models.Profile, error) {
	prefs, err := s.store.GetProfile(userID)
	if err != nil {
		return models.Profile{}, err
	}

	return prefs, nil
}
