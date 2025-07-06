package db

import "github.com/ajohnston1219/eatme/api/models"

type Store interface {
	CreateUser(email, password string) (models.User, error)
	SaveProfile(userID string, profile models.Profile) error
	GetProfile(userID string) (models.Profile, error)
}
