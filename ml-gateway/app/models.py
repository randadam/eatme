from pydantic import BaseModel

class Ingredient(BaseModel):
    name: str
    quantity: float
    unit: str

class Recipe(BaseModel):
    id: str
    title: str
    description: str
    total_time_minutes: int
    servings: int
    ingredients: list[Ingredient]
    steps: list[str]

class MealPlan(BaseModel):
    id: str
    recipes: list[Recipe]

class Profile(BaseModel):
    allergies: list[str]

class SuggestChatRequest(BaseModel):
    message: str
    profile: Profile

class SuggestChatResponse(BaseModel):
    response_text: str
    new_recipe: Recipe | None = None

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
    meal_plan: MealPlan

class GeneralChatResponse(BaseModel):
    response_text: str

