package db

import (
	"reflect"
	"testing"

	"github.com/ajohnston1219/eatme/api/models"
)

func TestMemoryStore_CreateUser(t *testing.T) {
	store := NewMemoryStore()

	// Test successful user creation
	user, err := store.CreateUser("Test", "User", "test@example.com")
	if err != nil {
		t.Errorf("unexpected error creating user: %v", err)
	}
	if user.FirstName != "Test" || user.LastName != "User" || user.Email != "test@example.com" {
		t.Errorf("user data mismatch. got first name=%s last name=%s email=%s, want first name='Test' last name='User' email='test@example.com'",
			user.FirstName, user.LastName, user.Email)
	}

	// Test duplicate email
	_, err = store.CreateUser("Another", "User", "test@example.com")
	if err == nil {
		t.Error("expected error for duplicate email, got nil")
	}
}

func TestMemoryStore_Preferences(t *testing.T) {
	store := NewMemoryStore()

	// Create a test user first
	user, err := store.CreateUser("Test", "User", "test@example.com")
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	// Test saving preferences
	prefs := models.Preferences{
		UserID:           user.ID,
		DietRestrictions: []models.DietRestriction{models.DietVegetarian},
		Allergies:        []models.Allergy{{Name: "peanuts", Severity: models.SeveritySevere}},
	}

	err = store.SavePreferences(prefs)
	if err != nil {
		t.Errorf("unexpected error saving preferences: %v", err)
	}

	// Test getting preferences
	savedPrefs, err := store.GetPreferences(user.ID)
	if err != nil {
		t.Errorf("unexpected error getting preferences: %v", err)
	}
	if !reflect.DeepEqual(savedPrefs.DietRestrictions, prefs.DietRestrictions) {
		t.Errorf("preferences diet mismatch. got=%s, want=%s", savedPrefs.DietRestrictions, prefs.DietRestrictions)
	}
	if len(savedPrefs.Allergies) != len(prefs.Allergies) {
		t.Errorf("preferences allergies length mismatch. got=%d, want=%d",
			len(savedPrefs.Allergies), len(prefs.Allergies))
	}
	if !reflect.DeepEqual(savedPrefs.Allergies, prefs.Allergies) {
		t.Errorf("preferences allergies mismatch. got=%s, want=%s", savedPrefs.Allergies, prefs.Allergies)
	}

	// Test getting preferences for non-existent user
	_, err = store.GetPreferences("non-existent-id")
	if err == nil {
		t.Error("expected error getting preferences for non-existent user, got nil")
	}

	// Test saving preferences for non-existent user
	badPrefs := models.Preferences{
		UserID:           "non-existent-id",
		DietRestrictions: []models.DietRestriction{models.DietVegetarian},
	}
	err = store.SavePreferences(badPrefs)
	if err == nil {
		t.Error("expected error saving preferences for non-existent user, got nil")
	}
}
