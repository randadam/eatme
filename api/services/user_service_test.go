// services/user_service_test.go
package services

import (
	"testing"

	"github.com/ajohnston1219/eatme/api/db"
	"github.com/ajohnston1219/eatme/api/models"
)

func setupTestService() *UserService {
	store := db.NewMemoryStore()
	return NewUserService(store)
}

func TestCreateUser(t *testing.T) {
	svc := setupTestService()

	user, err := svc.CreateUser("Test", "User", "test@example.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.ID == "" {
		t.Error("expected user ID to be set")
	}
	if user.FirstName != "Test" || user.LastName != "User" || user.Email != "test@example.com" {
		t.Errorf("user data mismatch: got %+v", user)
	}
}

func TestSaveAndGetPreferences(t *testing.T) {
	svc := setupTestService()

	user, _ := svc.CreateUser("Prefs", "User", "prefs@example.com")
	prefs := models.Preferences{
		UserID:           user.ID,
		DietRestrictions: []models.DietRestriction{models.DietKeto, models.DietVegan},
		Allergies: []models.Allergy{
			{Name: "peanuts", Severity: models.SeveritySevere},
		},
		Goals: &models.Goals{
			LimitCalories: ptrInt(2000),
			MacrosTarget: &models.Macros{
				ProteinGrams: 150,
				CarbsGrams:   100,
				FatGrams:     70,
			},
		},
	}

	err := svc.SavePreferences(prefs)
	if err != nil {
		t.Fatalf("expected no error saving preferences, got %v", err)
	}

	saved, err := svc.GetPreferences(user.ID)
	if err != nil {
		t.Fatalf("expected to retrieve preferences, got %v", err)
	}
	if saved.UserID != prefs.UserID {
		t.Error("user ID mismatch in saved preferences")
	}
	if len(saved.DietRestrictions) != 2 || saved.DietRestrictions[0] != models.DietKeto {
		t.Errorf("unexpected diet restrictions: %+v", saved.DietRestrictions)
	}
	if saved.Goals == nil || saved.Goals.MacrosTarget == nil {
		t.Error("expected macros goal to be set")
	}
}

func ptrInt(v int) *int {
	return &v
}
