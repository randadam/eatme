from fastapi import FastAPI
from models import RecipeSuggestion, SuggestChatRequest, SuggestChatResponse, ModifyChatRequest, ModifyChatResponse, GeneralChatRequest, GeneralChatResponse
from engines import suggest, modify, answer, generate_image

app = FastAPI()

@app.post("/chat/suggest", response_model=SuggestChatResponse)
async def chat(req: SuggestChatRequest):
    print("Incoming suggest request:", req)
    try:
        recipes = await suggest(req.profile, req.history, req.message)
        print("Got recipes:", recipes)
        for recipe in recipes:
            image_id = await generate_image(recipe)
            print(f"Generated image for recipe: {recipe.title} with ID {image_id}")
            recipe.image_url = f"http://localhost:8080/images/{image_id}.png"
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
        if resp.generate_new_image:
            print("Generating new image for recipe")
            image_id = await generate_image(resp.new_recipe)
            print(f"Generated image for recipe: {resp.new_recipe.title} with ID {image_id}")
            resp.new_recipe.image_url = f"http://localhost:8080/images/{image_id}.png"
        else:
            resp.new_recipe.image_url = req.recipe.image_url
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