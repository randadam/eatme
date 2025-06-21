package db

import (
	"errors"
	"sync"

	"github.com/ajohnston1219/eatme/api/models"
	"github.com/google/uuid"
)

// MemoryStore implements UserStore interface with in-memory storage
type MemoryStore struct {
	mu          sync.RWMutex
	users       map[string]models.User
	preferences map[string]models.Preferences
}

// NewMemoryStore creates a new instance of MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:       make(map[string]models.User),
		preferences: make(map[string]models.Preferences),
	}
}

// CreateUser creates a new user in memory
func (s *MemoryStore) CreateUser(firstName, lastName, email string) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if email already exists
	for _, user := range s.users {
		if user.Email == email {
			return models.User{}, errors.New("email already exists")
		}
	}

	user := models.User{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	s.users[user.ID] = user
	return user, nil
}

// SavePreferences saves user preferences in memory
func (s *MemoryStore) SavePreferences(prefs models.Preferences) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Verify user exists
	if _, exists := s.users[prefs.UserID]; !exists {
		return errors.New("user not found")
	}

	s.preferences[prefs.UserID] = prefs
	return nil
}

// GetPreferences retrieves user preferences from memory
func (s *MemoryStore) GetPreferences(userID string) (models.Preferences, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Verify user exists
	if _, exists := s.users[userID]; !exists {
		return models.Preferences{}, errors.New("user not found")
	}

	prefs, exists := s.preferences[userID]
	if !exists {
		return models.Preferences{}, errors.New("preferences not found")
	}

	return prefs, nil
}
