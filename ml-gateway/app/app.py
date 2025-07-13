from fastapi import FastAPI
from models import RecipeSuggestion, SuggestChatRequest, SuggestChatResponse, ModifyChatRequest, ModifyChatResponse, GeneralChatRequest, GeneralChatResponse
from engines import suggest, modify, answer

app = FastAPI()

@app.post("/chat/suggest", response_model=SuggestChatResponse)
async def chat(req: SuggestChatRequest):
    print("Incoming suggest request:", req)
    try:
        recipes = await suggest(req.profile, req.history, req.message)
        print("Got recipes:", recipes)
        text  = "Here is an idea ↓"
        return SuggestChatResponse(response_text=text,
                                   suggestions=[RecipeSuggestion(recipe=r, response_text=text) for r in recipes])
    except Exception as e:
        print("Error in suggest:", e)
        raise e

@app.post("/chat/modify", response_model=ModifyChatResponse)
async def chat(req: ModifyChatRequest):
    print("Incoming modify request:", req)
    try:
        resp = await modify(req.recipe, req.profile, req.message)
        print("Got updated recipe:", resp)
        return resp
    except Exception as e:
        print("Error in modify:", e)
        raise e

@app.post("/chat/general", response_model=GeneralChatResponse)
async def chat(req: GeneralChatRequest):
    print("Incoming general request:", req)
    try:
        response = await answer(req.recipe, req.profile, req.message)
        print("Got response:", response)
        return GeneralChatResponse(response_text=response)
    except Exception as e:
        print("Error in general:", e)
        raise e