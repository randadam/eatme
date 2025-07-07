import { useMutation, useQuery } from "@tanstack/react-query"
import api from "@/api"
import { useState } from "react"

export const recipeKeys = {
    all: ["recipes"] as const,
    byId: (id: string) => [...recipeKeys.all, id] as const,
}

export const suggestionThreadKeys = {
    byId: (id: string) => ["suggestion-thread", id] as const,
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

export function useStartSuggestionThread() {
    const { mutate: startThread, isPending: startThreadPending, error: startThreadError } = useMutation({
        mutationFn: (message: string) => api.suggestRecipe({ message }),
    })

    return { startThread, startThreadPending, startThreadError }
}

export function useGetSuggestionThread(threadId: string) {
    const { data: thread, isLoading: fetchLoading, error: fetchError } = useQuery({
        queryKey: suggestionThreadKeys.byId(threadId),
        queryFn: async () => {
            const resp = await api.getSuggestionThread(threadId)
            if (resp.status > 299) {
                throw new Error(JSON.stringify(resp.data))
            }
            return resp.data as unknown as api.ModelsSuggestionThread
        }
    })

    return { thread, fetchLoading, fetchError }
}


export function useSuggestionThread(initialThread: api.ModelsSuggestionThread) {
    const [thread, setThread] = useState(initialThread)
    const [index, setIndex] = useState(initialThread.suggestions.length - 1)

    const { mutate: getNextSuggestion, isPending: getNextSuggestionPending, error: getNextSuggestionError } = useMutation({
        mutationFn: async () => {
            const resp = await api.nextRecipeSuggestion(thread.id)
            if (resp.status > 299) {
                throw new Error(JSON.stringify(resp.data))
            }
            return resp.data as unknown as api.ModelsRecipeSuggestion
        },
        onSuccess: (newSuggestion) => {
            setThread(prev => ({
                ...prev,
                suggestions: [...prev.suggestions, newSuggestion],
            }))
            setIndex(prev => prev + 1)
        },
    })
    const reject = () => getNextSuggestion()

    const { mutate: acceptSuggestion, isPending: acceptSuggestionPending, error: acceptSuggestionError } = useMutation({
        mutationFn: async () => {
            const resp = await api.acceptRecipeSuggestion(thread.id, thread.suggestions[index].id)
            if (resp.status > 299) {
                throw new Error(JSON.stringify(resp.data))
            }
            return resp.data as unknown as api.ModelsUserRecipe
        },
    })
    const accept = () => acceptSuggestion()

    const loading = getNextSuggestionPending || acceptSuggestionPending
    const error = getNextSuggestionError || acceptSuggestionError

    return {
        thread,
        currentSuggestion: thread.suggestions[index],
        reject,
        accept,
        loading,
        error,
    }
}
