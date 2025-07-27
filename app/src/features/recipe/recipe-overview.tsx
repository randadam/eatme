import type api from "@/api"
import { Skeleton } from "@/components/ui/skeleton"
import { cn } from "@/lib/utils"

export interface RecipeOverviewProps {
    recipe: api.ModelsRecipeBody
    diff?: api.ModelsRecipeDiff
    thumbnail?: boolean
}

export function RecipeOverview({ recipe, diff, thumbnail }: RecipeOverviewProps) {
    const imageUrl = diff?.new_image_url || recipe.image_url
    return (
        <div className="flex flex-col space-y-2">
            <div className="flex space-x-2 items-center">
                {thumbnail && (
                    <div className="h-16 w-16">
                        {imageUrl ? (
                            <img className="h-full w-full" src={imageUrl} alt={recipe.title} />
                        ) : (
                            <Skeleton className="h-full w-full" />
                        )}
                    </div>
                )}
                <div className="flex-1">
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
                </div>
            </div>
            <div>
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
            </div>
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