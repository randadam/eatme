from .llm import chat, sys, usr
from models import MealPlan, Profile

async def answer(meal_plan: MealPlan, profile: Profile, message: str) -> str:
    messages = [
        sys("You are a helpful cooking assistant. Answer concisely but clearly."),
        usr(f"Preferences: {profile}\n\nMeal plan: {meal_plan}\n\nUser request: \"{message}\"")
    ]
    return await chat(messages)
