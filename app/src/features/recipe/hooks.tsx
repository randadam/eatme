import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import api from "@/api"
import { isHttpApiError } from "@/lib/error/error-provider"

export const recipeKeys = {
    all: ["recipes"] as const,
    byId: (id: string) => [...recipeKeys.all, id] as const,
}

export function useRecipe(recipeId: string) {

    const { data: recipe, isLoading, error } = useQuery({
        queryKey: recipeKeys.byId(recipeId),
        queryFn: async () => {
            const resp = await api.getRecipe(recipeId)
            return resp.data as unknown as api.ModelsUserRecipe
        },
    })

    const notFound = isHttpApiError(error) && error.status === 404

    return { recipe, isLoading, error, notFound }
}

export function useModifyRecipe(recipeId: string) {
    const qc = useQueryClient()

    const { mutate: modifyRecipe, isPending: modifyRecipePending, error: modifyRecipeError } = useMutation({
        mutationFn: async (request: api.ModelsModifyRecipeViaChatRequest) => {
            const resp = await api.modifyRecipe(recipeId, request)
            const result = resp.data as unknown as api.ModelsModifyChatResponse
            if (result.error) {
                throw new Error(result.error)
            }
            return result
        },
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: recipeKeys.byId(recipeId) })
        },
    })

    return { modifyRecipe, modifyRecipePending, modifyRecipeError }
}

export function useAllRecipes() {
    const { data: recipes, isLoading, error } = useQuery({
        queryKey: recipeKeys.all,
        queryFn: async () => {
            const resp = await api.getAllRecipes()
            return resp.data as unknown as api.ModelsUserRecipe[]
        },
    })

    return { recipes, isLoading, error }
}

export function useDeleteRecipe() {
    const qc = useQueryClient()

    const { mutate: deleteRecipe, isPending: deleteRecipePending, error: deleteRecipeError } = useMutation({
        mutationFn: async (recipeId: string) => {
            const resp = await api.deleteRecipe(recipeId)
            return resp.data as unknown as api.ModelsUserRecipe
        },
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: recipeKeys.all })
        },
    })

    return { deleteRecipe, deleteRecipePending, deleteRecipeError }
}