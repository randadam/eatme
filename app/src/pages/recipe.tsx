import type api from "@/api";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { ChatDrawer } from "../features/chat/chat-drawer";
import { useParams } from "react-router-dom";
import { useModifyRecipe, useRecipe } from "../features/recipe/hooks";
import { RecipeAccordion } from "@/features/recipe/recipe-accordion";

export default function RecipePage() {
    const recipeId = useParams().id
    if (!recipeId) {
        return <h1>Missing recipe plan ID</h1>
    }
    const { recipe, isLoading, error } = useRecipe(recipeId)
    return (
        <div>
            <h2 className="text-2xl font-bold pb-6">Recipe Details</h2>
            {isLoading && <p>Loading recipe...</p>}
            {error && <p>Error loading recipe: {error.message}</p>}
            {recipe && <Recipe recipe={recipe.data as api.ModelsUserRecipe}/>}
        </div>
    )
}

interface DrawerState {
    open: boolean;
    mode: "question" | "modify" | "suggest";
    recipe?: api.ModelsUserRecipe;
}

interface Props {
    recipe: api.ModelsUserRecipe
}

export function Recipe({ recipe }: Props) {
    const [drawerState, setDrawerState] = useState<DrawerState>({
        open: false,
        mode: "modify",
        recipe,
    })

    const { modifyRecipe, modifyRecipePending, modifyRecipeError } = useModifyRecipe(recipe.id)

    const openModify = (recipe: api.ModelsUserRecipe) => {
        setDrawerState({
            open: true,
            mode: "modify",
            recipe,
        })
    }

    const closeDrawer = () => {
        setDrawerState({
            open: false,
            mode: "modify",
            recipe,
        })
    }

    const handleSend = (message: string) => {
        switch (drawerState.mode) {
            case "modify":
                modifyRecipe({ message }, { onSuccess: () => closeDrawer() })
                break;
        }
    }

    return (
        <>
            <div className="border rounded p-2">
                <RecipeAccordion recipe={recipe} />
                <Button onClick={() => openModify(recipe)}>Modify</Button>
            </div>
            <ChatDrawer
                open={drawerState.open}
                onOpenChange={(open) => setDrawerState({ ...drawerState, open })}
                mode={drawerState.mode}
                recipe={drawerState.recipe}
                onSend={handleSend}
                loading={modifyRecipePending}
                error={modifyRecipeError?.message}
            />
        </>
    )
}