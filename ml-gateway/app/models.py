from pydantic import BaseModel

class Ingredient(BaseModel):
    name: str
    quantity: float
    unit: str

class Recipe(BaseModel):
    id: str
    title: str
    description: str
    total_time_minutes: float
    servings: int
    ingredients: list[Ingredient]
    steps: list[str]

class MealPlan(BaseModel):
    recipes: list[Recipe]

class Profile(BaseModel):
    allergies: list[str]

class ChatRequest(BaseModel):
    user_id: str
    message: str
    meal_plan: MealPlan
    profile: Profile

class ChatResponse(BaseModel):
    intent: str
    response_text: str
    new_meal_plan: MealPlan | None = None
    grocery_list: list[dict] | None = None
    needs_clarification: bool = False
