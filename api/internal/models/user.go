package models

// @Description User represents a user in the system
type User struct {
	// User's unique identifier
	ID string `json:"id" example:"usr_123456789"`
	// User's email address
	Email string `json:"email" example:"john.doe@example.com"`
}

type SetupStep string

const (
	SetupStepProfile   SetupStep = "profile"
	SetupStepSkill     SetupStep = "skill"
	SetupStepCuisines  SetupStep = "cuisines"
	SetupStepDiet      SetupStep = "diet"
	SetupStepEquipment SetupStep = "equipment"
	SetupStepAllergies SetupStep = "allergies"
	SetupStepDone      SetupStep = "done"
)

type Skill string

const (
	SkillBeginner     Skill = "beginner"
	SkillIntermediate Skill = "intermediate"
	SkillAdvanced     Skill = "advanced"
	SkillChef         Skill = "chef"
)

// @Description SignupRequest represents the user signup request payload
type SignupRequest struct {
	// User's email address
	Email string `json:"email" example:"john.doe@example.com" binding:"required"`
	// User's password
	Password string `json:"password" example:"Password123!" binding:"required"`
}

// @Description SignupResponse represents the user signup response
type SignupResponse struct {
	// Access token for user
	Token string `json:"token" example:"<JWT_TOKEN>" binding:"required"`
}

// @Description LoginRequest represents the user login request payload
type LoginRequest struct {
	// User's email address
	Email string `json:"email" example:"john.doe@example.com" binding:"required"`
	// User's password
	Password string `json:"password" example:"Password123!" binding:"required"`
}

// @Description LoginResponse represents the user login response
type LoginResponse struct {
	// Access token for user
	Token string `json:"token" example:"<JWT_TOKEN>" binding:"required"`
}

// @Description Profile represents a user's profile information
type Profile struct {
	// Setup Step
	SetupStep SetupStep `json:"setup_step" binding:"required"`
	// User's name
	Name string `json:"name" binding:"required"`
	// User's skill level
	Skill Skill `json:"skill" binding:"required"`
	// User's cuisines
	Cuisines []string `json:"cuisines" binding:"required"`
	// User's diet restrictions
	Diets []string `json:"diets" binding:"required"`
	// User's equipment
	Equipment []string `json:"equipment" binding:"required"`
	// User's allergies
	Allergies []string `json:"allergies" binding:"required"`
}

// @Description ProfileUpdateRequest represents a user's profile update request payload
type ProfileUpdateRequest struct {
	// Setup Step
	SetupStep SetupStep `json:"setup_step"`
	// User's name
	Name string `json:"name,omitempty"`
	// User's skill level
	Skill Skill `json:"skill,omitempty"`
	// User's cuisines
	Cuisines []string `json:"cuisines,omitempty"`
	// User's diet restrictions
	Diets []string `json:"diets,omitempty"`
	// User's equipment
	Equipment []string `json:"equipment,omitempty"`
	// User's allergies
	Allergies []string `json:"allergies,omitempty"`
}
