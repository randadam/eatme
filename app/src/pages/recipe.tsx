import type api from "@/api";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { useState } from "react";
import { ChatDrawer } from "../features/chat/chat-drawer";
import { useParams } from "react-router-dom";
import { useRecipe } from "../features/recipe/hooks";

export default function RecipePage() {
    const recipeId = useParams().id
    if (!recipeId) {
        return <h1>Missing recipe plan ID</h1>
    }
    const { recipe, isLoading, error } = useRecipe(recipeId)
    return (
        <>
            {isLoading && <p>Loading recipe...</p>}
            {error && <p>Error loading recipe: {error.message}</p>}
            {recipe && <Recipe recipe={recipe.data as api.ModelsUserRecipe}/>}
        </>
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
                console.log(message)
                break;
        }
    }

    return (
        <>
            <div className="border rounded p-2 mt-2">
                <RecipeAccordion recipe={recipe} />
                <Button onClick={() => openModify(recipe)}>Modify</Button>
            </div>
            <ChatDrawer
                open={drawerState.open}
                onOpenChange={(open) => setDrawerState({ ...drawerState, open })}
                mode={drawerState.mode}
                recipe={drawerState.recipe}
                onSend={handleSend}
                loading={false}
            />
        </>
    )
}

interface RecipeProps {
    recipe: api.ModelsUserRecipe
}

function RecipeAccordion({ recipe }: RecipeProps) {
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