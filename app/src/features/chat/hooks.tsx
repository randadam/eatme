import { useMutation, useQuery } from "@tanstack/react-query"
import api from "@/api"
import { useState } from "react"

export const suggestionThreadKeys = {
    byId: (id: string) => ["suggestion-thread", id] as const,
}

export function useStartSuggestionThread() {
    const { mutate: startThread, isPending: startThreadPending } = useMutation({
        mutationFn: async (message: string) => {
            const resp = await api.suggestRecipe({ message })
            return resp.data as api.ModelsSuggestChatResponse
        },
    })

    return { startThread, startThreadPending }
}

export function useGetSuggestionThread(threadId: string) {
    const { data: thread, isLoading: fetchLoading, error: fetchError } = useQuery({
        queryKey: suggestionThreadKeys.byId(threadId),
        queryFn: async () => {
            const resp = await api.getSuggestionThread(threadId)
            return resp.data as api.ModelsSuggestionThread
        },
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
            return resp.data as api.ModelsRecipeSuggestion
        },
        onSuccess: (newSuggestion) => {
            setThreadState(prev => ({
                thread: {
                    ...prev.thread,
                    suggestions: [...prev.thread.suggestions, newSuggestion],
                },
                currentIndex: prev.thread.suggestions.length,
            }))
        },
    })
    const reject = () => getNextSuggestion()

    const { mutate: acceptSuggestion, isPending: acceptSuggestionPending, error: acceptSuggestionError } = useMutation({
        mutationFn: async (cb?: (recipeId: string) => void) => {
            const resp = await api.acceptRecipeSuggestion(threadState.thread.id, threadState.thread.suggestions[threadState.currentIndex].id)
            const respData = resp.data as api.ModelsUserRecipe
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