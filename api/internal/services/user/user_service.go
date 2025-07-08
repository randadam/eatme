package user

import (
	"context"
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

func (s *UserService) CreateUser(ctx context.Context, email string, password string) (models.User, error) {
	var user models.User
	err := s.store.WithTx(func(tx db.Store) error {
		var err error
		user, err = tx.CreateUser(ctx, email, password)
		if err != nil {
			return err
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

		return tx.SaveProfile(ctx, user.ID, defaultProfile)
	})
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) SaveProfile(ctx context.Context, userID string, profile models.ProfileUpdateRequest) (models.Profile, error) {
	if userID == "" {
		return models.Profile{}, errors.New("user ID is required")
	}

	var currentProfile models.Profile

	err := s.store.WithTx(func(tx db.Store) error {
		var err error
		currentProfile, err = tx.GetProfile(ctx, userID)
		if err != nil {
			return err
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

		return tx.SaveProfile(ctx, userID, currentProfile)
	})

	if err != nil {
		return models.Profile{}, err
	}
	return currentProfile, nil
}

func (s *UserService) GetProfile(ctx context.Context, userID string) (models.Profile, error) {
	profile, err := s.store.GetProfile(ctx, userID)
	if err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}
