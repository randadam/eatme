package models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type DietRestriction string

type AllergySeverity string

const (
	SeverityMild         AllergySeverity = "mild"
	SeverityModerate     AllergySeverity = "moderate"
	SeveritySevere       AllergySeverity = "severe"
	SeverityAnaphylactic AllergySeverity = "anaphylactic"
)

type Allergy struct {
	Name        string          `json:"name" example:"peanuts"`
	Severity    AllergySeverity `json:"severity,omitempty"`
	Description string          `json:"description,omitempty" example:"Allergic reaction to peanuts"`
}

const (
	DietOmnivore    DietRestriction = "omnivore"
	DietVegetarian  DietRestriction = "vegetarian"
	DietVegan       DietRestriction = "vegan"
	DietPescatarian DietRestriction = "pescatarian"
	DietKeto        DietRestriction = "keto"
	DietPaleo       DietRestriction = "paleo"
	DietLowFODMAP   DietRestriction = "low_fodmap"
	DietOther       DietRestriction = "other"
)

// Preferences represents a user's dietary preferences
// @Description Preferences for the user's diet and health goals
type Preferences struct {
	UserID string `json:"user_id" example:"usr_123456789"`
	// List of dietary restrictions
	DietRestrictions []DietRestriction `json:"diet_restrictions"`
	// List of allergies
	Allergies []Allergy `json:"allergies"`
	// Health goals
	Goals *Goals `json:"goals,omitempty"`
}

type Goals struct {
	LimitCalories *int    `json:"limit_calories,omitempty"`
	MacrosTarget  *Macros `json:"macros_target,omitempty"`
	Notes         string  `json:"notes,omitempty"`
}

type Macros struct {
	ProteinGrams int `json:"protein_grams,omitempty"`
	CarbsGrams   int `json:"carbs_grams,omitempty"`
	FatGrams     int `json:"fat_grams,omitempty"`
}
