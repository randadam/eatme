import type api from "@/api"
import { Separator } from "@/components/ui/separator"
import { RecipeOverview } from "./recipe-overview"

export interface FullRecipeProps {
    recipe: api.ModelsRecipeBody
}

export function FullRecipe({ recipe }: FullRecipeProps) {
    return (
        <div>
            <div className="p-2">
                <RecipeOverview recipe={recipe} />
            </div>
            <Separator />
            <div className="p-2">
                <h2 className="font-semibold">Ingredients:</h2>
                <ul className="text-left">
                    {recipe.ingredients.map(i => (
                        <li key={i.name} className="flex space-x-2">
                            <p>{i.quantity}</p>
                            <p>{i.unit}</p>
                            <p>{i.name}</p>
                        </li>
                    ))}
                </ul>
            </div>
            <Separator />
            <div className="p-2">
                <h2 className="font-semibold">Steps:</h2>
                <ol className="text-left">
                    {recipe.steps.map((s, i) => (
                        <li key={i}>{i + 1}. {s}</li>
                    ))}
                </ol>
            </div>
        </div>
    )
}