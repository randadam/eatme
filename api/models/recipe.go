package models

import "time"

type MeasurementUnit string

const (
	MeasurementUnitGram       MeasurementUnit = "g"
	MeasurementUnitMilliliter MeasurementUnit = "ml"
	MeasurementUnitTeaspoon   MeasurementUnit = "tsp"
	MeasurementUnitTablespoon MeasurementUnit = "tbsp"
	MeasurementUnitCup        MeasurementUnit = "cup"
	MeasurementUnitOunce      MeasurementUnit = "oz"
	MeasurementUnitPound      MeasurementUnit = "lb"
)

// Ingredient represents an ingredient in a recipe
type Ingredient struct {
	Name     string          `json:"name" example:"Flour" binding:"required"`
	Quantity float64         `json:"quantity" example:"1" binding:"required"`
	Unit     MeasurementUnit `json:"unit" example:"cup" binding:"required"`
}

// RecipeBody represents the contents of a recipe
// @Description Contents of a recipe
type RecipeBody struct {
	Title            string       `json:"title" example:"Veal Bolognese" binding:"required"`
	Description      string       `json:"description" example:"A classic Italian dish" binding:"required"`
	Ingredients      []Ingredient `json:"ingredients" binding:"required"`
	Steps            []string     `json:"steps" binding:"required"`
	Servings         int          `json:"servings" example:"4" binding:"required"`
	TotalTimeMinutes int          `json:"total_time_minutes" example:"120" binding:"required"`
}

type RecipeSource string

const (
	RecipeSourceScraped   RecipeSource = "scraped"
	RecipeSourceGenerated RecipeSource = "generated"
)

// GlobalRecipe represents a global recipe that can be "forked" into a user's recipe book
// @Description A global recipe that can be "forked" into a user's recipe book
type GlobalRecipe struct {
	ID         string       `json:"id" binding:"required"`
	SourceType RecipeSource `json:"source_type" binding:"required"`
	CreatedAt  time.Time    `json:"created_at" binding:"required"`
	UpdatedAt  time.Time    `json:"updated_at" binding:"required"`
	RecipeBody RecipeBody
}

// UserRecipe is the user's personal copy (favorites, edits).
// @Description The user's personal copy (favorites, edits).
type UserRecipe struct {
	ID              string    `json:"id" binding:"required"`
	UserID          string    `json:"user_id" binding:"required"`
	GlobalRecipeID  *string   `json:"global_recipe_id,omitempty"`
	IsFavorite      bool      `json:"is_favorite" binding:"required"`
	LatestVersionID string    `json:"latest_version_id" binding:"required"`
	CreatedAt       time.Time `json:"created_at" binding:"required"`
	UpdatedAt       time.Time `json:"updated_at" binding:"required"`
	RecipeBody
}

// RecipeVersion is an immutable snapshot used inside meal plans.
// @Description An immutable snapshot used inside meal plans.
type RecipeVersion struct {
	ID           string    `json:"id" binding:"required"`
	UserRecipeID string    `json:"user_recipe_id" binding:"required"`
	ParentID     *string   `json:"parent_id,omitempty"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
	Notes        *string   `json:"notes,omitempty"`
	RecipeBody
}

// MealPlanRecipe is a recipe in a meal plan
// @Description A recipe in a meal plan
type MealPlanRecipe struct {
	PlanID string `json:"plan_id" binding:"required"`
	Day    int    `json:"day,omitempty"`
	RecipeBody
}

// MealPlan represents a meal plan
// @Description A meal plan
type MealPlan struct {
	ID      string            `json:"id" example:"12345678-1234-1234-1234-123456789012" binding:"required"`
	UserID  string            `json:"user_id" example:"12345678-1234-1234-1234-123456789012" binding:"required"`
	Name    string            `json:"name" example:"My Meal Plan" binding:"required"`
	Recipes []*MealPlanRecipe `json:"recipes" binding:"required"`
}
