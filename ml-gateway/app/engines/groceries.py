from collections import defaultdict
from models import MealPlan, Profile

def build(meal_plan: MealPlan, profile: Profile) -> list[dict]:
    merged = defaultdict(lambda: {"name": "", "quantity": 0, "unit": ""})
    for recipe in meal_plan.recipes:
        for ing in recipe.ingredients:
            key = ing.name.lower()
            if merged[key]["unit"] in ("", ing.unit):
                merged[key]["unit"] = ing["unit"]
                merged[key]["quantity"] += ing["quantity"]
            merged[key]["name"] = key
    return list(merged.values())
