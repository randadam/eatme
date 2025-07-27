from uuid import uuid4
from .llm import as_json, sys, usr
import replicate
from models import ModifyChatResponse, Recipe, Profile, SuggestionRecipes

def suggest_template(profile: Profile, history: list[str], message: str, num_suggestions: int) -> str:
    return f"""
    Profile: {profile}
    Previously Rejected: {", ".join(history)}
    User request: "{message}"
    Return {num_suggestions} recipes in JSON format. Do not repeat any of the previously rejected recipes.
    All recipes should be wrapped in a top-level `suggestions` field.
    """

def modify_template(recipe: Recipe, profile: Profile, message: str) -> str:
    return f"""
    Profile: {profile}
    User request: "{message}"
    Current recipe: {recipe}
    Return the modified recipe in JSON format following the given schema, which will have `new_recipe`,
    `response_text`, and an optional `error` field.  If the recipe modification violates a rule from
    the user's profile (e.g. allergies), return the original recipe with a description of the violation
    in the `error` field.
    """

def image_template(recipe: Recipe) -> str:
    return f"""
    Recipe: {recipe}
    Generate a high quality photo of the recipe.
    """

async def suggest(profile: Profile, history: list[str], message: str, num_suggestions: int = 3) -> list[Recipe]:
    messages = [
        sys("You are a recipe suggestion assistant."),
        usr(suggest_template(profile, history, message, num_suggestions))
    ]
    print(f"Suggest messages: {messages}")
    result = await as_json(messages, schema=SuggestionRecipes)
    return result.suggestions


async def modify(recipe: Recipe, profile: Profile, message: str) -> ModifyChatResponse:
    messages = [
        sys("You are a recipe modifier."),
        usr(modify_template(recipe, profile, message))
    ]
    print(f"Modify messages: {messages}")
    return await as_json(messages, schema=ModifyChatResponse)

async def generate_image(recipe: Recipe) -> str:
    outputs = await replicate.async_run(
        "black-forest-labs/flux-dev",
        input={
            "prompt": image_template(recipe),
        }
    )
    output = outputs[0]
    id = uuid4()
    with open(f"/images/{id}.png", "wb") as f:
        f.write(output.read())
    return id