package meal

import (
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/models"
	"github.com/google/uuid"
)

type MealService struct {
	store db.Store
}

func NewMealService(store db.Store) *MealService {
	return &MealService{store: store}
}

func (s *MealService) NewMealPlan(userID string) (models.MealPlan, error) {
	id := uuid.New().String()
	plan := models.MealPlan{
		ID:      id,
		Recipes: []*models.Recipe{},
	}
	err := s.store.SaveMealPlan(userID, plan)
	if err != nil {
		return models.MealPlan{}, err
	}
	return plan, nil
}

func (s *MealService) SaveMealPlan(userID string, mealPlan models.MealPlan) error {
	return s.store.SaveMealPlan(userID, mealPlan)
}

func (s *MealService) GetMealPlan(userID string, mealPlanID string) (models.MealPlan, error) {
	return s.store.GetMealPlan(userID, mealPlanID)
}

func (s *MealService) GetAllPlans(userID string) ([]models.MealPlan, error) {
	return s.store.GetAllPlans(userID)
}
