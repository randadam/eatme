import type api from "@/api"

export interface RecipeOverviewProps {
    recipe: api.ModelsRecipeBody
}

export function RecipeOverview({ recipe }: RecipeOverviewProps) {
    return (
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
    )
}