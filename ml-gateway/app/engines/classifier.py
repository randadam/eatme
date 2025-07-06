from .llm import chat, sys, usr

TEMPLATE = """
You are MealPlan-Router-Bot. Return only:
suggest_recipe | modify_recipe | grocery_list | general_question | ambiguous
<examples>%s</examples>
"""

EXAMPLES = """
User: "Give me something spicy"
Assistant: suggest_recipe
User: "Swap chicken for tofu"
Assistant: modify_recipe
User: "What is gochujang?"
Assistant: general_question
User: "What do I need to buy?"
Assistant: grocery_list
User: "Maybe tacos, but can we make them vegetarian?"
Assistant: ambiguous
"""

async def classify_intent(msg: str) -> str:
    messages = [
        sys(TEMPLATE % (EXAMPLES)),
        usr(msg),
    ]
    return (await chat(messages)).strip().lower()
