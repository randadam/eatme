import { useQuery } from "@tanstack/react-query"
import api from "@/api"

export const recipeKeys = {
    all: ["recipes"] as const,
    byId: (id: string) => [...recipeKeys.all, id] as const,
}

export function useRecipe(recipeId: string) {
    const { data: recipe, isLoading, error } = useQuery({
        queryKey: recipeKeys.byId(recipeId),
        queryFn: () => api.getRecipe(recipeId),
    })

    return { recipe, isLoading, error }
}

export function useAllRecipes() {
    const { data: recipes, isLoading, error } = useQuery({
        queryKey: recipeKeys.all,
        queryFn: () => api.getAllRecipes(),
    })

    return { recipes, isLoading, error }
}