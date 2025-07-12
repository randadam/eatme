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

class Profile(BaseModel):
    allergies: list[str]

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

class GeneralChatRequest(BaseModel):
    message: str
    profile: Profile
    recipe: Recipe

class GeneralChatResponse(BaseModel):
    response_text: str

