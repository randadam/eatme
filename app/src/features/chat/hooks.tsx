import api from "@/api"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { mealPlanKeys } from "../plan/hooks"


export function useSuggestChat(mealPlanId: string) {
    const queryClient = useQueryClient()

    return useMutation({
        mutationFn: async (message: string) => {
            const res = await api.suggestRecipe(mealPlanId, {
                message,
            })
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
            return res.data as api.ModelsSuggestChatResponse
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: mealPlanKeys.byId(mealPlanId),
            })
        },
    })
}

export function useModifyChat(mealPlanId: string, recipeId: string) {
    const queryClient = useQueryClient()

    return useMutation({
        mutationFn: async (message: string) => {
            const res = await api.modifyRecipe(mealPlanId, recipeId, {
                message,
            })
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: mealPlanKeys.byId(mealPlanId),
            })
        },
    })
}

export function useGeneralChat(mealPlanId: string) {
    return useMutation({
        mutationFn: async (message: string) => {
            const res = await api.generalChat(mealPlanId, {
                message,
            })
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
            return res.data as api.ModelsGeneralChatResponse
        },
    })
}
    