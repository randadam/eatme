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
	defaultProfile := models.Profile{
		SetupStep: models.SetupStepProfile,
		Name:      "",
		Skill:     models.SkillBeginner,
		Cuisines:  []models.Cuisine{},
		Diet:      []models.Diet{},
		Equipment: []models.Equipment{},
		Allergies: []models.Allergy{},
	}
	err = s.store.SaveProfile(user.ID, defaultProfile)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) SaveProfile(userID string, profile models.ProfileUpdateRequest) (models.Profile, error) {
	if userID == "" {
		return models.Profile{}, errors.New("user ID is required")
	}

	currentProfile, err := s.store.GetProfile(userID)
	if err != nil {
		return models.Profile{}, err
	}

	currentProfile.SetupStep = profile.SetupStep
	if profile.Name != "" {
		currentProfile.Name = profile.Name
	}
	if profile.Skill != "" {
		currentProfile.Skill = profile.Skill
	}
	if profile.Cuisines != nil {
		currentProfile.Cuisines = profile.Cuisines
	}
	if profile.Diet != nil {
		currentProfile.Diet = profile.Diet
	}
	if profile.Equipment != nil {
		currentProfile.Equipment = profile.Equipment
	}
	if profile.Allergies != nil {
		currentProfile.Allergies = profile.Allergies
	}

	err = s.store.SaveProfile(userID, currentProfile)
	if err != nil {
		return models.Profile{}, err
	}
	return currentProfile, nil
}

func (s *UserService) GetProfile(userID string) (models.Profile, error) {
	prefs, err := s.store.GetProfile(userID)
	if err != nil {
		return models.Profile{}, err
	}

	return prefs, nil
}
