import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import api from "@/api"
import { useErrorHandler } from "@/lib/error/error-provider"

export const recipeKeys = {
    all: ["recipes"] as const,
    byId: (id: string) => [...recipeKeys.all, id] as const,
}

export function useRecipe(recipeId: string) {
    const showError = useErrorHandler()

    const { data: recipe, isLoading, error } = useQuery({
        queryKey: recipeKeys.byId(recipeId),
        queryFn: async () => {
            try {
                const resp = await api.getRecipe(recipeId)
                return resp.data as unknown as api.ModelsUserRecipe
            } catch (err: any) {
                showError(err)
                throw err
            }
        },
    })

    return { recipe, isLoading, error }
}

export function useModifyRecipe(recipeId: string) {
    const qc = useQueryClient()
    const showError = useErrorHandler()

    const { mutate: modifyRecipe, isPending: modifyRecipePending, error: modifyRecipeError } = useMutation({
        mutationFn: async (recipe: api.ModelsModifyChatRequest) => {
            const resp = await api.modifyRecipe(recipeId, recipe)
            return resp.data as unknown as api.ModelsUserRecipe
        },
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: recipeKeys.byId(recipeId) })
        },
        onError: (err: any) => {
            showError(err)
        }
    })

    return { modifyRecipe, modifyRecipePending, modifyRecipeError }
}

export function useAllRecipes() {
    const showError = useErrorHandler()

    const { data: recipes, isLoading, error } = useQuery({
        queryKey: recipeKeys.all,
        queryFn: async () => {
            try {
                const resp = await api.getAllRecipes()
                return resp.data as unknown as api.ModelsUserRecipe[]
            } catch (err: any) {
                showError(err)
                throw err
            }
        },
    })

    return { recipes, isLoading, error }
}

export function useDeleteRecipe() {
    const qc = useQueryClient()
    const showError = useErrorHandler()

    const { mutate: deleteRecipe, isPending: deleteRecipePending, error: deleteRecipeError } = useMutation({
        mutationFn: async (recipeId: string) => {
            const resp = await api.deleteRecipe(recipeId)
            return resp.data as unknown as api.ModelsUserRecipe
        },
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: recipeKeys.all })
        },
        onError: (err: any) => {
            showError(err)
        }
    })

    return { deleteRecipe, deleteRecipePending, deleteRecipeError }
}