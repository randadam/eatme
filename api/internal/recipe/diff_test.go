package recipe

import (
	"testing"

	"github.com/ajohnston1219/eatme/api/internal/models"
)

var testCases = []struct {
	currentRecipe  *models.RecipeBody
	proposedRecipe *models.RecipeBody
	expectedDiff   *models.RecipeDiff
}{
	{
		currentRecipe: &models.RecipeBody{
			Title:       "Veal Bolognese",
			Description: "A classic Italian dish",
			Ingredients: []models.Ingredient{
				{Name: "Veal", Quantity: 1, Unit: models.MeasurementUnitGram},
				{Name: "Crushed Tomato", Quantity: 3, Unit: models.MeasurementUnitCup},
				{Name: "Flour", Quantity: 1, Unit: models.MeasurementUnitGram},
				{Name: "Sugar", Quantity: 1, Unit: models.MeasurementUnitGram},
			},
			Steps: []models.Step{
				"Step 1",
				"Step 2",
				"Step 3",
			},
			Servings:         4,
			TotalTimeMinutes: 120,
		},
		proposedRecipe: &models.RecipeBody{
			Title:       "Beef Bolognese",
			Description: "A classic Italian dish with beef",
			Ingredients: []models.Ingredient{
				{Name: "Beef", Quantity: 1, Unit: models.MeasurementUnitGram},
				{Name: "Flour", Quantity: 1, Unit: models.MeasurementUnitGram},
				{Name: "Sugar", Quantity: 2, Unit: models.MeasurementUnitGram},
				{Name: "Crushed Tomato", Quantity: 3, Unit: models.MeasurementUnitCup},
				{Name: "Onion", Quantity: 1, Unit: models.MeasurementUnitCount},
			},
			Steps: []models.Step{
				"Step 1",
				"Step 2 Changed",
				"Step 3",
				"Step 4",
			},
			Servings:         6,
			TotalTimeMinutes: 150,
		},
		expectedDiff: &models.RecipeDiff{
			NewTitle:            stringPointer("Beef Bolognese"),
			NewDescription:      stringPointer("A classic Italian dish with beef"),
			NewServings:         intPointer(6),
			NewTotalTimeMinutes: intPointer(150),
			AddedIngredients: []models.Ingredient{
				{Name: "Beef", Quantity: 1, Unit: models.MeasurementUnitGram},
				{Name: "Onion", Quantity: 1, Unit: models.MeasurementUnitCount},
			},
			RemovedIngredients: []models.RemovedIngredient{
				{Index: 0},
			},
			ModifiedIngredients: []models.ModifiedIngredient{
				{
					Index:    3,
					Name:     "Sugar",
					Quantity: 2,
					Unit:     models.MeasurementUnitGram,
				},
			},
			NewSteps: []models.DiffStep{
				{
					Step:  "Step 1",
					IsNew: false,
				},
				{
					Step:  "Step 2 Changed",
					IsNew: true,
				},
				{
					Step:  "Step 3",
					IsNew: false,
				},
				{
					Step:  "Step 4",
					IsNew: true,
				},
			},
		},
	},
}

func TestGetRecipeDiff(t *testing.T) {
	for _, tc := range testCases {
		diff := GetRecipeDiff(tc.currentRecipe, tc.proposedRecipe)

		expectEqualButCheckNilString("NewTitle", t, tc.expectedDiff.NewTitle, diff.NewTitle)
		expectEqualButCheckNilString("NewDescription", t, tc.expectedDiff.NewDescription, diff.NewDescription)
		expectEqualButCheckNilInt("NewServings", t, tc.expectedDiff.NewServings, diff.NewServings)
		expectEqualButCheckNilInt("NewTotalTimeMinutes", t, tc.expectedDiff.NewTotalTimeMinutes, diff.NewTotalTimeMinutes)

		for i, ingredient := range diff.AddedIngredients {
			if ingredient != tc.expectedDiff.AddedIngredients[i] {
				t.Errorf("expected added ingredient diff %v, got %v", tc.expectedDiff.AddedIngredients[i], ingredient)
			}
		}
		for i, ingredient := range diff.RemovedIngredients {
			if ingredient != tc.expectedDiff.RemovedIngredients[i] {
				t.Errorf("expected removed ingredient diff %v, got %v", tc.expectedDiff.RemovedIngredients[i], ingredient)
			}
		}
		for i, ingredient := range diff.ModifiedIngredients {
			if ingredient != tc.expectedDiff.ModifiedIngredients[i] {
				t.Errorf("expected modified ingredient diff %v, got %v", tc.expectedDiff.ModifiedIngredients[i], ingredient)
			}
		}
		for i, step := range diff.NewSteps {
			if step != tc.expectedDiff.NewSteps[i] {
				t.Errorf("expected step diff %v, got %v", tc.expectedDiff.NewSteps[i], step)
			}
		}
	}
}

func stringPointer(s string) *string {
	return &s
}

func intPointer(i int) *int {
	return &i
}

func expectEqualButCheckNilString(fieldName string, t *testing.T, expected *string, actual *string) {
	if expected == nil && actual != nil {
		t.Errorf("for %s expected nil, got %s", fieldName, *actual)
	}
	if expected != nil && actual == nil {
		t.Errorf("for %s expected %s, got nil", fieldName, *expected)
	}
	if expected != nil && actual != nil {
		if *expected != *actual {
			t.Errorf("expected %s, got %s", *expected, *actual)
		}
	}
}

func expectEqualButCheckNilInt(fieldName string, t *testing.T, expected *int, actual *int) {
	if expected == nil && actual != nil {
		t.Errorf("for %s expected nil, got %d", fieldName, *actual)
	}
	if expected != nil && actual == nil {
		t.Errorf("for %s expected %d, got nil", fieldName, *expected)
	}
	if expected != nil && actual != nil {
		if *expected != *actual {
			t.Errorf("for %s expected %d, got %d", fieldName, *expected, *actual)
		}
	}
}
