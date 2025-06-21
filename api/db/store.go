// db/store.go
package db

import "github.com/ajohnston1219/eatme/api/models"

type Store interface {
	CreateUser(firstName, lastName, email string) (models.User, error)
	SavePreferences(prefs models.Preferences) error
	GetPreferences(userID string) (models.Preferences, error)
}
