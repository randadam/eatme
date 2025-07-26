import type api from "@/api"
import { cn } from "@/lib/utils"

export interface RecipeOverviewProps {
    recipe: api.ModelsRecipeBody
    diff?: api.ModelsRecipeDiff
}

export function RecipeOverview({ recipe, diff }: RecipeOverviewProps) {
    console.log('recipe', recipe)
    console.log('diff', diff)
    return (
        <div className="flex flex-col space-y-2">
            <h2
                className={cn(
                    "font-semibold",
                    diff?.new_title && "line-through decoration-red-500"
                )}
            >
                {recipe.title}
            </h2>
            {diff?.new_title && (
                <h2 className="font-semibold">{diff.new_title}</h2>
            )}
            <p className={cn(
                "text-muted-foreground",
                diff?.new_description && "line-through decoration-red-500"
            )}>
                {recipe.description}
            </p>
            {diff?.new_description && (
                <p className="text-muted-foreground decoration-red-500">
                    {diff.new_description}
                </p>
            )}
            <div className="flex space-x-2">
                <div className="flex space-x-2">
                    <p>Servings:</p>
                    <p className={cn(
                        "text-muted-foreground",
                        diff?.new_servings && "line-through decoration-red-500"
                    )}>{recipe.servings}</p>
                    {diff?.new_servings && (
                        <p className="text-muted-foreground decoration-red-500">{diff.new_servings}</p>
                    )}
                </div>
                <div className="flex space-x-2">
                    <p>Total Time:</p>
                    <p className={cn(
                        "text-muted-foreground",
                        diff?.new_total_time_minutes && "line-through decoration-red-500"
                    )}>{recipe.total_time_minutes} minutes</p>
                    {diff?.new_total_time_minutes && (
                        <p className="text-muted-foreground decoration-red-500">{diff.new_total_time_minutes} minutes</p>
                    )}
                </div>
            </div>
        </div>
    )
}