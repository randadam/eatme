import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
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
    const [threadState, setThreadState] = useState({
        thread: initialThread,
        currentIndex: initialThread.suggestions.length - 1,
    })

    const { mutate: getNextSuggestion, isPending: getNextSuggestionPending, error: getNextSuggestionError } = useMutation({
        mutationFn: async () => {
            const resp = await api.nextRecipeSuggestion(threadState.thread.id)
            if (resp.status > 299) {
                throw new Error(JSON.stringify(resp.data))
            }
            return resp.data as unknown as api.ModelsRecipeSuggestion
        },
        onSuccess: (newSuggestion) => {
            setThreadState(prev => ({
                ...prev,
                suggestions: [...prev.thread.suggestions, newSuggestion],
                currentIndex: prev.thread.suggestions.length,
            }))
        },
    })
    const reject = () => getNextSuggestion()

    const { mutate: acceptSuggestion, isPending: acceptSuggestionPending, error: acceptSuggestionError } = useMutation({
        mutationFn: async (cb?: (recipeId: string) => void) => {
            const resp = await api.acceptRecipeSuggestion(threadState.thread.id, threadState.thread.suggestions[threadState.currentIndex].id)
            if (resp.status > 299) {
                throw new Error(JSON.stringify(resp.data))
            }
            const respData = resp.data as unknown as api.ModelsUserRecipe
            if (cb) {
                cb(respData.id)
            }
            return respData
        },
    })
    const accept = (cb?: (recipeId: string) => void) => acceptSuggestion(cb)

    const error = getNextSuggestionError || acceptSuggestionError

    const back = () => {
        setThreadState(prev => ({
            ...prev,
            currentIndex: prev.currentIndex - 1,
        }))
    }
    const forward = () => {
        setThreadState(prev => {
            if (prev.currentIndex < prev.thread.suggestions.length - 1) {
                return {
                    ...prev,
                    currentIndex: prev.currentIndex + 1,
                }
            }
            return prev
        })
    }

    return {
        thread: threadState.thread,
        currentSuggestion: threadState.thread.suggestions[threadState.currentIndex],
        currentIndex: threadState.currentIndex,
        reject,
        rejectLoading: getNextSuggestionPending,
        accept,
        acceptLoading: acceptSuggestionPending,
        error,
        back,
        forward,
    }
}
