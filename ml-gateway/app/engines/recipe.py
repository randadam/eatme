from .llm import as_json, sys, usr
from models import Recipe, MealPlan, Profile

async def suggest(profile: Profile, message: str) -> MealPlan:
    messages = [
        sys("You are a meal-planning assistant."),
        usr(f"Preferences: {profile}\n\nUser request: \"{message}\"\n\nReturn 2–3 recipes in JSON format.")
    ]
    return await as_json(messages, schema=MealPlan)


async def modify(meal_plan: MealPlan, profile: Profile, message: str) -> MealPlan:
    messages = [
        sys("You are a recipe modifier."),
        usr(f"Preferences: {profile}\n\nUser request: \"{message}\"\n\nCurrent meal plan: {meal_plan}\n\nReturn the modified recipe in JSON format.")
    ]
    return await as_json(messages, schema=MealPlan)