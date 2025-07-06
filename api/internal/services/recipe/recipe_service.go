package recipe

import (
	"github.com/ajohnston1219/eatme/api/internal/db"
	"github.com/ajohnston1219/eatme/api/models"
)

type RecipeService struct {
	store db.Store
}

func NewRecipeService(store db.Store) *RecipeService {
	return &RecipeService{store: store}
}

func (s *RecipeService) GetMealPlan(userID string, mealPlanID string) (models.MealPlan, error) {
	recipes := []models.Recipe{
		{
			ID:               "pasta_tomato_basil",
			Title:            "Classic Tomato Basil Pasta",
			Description:      "Al dente spaghetti tossed in a quick garlic-tomato sauce with fresh basil.",
			TotalTimeMinutes: 25,
			Servings:         2,
			Ingredients: []models.Ingredient{
				{Name: "Spaghetti", Quantity: 200, Unit: "g"},
				{Name: "Garlic cloves", Quantity: 2, Unit: ""},
				{Name: "Crushed tomatoes", Quantity: 400, Unit: "g"},
				{Name: "Fresh basil", Quantity: 0.5, Unit: "cup"},
			},
			Steps: []string{
				"Cook spaghetti according to package.",
				"Sauté garlic, add tomatoes, simmer 10 min.",
				"Toss pasta with sauce and basil.",
			},
		},
		{
			ID:               "pasta_pesto",
			Title:            "Classic Tomato Basil Pasta",
			Description:      "Al dente spaghetti tossed in a quick garlic-tomato sauce with fresh basil.",
			TotalTimeMinutes: 25,
			Servings:         2,
			Ingredients: []models.Ingredient{
				{Name: "Spaghetti", Quantity: 200, Unit: "g"},
				{Name: "Pesto", Quantity: 1, Unit: "tbsp"},
				{Name: "Crushed tomatoes", Quantity: 400, Unit: "g"},
				{Name: "Fresh basil", Quantity: 0.5, Unit: "cup"},
			},
			Steps: []string{
				"Cook spaghetti according to package.",
				"Sauté garlic, add tomatoes, simmer 10 min.",
				"Toss pasta with sauce and basil.",
			},
		},
		{
			ID:               "pasta_alfredo",
			Title:            "Classic Alfredo Pasta",
			Description:      "Al dente spaghetti tossed in a quick garlic-tomato sauce with fresh basil.",
			TotalTimeMinutes: 25,
			Servings:         2,
			Ingredients: []models.Ingredient{
				{Name: "Spaghetti", Quantity: 200, Unit: "g"},
				{Name: "Butter", Quantity: 1, Unit: "tbsp"},
				{Name: "Heavy cream", Quantity: 200, Unit: "ml"},
				{Name: "Grated Parmesan", Quantity: 50, Unit: "g"},
			},
			Steps: []string{
				"Cook spaghetti according to package.",
				"Sauté garlic, add tomatoes, simmer 10 min.",
				"Toss pasta with sauce and basil.",
			},
		},
	}

	return models.MealPlan{Recipes: recipes}, nil
}
