import type api from "@/api";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { ChatDrawer } from "../features/chat/chat-drawer";
import { useNavigate, useParams } from "react-router-dom";
import { useModifyRecipe, useRecipe } from "../features/recipe/hooks";
import { RecipeAccordion } from "@/features/recipe/recipe-accordion";
import { NotFoundPage } from "./not-found";
import { ErrorPage } from "./error";
import RecipeSkeleton from "@/features/recipe/recipe-skeleton";

export default function RecipePage() {
    const recipeId = useParams().id
    return (
        <div>
            <h2 className="text-2xl font-bold pb-6">Recipe Details</h2>
            <Recipe recipeId={recipeId!}/>
        </div>
    )
}

interface DrawerState {
    open: boolean;
    mode: "question" | "modify" | "suggest";
    recipe?: api.ModelsUserRecipe;
}

interface Props {
    recipeId: string
}

export function Recipe({ recipeId }: Props) {
    const nav = useNavigate()
    const { recipe, isLoading, error, notFound } = useRecipe(recipeId)

    const [drawerState, setDrawerState] = useState<DrawerState>({
        open: false,
        mode: "modify",
        recipe,
    })

    const { modifyRecipe, modifyRecipePending, modifyRecipeError } = useModifyRecipe(recipe?.id ?? "")

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
                modifyRecipe({ prompt: message }, { onSuccess: () => closeDrawer() })
                break;
        }
    }

    if (notFound) {
        return <NotFoundPage />
    }

    if (error) {
        return <ErrorPage title="Error loading recipe" description={error.message} />
    }

    if (isLoading || !recipe) {
        return <RecipeSkeleton />
    }

    return (
        <>
            <div className="border rounded p-2">
                <RecipeAccordion id={recipe.id} recipe={recipe} />
                <div className="flex justify-between">
                    <Button variant="outline"onClick={() => openModify(recipe)}>Modify</Button>
                    <Button onClick={() => nav(`/cook/${recipe.id}`)}>Start Cooking</Button>
                </div>
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