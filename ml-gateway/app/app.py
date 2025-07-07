from fastapi import FastAPI
from models import ChatRequest, ChatResponse
from engines import classify_intent, recipe, qa

app = FastAPI()

@app.post("/chat", response_model=ChatResponse)
async def chat(req: ChatRequest):
    intent = await classify_intent(req.message)

    if intent == "suggest_recipe":
        plan = await recipe.suggest(req.profile, req.message)
        text  = f"Here are {len(plan.recipes)} ideas ↓"
        return ChatResponse(intent=intent,
                            response_text=text,
                            new_meal_plan=plan)

    if intent == "modify_recipe":
        updated = await recipe.modify(req.meal_plan, req.message)
        text  = "Got it—modified recipe below."
        return ChatResponse(intent=intent,
                            response_text=text,
                            new_meal_plan=updated)

    if intent == "grocery_list":
        text  = "Here’s your merged grocery list 🛒"
        return ChatResponse(intent=intent,
                            response_text=text)

    if intent == "general_question":
        answer = await qa.answer(req.message)
        return ChatResponse(intent=intent,
                            response_text=answer)

    clarifier = ("Just to confirm—would you like new recipes "
                 "or to change one you already picked?")
    return ChatResponse(intent="ambiguous",
                        response_text=clarifier,
                        needs_clarification=True)
