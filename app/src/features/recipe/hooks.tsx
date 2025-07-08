import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
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

export function useModifyRecipe(recipeId: string) {
    const qc = useQueryClient()
    const { mutate: modifyRecipe, isPending: modifyRecipePending, error: modifyRecipeError } = useMutation({
        mutationFn: async (recipe: api.ModelsModifyChatRequest) => {
            const resp = await api.modifyRecipe(recipeId, recipe)
            if (resp.status > 299) {
                throw new Error(JSON.stringify(resp.data))
            }
            return resp.data as unknown as api.ModelsUserRecipe
        },
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: recipeKeys.byId(recipeId) })
        }
    })

    return { modifyRecipe, modifyRecipePending, modifyRecipeError }
}

export function useAllRecipes() {
    const { data: recipes, isLoading, error } = useQuery({
        queryKey: recipeKeys.all,
        queryFn: () => api.getAllRecipes(),
    })

    return { recipes, isLoading, error }
}
