package recipe

import (
	"slices"

	"github.com/ajohnston1219/eatme/api/internal/models"
)

func GetRecipeDiff(currentRecipe *models.RecipeBody, proposedRecipe *models.RecipeBody) *models.RecipeDiff {
	diff := &models.RecipeDiff{
		NewTitle:            nil,
		NewDescription:      nil,
		NewServings:         nil,
		NewTotalTimeMinutes: nil,
		AddedIngredients:    []models.Ingredient{},
		RemovedIngredients:  []models.RemovedIngredient{},
		ModifiedIngredients: []models.ModifiedIngredient{},
		NewSteps:            []models.DiffStep{},
	}

	if currentRecipe.Title != proposedRecipe.Title {
		diff.NewTitle = &proposedRecipe.Title
	}

	if currentRecipe.Description != proposedRecipe.Description {
		diff.NewDescription = &proposedRecipe.Description
	}

	if currentRecipe.Servings != proposedRecipe.Servings {
		diff.NewServings = &proposedRecipe.Servings
	}

	if currentRecipe.TotalTimeMinutes != proposedRecipe.TotalTimeMinutes {
		diff.NewTotalTimeMinutes = &proposedRecipe.TotalTimeMinutes
	}

	currIdx := 0
	foundProposed := make([]bool, len(proposedRecipe.Ingredients))
	for currIdx < len(currentRecipe.Ingredients) {
		currentIngredient := currentRecipe.Ingredients[currIdx]
		found := false
		for j, proposedIngredient := range proposedRecipe.Ingredients {
			if currentIngredient.Name == proposedIngredient.Name {
				if currentIngredient.Quantity != proposedIngredient.Quantity || currentIngredient.Unit != proposedIngredient.Unit {
					diff.ModifiedIngredients = append(diff.ModifiedIngredients, models.ModifiedIngredient{
						Index:    currIdx,
						Name:     proposedIngredient.Name,
						Quantity: proposedIngredient.Quantity,
						Unit:     proposedIngredient.Unit,
					})
				}
				found = true
				foundProposed[j] = true
				break
			}
		}
		if !found {
			diff.RemovedIngredients = append(diff.RemovedIngredients, models.RemovedIngredient{
				Index: currIdx,
			})
		}
		currIdx++
	}

	for j, proposedIngredient := range proposedRecipe.Ingredients {
		if !foundProposed[j] {
			diff.AddedIngredients = append(diff.AddedIngredients, proposedIngredient)
		}
	}

	for _, proposedStep := range proposedRecipe.Steps {
		stepDiff := models.DiffStep{
			Step:  proposedStep,
			IsNew: false,
		}
		found := slices.Contains(currentRecipe.Steps, proposedStep)
		if !found {
			stepDiff.IsNew = true
		}

		diff.NewSteps = append(diff.NewSteps, stepDiff)
	}

	return diff
}
