import type api from "@/api";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { useState } from "react";
import { useModifyChat, useSuggestChat } from "../chat/hooks";
import { ChatDrawer } from "../chat/chat-drawer";
import { useParams } from "react-router-dom";
import { useMealPlan } from "./hooks";

export default function MealPlanLayout() {
    const mealPlanId = useParams().id
    if (!mealPlanId) {
        return <h1>Missing meal plan ID</h1>
    }
    const { data: mealPlan, isLoading, error } = useMealPlan(mealPlanId)
    console.log('mealPlan', mealPlan)
    return (
        <>
            {isLoading && <p>Loading meal plan...</p>}
            {error && <p>Error loading meal plan: {error.message}</p>}
            {mealPlan && <MealPlan plan={mealPlan}/>}
        </>
    )
}

interface DrawerState {
    open: boolean;
    mode: "question" | "modify" | "suggest";
    recipe?: api.ModelsRecipe;
}

interface Props {
    plan: api.ModelsMealPlan
}

export function MealPlan({ plan }: Props) {
    const [drawerState, setDrawerState] = useState<DrawerState>({
        open: false,
        mode: "suggest",
        recipe: undefined,
    })

    const openSuggest = () => {
        setDrawerState({
            open: true,
            mode: "suggest",
            recipe: undefined,
        })
    }

    const openModify = (recipe: api.ModelsRecipe) => {
        setDrawerState({
            open: true,
            mode: "modify",
            recipe,
        })
    }

    const closeDrawer = () => {
        setDrawerState({
            open: false,
            mode: "suggest",
            recipe: undefined,
        })
    }

    const {
        mutate: suggestRecipe,
        isPending: suggestRecipePending,
        error: suggestRecipeError
    } = useSuggestChat(plan.id)
    const {
        mutate: modifyRecipe,
        isPending: modifyRecipePending,
        error: modifyRecipeError
    } = useModifyChat(plan.id, drawerState.recipe?.id ?? "")

    const handleSend = (message: string) => {
        switch (drawerState.mode) {
            case "suggest":
                suggestRecipe(message, {
                    onSuccess: () => closeDrawer(),
                })
                break;
            case "modify":
                modifyRecipe(message, {
                    onSuccess: () => closeDrawer(),
                })
                break;
        }
    }

    return (
        <div>
            <div>
                <h1>Meal Plan</h1>
                {plan.recipes.map(r => (
                    <div className="border rounded p-2 mt-2" key={r.id}>
                        <RecipeAccordion recipe={r} />
                        <Button onClick={() => openModify(r)}>Modify</Button>
                    </div>
                ))}
            </div>
            <div className="mt-2">
                <Button onClick={openSuggest}>Add Recipe</Button>
            </div>
            <ChatDrawer
                open={drawerState.open}
                onOpenChange={(open) => setDrawerState({ ...drawerState, open })}
                mealPlanId={plan.id}
                mode={drawerState.mode}
                recipe={drawerState.recipe}
                onSend={handleSend}
                loading={
                    suggestRecipePending ||
                    modifyRecipePending
                }
            />
            <div>
                {suggestRecipeError && <p>{suggestRecipeError.message}</p>}
                {modifyRecipeError && <p>{modifyRecipeError.message}</p>}
            </div>
        </div>
    )
}

interface RecipeProps {
    recipe: api.ModelsRecipe
}

function RecipeAccordion({ recipe }: RecipeProps) {
    console.log('recipe', recipe)
    return (
        <Accordion type="single" collapsible>
            <AccordionItem value={recipe.id}>
                <AccordionTrigger>
                    <div className="flex flex-col space-y-2">
                        <h2>{recipe.title}</h2>
                        <p className="text-muted-foreground">{recipe.description}</p>
                        <div className="flex space-x-2">
                            <div className="flex space-x-2">
                                <p>Servings:</p>
                                <p className="text-muted-foreground">{recipe.servings}</p>
                            </div>
                            <div className="flex space-x-2">
                                <p>Time:</p>
                                <p className="text-muted-foreground">{recipe.total_time_minutes} minutes</p>
                            </div>
                        </div>
                    </div>
                </AccordionTrigger>
                <AccordionContent className="text-left space-y-2">
                    <Separator />
                    <h2 className="font-semibold">Ingredients:</h2>
                    <ul>
                        {recipe.ingredients.map(i => (
                            <li key={i.name} className="flex space-x-2">
                                <p>{i.quantity}</p>
                                <p>{i.unit}</p>
                                <p>{i.name}</p>
                            </li>
                        ))}
                    </ul>
                    <Separator />
                    <h2 className="font-semibold">Steps:</h2>
                    <ol>
                        {recipe.steps.map((s, i) => (
                            <li key={i}>{i + 1}. {s}</li>
                        ))}
                    </ol>
                </AccordionContent>
            </AccordionItem>
        </Accordion>
    )
}