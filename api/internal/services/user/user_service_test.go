package user_test

import (
	"errors"
	"testing"

	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/services/user"
	"github.com/ajohnston1219/eatme/api/internal/testutil"
	"github.com/ajohnston1219/eatme/api/models"
)

// ---------- happy path ------------------------------------------------------

func TestUserLifecycle(t *testing.T) {
	store := testutil.NewTestSQLiteStore(t)
	svc := user.NewUserService(store)

	// 1) Create user
	u, err := svc.CreateUser("alice@example.com", "SuperSecret123")
	if err != nil {
		t.Fatalf("CreateUser: %v", err)
	}
	if u.Email != "alice@example.com" {
		t.Fatalf("CreateUser: wrong email got %s", u.Email)
	}
	savedProfile, err := store.GetProfile(u.ID)
	if err != nil {
		t.Fatalf("GetProfile: %v", err)
	}
	if savedProfile.SetupStep != models.SetupStepProfile {
		t.Fatalf("GetProfile: wrong setup step got %s", savedProfile.SetupStep)
	}

	// 2) Save profile
	in := models.Profile{
		SetupStep: models.SetupStepDiet,
		Name:      "Alice",
		Skill:     models.SkillIntermediate,
		Cuisines:  []models.Cuisine{models.CuisineItalian, models.CuisineMexican},
		Diet:      []models.Diet{models.DietVegetarian},
		Equipment: []models.Equipment{models.EquipmentOven, models.EquipmentGrill},
		Allergies: []models.Allergy{models.AllergyPeanuts},
	}
	if err := svc.SaveProfile(u.ID, in); err != nil {
		t.Fatalf("SaveProfile: %v", err)
	}

	// 3) Get profile back
	got, err := svc.GetProfile(u.ID)
	if err != nil {
		t.Fatalf("GetProfile: %v", err)
	}
	if got.Name != in.Name || got.Skill != in.Skill {
		t.Errorf("GetProfile mismatch: got %+v want %+v", got, in)
	}
	if len(got.Cuisines) != 2 || got.Cuisines[0] != models.CuisineItalian {
		t.Errorf("Cuisines not stored correctly: %+v", got.Cuisines)
	}
}

// ---------- duplicate email -------------------------------------------------

func TestCreateUserDuplicateEmail(t *testing.T) {
	store := testutil.NewTestSQLiteStore(t)
	svc := user.NewUserService(store)

	if _, err := svc.CreateUser("dup@example.com", "pw1"); err != nil {
		t.Fatalf("first insert: %v", err)
	}
	_, err := svc.CreateUser("dup@example.com", "pw2")
	if err == nil {
		t.Fatalf("expected duplicate email error, got nil")
	}
	if !errors.Is(err, db.ErrEmailExists) {
		t.Fatalf("expected ErrEmailExists, got %v", err)
	}
}

// ---------- missing userID on SaveProfile -----------------------------------

func TestSaveProfileMissingUserID(t *testing.T) {
	store := testutil.NewTestSQLiteStore(t)
	svc := user.NewUserService(store)

	err := svc.SaveProfile("", models.Profile{})
	if err == nil {
		t.Fatalf("expected error for empty userID, got nil")
	}
}
