from fastapi import FastAPI
from models import SuggestChatRequest, SuggestChatResponse, ModifyChatRequest, ModifyChatResponse, GeneralChatRequest, GeneralChatResponse
from engines import suggest, modify, answer

app = FastAPI()

@app.post("/chat/suggest", response_model=SuggestChatResponse)
async def chat(req: SuggestChatRequest):
    print("Incoming suggest request:", req)
    recipes = await suggest(req.profile, req.history, req.message)
    text  = "Here is an idea ↓"
    return SuggestChatResponse(response_text=text,
                               suggestions=recipes)

@app.post("/chat/modify", response_model=ModifyChatResponse)
async def chat(req: ModifyChatRequest):
    print("Incoming modify request:", req)
    updated = await modify(req.recipe, req.profile, req.message)
    text  = "Got it—modified recipe below."
    return ModifyChatResponse(response_text=text,
                              new_recipe=updated)

@app.post("/chat/general", response_model=GeneralChatResponse)
async def chat(req: GeneralChatRequest):
    print("Incoming general request:", req)
    response = await answer(req.meal_plan, req.profile, req.message)
    return GeneralChatResponse(response_text=response)