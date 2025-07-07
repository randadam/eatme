from .llm import as_json, sys, usr
from models import Recipe, Profile

async def suggest(profile: Profile, message: str) -> list[Recipe]:
    messages = [
        sys("You are a meal-planning assistant."),
        usr(f"Preferences: {profile}\n\nUser request: \"{message}\"\n\nReturn a recipe in JSON format.")
    ]
    return await as_json(messages, schema=Recipe)


async def modify(recipe: Recipe, profile: Profile, message: str) -> Recipe:
    messages = [
        sys("You are a recipe modifier."),
        usr(f"Preferences: {profile}\n\nUser request: \"{message}\"\n\nCurrent recipe: {recipe}\n\nReturn the modified recipe in JSON format.")
    ]
    return await as_json(messages, schema=Recipe)