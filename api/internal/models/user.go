package models

// User represents a user in the system
// @Name User
// @Description User information
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

type Cuisine string

const (
	CuisineAmerican   Cuisine = "american"
	CuisineBritish    Cuisine = "british"
	CuisineChinese    Cuisine = "chinese"
	CuisineFrench     Cuisine = "french"
	CuisineGerman     Cuisine = "german"
	CuisineIndian     Cuisine = "indian"
	CuisineItalian    Cuisine = "italian"
	CuisineJapanese   Cuisine = "japanese"
	CuisineMexican    Cuisine = "mexican"
	CuisineSpanish    Cuisine = "spanish"
	CuisineThai       Cuisine = "thai"
	CuisineVietnamese Cuisine = "vietnamese"
)

type Diet string

const (
	DietVegetarian  Diet = "vegetarian"
	DietVegan       Diet = "vegan"
	DietKeto        Diet = "keto"
	DietPaleo       Diet = "paleo"
	DietLowCarb     Diet = "low_carb"
	DietHighProtein Diet = "high_protein"
)

type Equipment string

const (
	EquipmentStove          Equipment = "stove"
	EquipmentOven           Equipment = "oven"
	EquipmentMicrowave      Equipment = "microwave"
	EquipmentToaster        Equipment = "toaster"
	EquipmentGrill          Equipment = "grill"
	EquipmentSmoker         Equipment = "smoker"
	EquipmentSlowCooker     Equipment = "slow_cooker"
	EquipmentPressureCooker Equipment = "pressure_cooker"
	EquipmentSousVide       Equipment = "sous_vide"
)

type Allergy string

const (
	AllergyDairy    Allergy = "dairy"
	AllergyEggs     Allergy = "eggs"
	AllergyFish     Allergy = "fish"
	AllergyGluten   Allergy = "gluten"
	AllergyPeanuts  Allergy = "peanuts"
	AllergySoy      Allergy = "soy"
	AllergyTreeNuts Allergy = "tree_nuts"
	AllergyWheat    Allergy = "wheat"
)

// SignupRequest represents the user signup request payload
// @Name SignupRequest
// @Description User signup request
type SignupRequest struct {
	// User's email address
	Email string `json:"email" example:"john.doe@example.com" binding:"required"`
	// User's password
	Password string `json:"password" example:"Password123!" binding:"required"`
}

// SignupResponse represents the user signup response
// @Name SignupResponse
// @Description User signup response containing the new user's ID
type SignupResponse struct {
	// Access token for user
	Token string `json:"token" example:"<JWT_TOKEN>" binding:"required"`
}

// Profile represents a user's profile information
// @Name Profile
// @Description User profile information
type Profile struct {
	// Setup Step
	SetupStep SetupStep `json:"setup_step" binding:"required"`
	// User's name
	Name string `json:"name" binding:"required"`
	// User's skill level
	Skill Skill `json:"skill" binding:"required"`
	// User's cuisines
	Cuisines []Cuisine `json:"cuisines" binding:"required"`
	// User's diet restrictions
	Diet []Diet `json:"diet" binding:"required"`
	// User's equipment
	Equipment []Equipment `json:"equipment" binding:"required"`
	// User's allergies
	Allergies []Allergy `json:"allergies" binding:"required"`
}

// ProfileUpdateRequest represents a user's profile update request payload
// @Name ProfileUpdateRequest
// @Description User profile update request
type ProfileUpdateRequest struct {
	// Setup Step
	SetupStep SetupStep `json:"setup_step" binding:"required"`
	// User's name
	Name string `json:"name,omitempty"`
	// User's skill level
	Skill Skill `json:"skill,omitempty"`
	// User's cuisines
	Cuisines []Cuisine `json:"cuisines,omitempty"`
	// User's diet restrictions
	Diet []Diet `json:"diet,omitempty"`
	// User's equipment
	Equipment []Equipment `json:"equipment,omitempty"`
	// User's allergies
	Allergies []Allergy `json:"allergies,omitempty"`
}
