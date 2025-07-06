package models

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

// Recipe represents a recipe
// @Description A recipe
type Recipe struct {
	ID               string       `json:"id" example:"12345678-1234-1234-1234-123456789012" binding:"required"`
	Title            string       `json:"title" example:"Veal Bolognese" binding:"required"`
	Description      string       `json:"description" example:"A classic Italian dish" binding:"required"`
	TotalTimeMinutes int          `json:"total_time_minutes" example:"120" binding:"required"`
	Servings         int          `json:"servings" example:"4" binding:"required"`
	Ingredients      []Ingredient `json:"ingredients" binding:"required"`
	Steps            []string     `json:"steps" binding:"required"`
}

// MealPlan represents a meal plan
// @Description A meal plan
type MealPlan struct {
	Recipes []Recipe `json:"recipes" binding:"required"`
}
