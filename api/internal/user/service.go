package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/internal/models"
	"go.uber.org/zap"
)

type UserService struct {
	store db.Store
}

func NewUserService(store db.Store) *UserService {
	return &UserService{store: store}
}

func (s *UserService) getStore(ctx context.Context) db.Store {
	if tx, ok := db.GetTx(ctx); ok {
		return tx
	}
	return s.store
}

func (s *UserService) CreateUser(ctx context.Context, email string, password string) (*models.User, error) {
	var user models.User

	store := s.getStore(ctx)
	user, err := store.CreateUser(ctx, email, password)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrEmailExists):
			return nil, ErrEmailExists
		default:
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}
	zap.L().Debug("created user")
	defaultProfile := models.Profile{
		SetupStep: models.SetupStepProfile,
		Name:      "",
		Skill:     models.SkillBeginner,
		Cuisines:  []models.Cuisine{},
		Diet:      []models.Diet{},
		Equipment: []models.Equipment{},
		Allergies: []models.Allergy{},
	}

	err = store.SaveProfile(ctx, user.ID, defaultProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to save profile in create user: %w", err)
	}
	zap.L().Debug("saved default profile")
	return &user, nil
}

func (s *UserService) SaveProfile(ctx context.Context, userID string, profile models.ProfileUpdateRequest) (*models.Profile, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	store := s.getStore(ctx)
	currentProfile, err := store.GetProfile(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrNotFound):
			return nil, ErrProfileNotFound
		default:
			return nil, fmt.Errorf("failed to get profile: %w", err)
		}
	}
	zap.L().Debug("found profile")

	if profile.SetupStep != "" {
		currentProfile.SetupStep = profile.SetupStep
	}
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

	err = store.SaveProfile(ctx, userID, currentProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to save profile: %w", err)
	}
	zap.L().Debug("saved profile")
	return &currentProfile, nil
}

func (s *UserService) GetProfile(ctx context.Context, userID string) (*models.Profile, error) {
	store := s.getStore(ctx)
	profile, err := store.GetProfile(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrNotFound):
			return nil, ErrProfileNotFound
		default:
			return nil, fmt.Errorf("failed to get profile: %w", err)
		}
	}

	return &profile, nil
}
