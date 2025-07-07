from .llm import as_json, sys, usr
from models import Recipe, Profile

def suggest_template(profile: Profile, history: list[str], message: str) -> str:
    return f"""
    Preferences: {profile}
    Previously Rejected: {", ".join(history)}
    User request: "{message}"
    Return a recipe in JSON format. Do not repeat any of the previously rejected recipes.
    """

def modify_template(recipe: Recipe, profile: Profile, message: str) -> str:
    return f"""
    Preferences: {profile}
    User request: "{message}"
    Current recipe: {recipe}
    Return the modified recipe in JSON format.
    """

async def suggest(profile: Profile, history: list[str], message: str) -> list[Recipe]:
    messages = [
        sys("You are a meal-planning assistant."),
        usr(suggest_template(profile, history, message))
    ]
    return await as_json(messages, schema=Recipe)


async def modify(recipe: Recipe, profile: Profile, message: str) -> Recipe:
    messages = [
        sys("You are a recipe modifier."),
        usr(modify_template(recipe, profile, message))
    ]
    return await as_json(messages, schema=Recipe)