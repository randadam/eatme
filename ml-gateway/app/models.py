from pydantic import BaseModel

class Ingredient(BaseModel):
    name: str
    quantity: float
    unit: str

class Recipe(BaseModel):
    title: str
    description: str
    total_time_minutes: int
    servings: int
    ingredients: list[Ingredient]
    steps: list[str]
    image_url: str | None = None

class Profile(BaseModel):
    name: str
    skill: str
    cuisines: list[str]
    diets: list[str]
    allergies: list[str]
    equipment: list[str]

    def __str__(self):
        return f"Profile(name={self.name}, skill_level={self.skill}, cuisines={', '.join(self.cuisines)}, diets={', '.join(self.diets)}, allergies={', '.join(self.allergies)}, equipment={', '.join(self.equipment)})"

class SuggestChatRequest(BaseModel):
    message: str
    history: list[str]
    profile: Profile

class SuggestionRecipes(BaseModel):
    suggestions: list[Recipe]

class RecipeSuggestion(BaseModel):
    recipe: Recipe
    response_text: str

class SuggestChatResponse(BaseModel):
    response_text: str
    suggestions: list[RecipeSuggestion]

class ModifyChatRequest(BaseModel):
    message: str
    profile: Profile
    recipe: Recipe

class ModifyChatResponse(BaseModel):
    response_text: str
    new_recipe: Recipe | None = None
    error: str | None = None

class GeneralChatRequest(BaseModel):
    message: str
    profile: Profile
    recipe: Recipe

class GeneralChatResponse(BaseModel):
    response_text: str

