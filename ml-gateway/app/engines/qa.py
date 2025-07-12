from .llm import chat, sys, usr
from models import Profile, Recipe

async def answer(recipe: Recipe, profile: Profile, message: str) -> str:
    messages = [
        sys("You are a helpful cooking assistant. Answer concisely but clearly."),
        usr(f"Preferences: {profile}\n\nRecipe: {recipe}\n\nUser request: \"{message}\"")
    ]
    return await chat(messages)
