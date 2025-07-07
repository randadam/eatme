package db

import "github.com/ajohnston1219/eatme/api/models"

type Store interface {
	CreateUser(email, password string) (models.User, error)

	GetProfile(userID string) (models.Profile, error)
	SaveProfile(userID string, profile models.Profile) error

	GetAllPlans(userID string) ([]models.MealPlan, error)
	GetMealPlan(userID string, mealPlanID string) (models.MealPlan, error)
	SaveMealPlan(userID string, mealPlan models.MealPlan) error
}
