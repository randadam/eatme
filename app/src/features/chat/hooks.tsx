import { useMutation, useQuery } from "@tanstack/react-query"
import api from "@/api"
import { useState } from "react"

export const suggestionThreadKeys = {
    byId: (id: string) => ["suggestion-thread", id] as const,
}

export function useStartSuggestionThread() {
    const { mutate: startThread, isPending: startThreadPending } = useMutation({
        mutationFn: async (message: string) => {
            const resp = await api.startSuggestionThread({ prompt: message })
            return resp.data as api.ModelsThreadState
        },
    })

    return { startThread, startThreadPending }
}

export function useGetSuggestionThread(threadId: string) {
    const { data: thread, isLoading: fetchLoading, error: fetchError } = useQuery({
        queryKey: suggestionThreadKeys.byId(threadId),
        queryFn: async () => {
            const resp = await api.getThread(threadId)
            return resp.data as api.ModelsThreadState
        },
    })

    return { thread, fetchLoading, fetchError }
}


export function useSuggestionThread(initialThread: api.ModelsThreadState) {
    let firstNotSeen = initialThread.suggestions.findIndex(s => !s.accepted && !s.rejected)
    if (firstNotSeen === -1) {
        firstNotSeen = initialThread.suggestions.length - 1
    }

    const [threadState, setThreadState] = useState({
        thread: initialThread,
        currentIndex: firstNotSeen,
    })

    const { mutate: getNextSuggestion, isPending: getNextSuggestionPending, error: getNextSuggestionError } = useMutation({
        mutationFn: async (threadId: string, updatedPrompt?: string) => {
            const req: api.ModelsGetNewSuggestionsRequest = {
                prompt: updatedPrompt,
            }
            const resp = await api.getNewSuggestions(threadId, req)
            return resp.data as api.ModelsRecipeSuggestion[]
        },
        onSuccess: (newSuggestions) => {
            setThreadState(prev => ({
                thread: {
                    ...prev.thread,
                    suggestions: [...prev.thread.suggestions, ...newSuggestions],
                },
                currentIndex: prev.currentIndex + 1,
            }))
        },
    })
    const reject = () => {
        const nextIndex = threadState.currentIndex + 1;
        console.log('nextIndex', nextIndex)
        console.log('threadState.thread.suggestions.length', threadState.thread.suggestions.length)
        if (nextIndex < threadState.thread.suggestions.length) {
            setThreadState(prev => ({
                ...prev,
                currentIndex: nextIndex,
            }))
        } else {
            getNextSuggestion(threadState.thread.id)
        }
    }

    const { mutate: acceptSuggestion, isPending: acceptSuggestionPending, error: acceptSuggestionError } = useMutation({
        mutationFn: async (cb?: (recipeId: string) => void) => {
            const resp = await api.acceptSuggestion(threadState.thread.id, threadState.thread.suggestions[threadState.currentIndex].id)
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