import type api from "@/api"
import { Separator } from "@/components/ui/separator"
import { RecipeOverview } from "./recipe-overview"
import { cn } from "@/lib/utils"

export interface FullRecipeProps {
    recipe: api.ModelsRecipeBody
    diff?: api.ModelsRecipeDiff
}

interface ModifiedIngredient extends api.ModelsIngredient {
    action: "add" | "remove" | "update" | "keep"
}

export function FullRecipe({ recipe, diff }: FullRecipeProps) {
    const ingredients: ModifiedIngredient[] = recipe.ingredients.map(i => ({
        ...i,
        action: "keep"
    }))
    for (const modifiedIngredient of diff?.modified_ingredients ?? []) {
        ingredients[modifiedIngredient.index] = {
            ...modifiedIngredient,
            action: "update"
        }
    }
    for (const removedIngredient of diff?.removed_ingredients ?? []) {
        ingredients[removedIngredient.index].action = "remove"
    }
    for (const addedIngredient of diff?.added_ingredients ?? []) {
        ingredients.push({
            ...addedIngredient,
            action: "add"
        })
    }

    const steps: api.ModelsDiffStep[] = diff?.new_steps ? diff.new_steps : recipe.steps.map(s => ({
        step: s,
        is_new: false
    }))

    return (
        <div>
            <div className="flex p-2">
                <RecipeOverview recipe={recipe} diff={diff} thumbnail />
            </div>
            <Separator />
            <div className="p-2">
                <h2 className="font-semibold">Ingredients:</h2>
                <ul className="text-left">
                    {ingredients.map(i => {
                        return (
                            <li
                                key={i.name}
                                className="flex space-x-2"
                            >
                                {diff && (
                                    <div className="min-w-4">
                                        {i.action === "add" && <p className="text-green-500">+</p>}
                                        {i.action === "remove" && <p className="text-red-500">-</p>}
                                        {i.action === "update" && <p className="text-yellow-500">~</p>}
                                        {i.action === "keep" && <p>&nbsp;</p>}
                                    </div>
                                )}
                                <p
                                    className={cn(
                                        "flex",
                                        i.action === "remove" && "line-through decoration-red-500",
                                        i.action === "update" || i.action === "add" && "font-semibold"
                                    )}
                                >
                                    {i.quantity} {i.unit} {i.name}
                                </p>
                            </li>
                        )
                    })}
                </ul>
            </div>
            <Separator />
            <div className="p-2">
                <h2 className="font-semibold">Steps:</h2>
                <ol className="text-left">
                    {steps.map((s, i) => {
                        return (
                            <li className="flex space-x-2" key={i}>
                                {diff && (
                                    <div className="min-w-4">
                                        {s.is_new && <p className="text-yellow-500">~</p>}
                                    </div>
                                )}
                                <p>{i + 1}</p>
                                <p className={s.is_new ? "font-semibold" : ""}>
                                    {s.step}
                                </p>
                            </li>
                        )
                    })}
                </ol>
            </div>
        </div>
    )
}