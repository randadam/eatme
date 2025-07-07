import api from "@/api"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { recipeKeys } from "../recipe/hooks"


export function useSuggestChat() {
    const { mutate: suggestRecipe, isPending, error } = useMutation({
        mutationFn: async (message: string) => {
            const res = await api.suggestRecipe({
                message,
            })
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
            return res.data as api.ModelsSuggestChatResponse
        },
    })

    return { suggestRecipe, isPending, error }
}

export function useModifyChat(recipeId: string) {
    const queryClient = useQueryClient()

    const { mutate: modifyRecipe, isPending, error } = useMutation({
        mutationFn: async (message: string) => {
            const res = await api.modifyRecipe(recipeId, {
                message,
            })
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({
                queryKey: recipeKeys.byId(recipeId),
            })
        },
    })

    return { modifyRecipe, isPending, error }
}

export function useGeneralChat(recipeId: string) {
    const { mutate: generalChat, isPending, error } = useMutation({
        mutationFn: async (message: string) => {
            const res = await api.generalChat(recipeId, {
                message,
            })
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
            return res.data as api.ModelsGeneralChatResponse
        },
    })

    return { generalChat, isPending, error }
}
    