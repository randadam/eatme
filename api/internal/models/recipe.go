package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type MeasurementUnit string

const (
	MeasurementUnitGram       MeasurementUnit = "g"
	MeasurementUnitMilliliter MeasurementUnit = "ml"
	MeasurementUnitTeaspoon   MeasurementUnit = "tsp"
	MeasurementUnitTablespoon MeasurementUnit = "tbsp"
	MeasurementUnitCup        MeasurementUnit = "cup"
	MeasurementUnitOunce      MeasurementUnit = "oz"
	MeasurementUnitPound      MeasurementUnit = "lb"
	MeasurementUnitCount      MeasurementUnit = "count"
)

// @Description Ingredient represents an ingredient in a recipe
type Ingredient struct {
	Name     string          `json:"name" example:"Flour" binding:"required"`
	Quantity float64         `json:"quantity" example:"1" binding:"required"`
	Unit     MeasurementUnit `json:"unit" example:"cup" binding:"required"`
}

type Ingredients []Ingredient

func (i *Ingredients) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan ingredients: expected []byte, got %T", value)
	}
	return json.Unmarshal(b, i)
}

type Step string

type Steps []Step

func (s *Steps) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan steps: expected []byte, got %T", value)
	}
	return json.Unmarshal(b, s)
}

// @Description RecipeBody represents the contents of a recipe
type RecipeBody struct {
	Title            string      `json:"title" example:"Veal Bolognese" binding:"required"`
	Description      string      `json:"description" example:"A classic Italian dish" binding:"required"`
	Ingredients      Ingredients `json:"ingredients" binding:"required"`
	Steps            Steps       `json:"steps" binding:"required"`
	Servings         int         `json:"servings" example:"4" binding:"required"`
	TotalTimeMinutes int         `json:"total_time_minutes" example:"120" binding:"required"`
	ImageURL         string      `json:"image_url,omitempty"`
}

type RecipeSource string

const (
	RecipeSourceScraped   RecipeSource = "scraped"
	RecipeSourceGenerated RecipeSource = "generated"
)

// @Description GlobalRecipe represents a global recipe that can be "forked" into a user's recipe book
type GlobalRecipe struct {
	ID         string       `json:"id" binding:"required"`
	SourceType RecipeSource `json:"source_type" binding:"required"`
	CreatedAt  time.Time    `json:"created_at" binding:"required"`
	UpdatedAt  time.Time    `json:"updated_at" binding:"required"`
	RecipeBody
}

// @Description UserRecipe is the user's personal copy (favorites, edits).
type UserRecipe struct {
	ID              string    `json:"id" binding:"required"`
	UserID          string    `json:"user_id" binding:"required"`
	GlobalRecipeID  *string   `json:"global_recipe_id,omitempty"`
	ThreadID        string    `json:"thread_id" binding:"required"`
	IsFavorite      bool      `json:"is_favorite" binding:"required"`
	LatestVersionID string    `json:"latest_version_id" binding:"required"`
	CreatedAt       time.Time `json:"created_at" binding:"required"`
	UpdatedAt       time.Time `json:"updated_at" binding:"required"`
	RecipeBody
}

// @Description RecipeVersion is an immutable snapshot used inside meal plans.
type RecipeVersion struct {
	ID           string    `json:"id" binding:"required"`
	UserRecipeID string    `json:"user_recipe_id" binding:"required"`
	ParentID     *string   `json:"parent_id,omitempty"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
	Notes        *string   `json:"notes,omitempty"`
	RecipeBody
}

// @Description ModifiedIngredient represents a modification to an ingredient
type ModifiedIngredient struct {
	Index    int             `json:"index" binding:"required"`
	Name     string          `json:"name" example:"Flour" binding:"required"`
	Quantity float64         `json:"quantity" example:"1" binding:"required"`
	Unit     MeasurementUnit `json:"unit" example:"cup" binding:"required"`
}

// @Description RemovedIngredient represents an ingredient that was removed from a recipe
type RemovedIngredient struct {
	Index int `json:"index" binding:"required"`
}

// @Description DiffStep represents a step in a recipe with information about the type of change
type DiffStep struct {
	Step  Step `json:"step" binding:"required"`
	IsNew bool `json:"is_new" binding:"required"`
}

// @Description RecipeDiff represents the difference between two recipe versions
type RecipeDiff struct {
	NewTitle            *string              `json:"new_title,omitempty"`
	NewDescription      *string              `json:"new_description,omitempty"`
	NewServings         *int                 `json:"new_servings,omitempty"`
	NewTotalTimeMinutes *int                 `json:"new_total_time_minutes,omitempty"`
	AddedIngredients    []Ingredient         `json:"added_ingredients" binding:"required"`
	ModifiedIngredients []ModifiedIngredient `json:"modified_ingredients" binding:"required"`
	RemovedIngredients  []RemovedIngredient  `json:"removed_ingredients" binding:"required"`
	NewSteps            []DiffStep           `json:"new_steps" binding:"required"`
}

// @Description ModifyRecipeResponse represents the response to a recipe modification
type ModifyRecipeResponse struct {
	CurrentRecipe RecipeBody `json:"current_recipe" binding:"required"`
	Diff          RecipeDiff `json:"diff" binding:"required"`
	ResponseText  string     `json:"response_text" binding:"required"`
}

// @Description MealPlanRecipe is a recipe in a meal plan
type MealPlanRecipe struct {
	PlanID string `json:"plan_id" binding:"required"`
	Day    int    `json:"day,omitempty"`
	RecipeBody
}

// @Description MealPlan represents a meal plan
type MealPlan struct {
	ID      string            `json:"id" example:"12345678-1234-1234-1234-123456789012" binding:"required"`
	UserID  string            `json:"user_id" example:"12345678-1234-1234-1234-123456789012" binding:"required"`
	Name    string            `json:"name" example:"My Meal Plan" binding:"required"`
	Recipes []*MealPlanRecipe `json:"recipes" binding:"required"`
}
