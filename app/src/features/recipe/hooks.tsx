import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import api from "@/api"
import { isHttpApiError, useErrorHandler } from "@/lib/error/error-provider"
import { useState } from "react"

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
    const showError = useErrorHandler()

    const [proposedRecipe, setProposedRecipe] = useState<api.ModelsRecipeBody | undefined>(undefined)

    const { mutate: modifyRecipe, isPending: modifyRecipePending, error: modifyRecipeError } = useMutation({
        mutationFn: async (request: api.ModelsModifyRecipeViaChatRequest) => {
            const resp = await api.modifyRecipe(recipeId, request)
            const result = resp.data as unknown as api.ModelsModifyChatResponse
            if (result.error) {
                throw new Error(result.error)
            }
            return result
        },
        onSuccess: (result: api.ModelsModifyChatResponse) => {
            setProposedRecipe(result.new_recipe)
        },
        onError: (error) => {
            console.error("Failed to modify recipe", error)
            showError("Failed to modify recipe")
        },
    })

    const { mutate: reject, isPending: rejectPending, error: rejectError } = useMutation({
        mutationFn: async () => {
            await api.rejectRecipeModification(recipeId)
        },
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: recipeKeys.byId(recipeId) })
            setProposedRecipe(undefined)
        },
        onError: (error) => {
            console.error("Failed to reject recipe modification", error)
            showError("Failed to reject recipe modification")
        },
    })

    const { mutate: accept, isPending: acceptPending, error: acceptError } = useMutation({
        mutationFn: async () => {
            await api.acceptRecipeModification(recipeId)
        },
        onSuccess: () => {
            qc.invalidateQueries({ queryKey: recipeKeys.byId(recipeId) })
            setProposedRecipe(undefined)
        },
        onError: (error) => {
            console.error("Failed to accept recipe modification", error)
            showError("Failed to accept recipe modification")
        }
    })

    return {
        modifyRecipe,
        modifyRecipePending,
        modifyRecipeError,
        reject,
        rejectPending,
        rejectError,
        accept,
        acceptPending,
        acceptError,
        proposedRecipe,
    }
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