package services

import (
	"errors"

	"github.com/ajohnston1219/eatme/api/db"
	"github.com/ajohnston1219/eatme/api/models"
)

type UserService struct {
	store db.Store
}

func NewUserService(store db.Store) *UserService {
	return &UserService{store: store}
}

func (s *UserService) CreateUser(firstName, lastName, email string) (models.User, error) {
	user, err := s.store.CreateUser(firstName, lastName, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) SavePreferences(prefs models.Preferences) error {
	if prefs.UserID == "" {
		return errors.New("user ID is required")
	}

	err := s.store.SavePreferences(prefs)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetPreferences(userID string) (models.Preferences, error) {
	prefs, err := s.store.GetPreferences(userID)
	if err != nil {
		return models.Preferences{}, err
	}

	return prefs, nil
}
