"use client"

import { useParams } from "react-router-dom"
import { Button } from "@/components/ui/button"
import { Loader2 } from "lucide-react"
import { useGetSuggestionThread, useSuggestionThread } from "@/features/recipe/hooks"
import type api from "@/api"

export default function Suggest() {
    const threadId = useParams().threadId
    if (!threadId) {
        return <div>Missing thread ID</div>
    }

    const { thread: initialThread, fetchLoading, fetchError } = useGetSuggestionThread(threadId)

    if (fetchLoading) {
        return <div>Loading...</div>
    }
    if (fetchError) {
        return <div>Error: {fetchError.message}</div>
    }
    if (!initialThread) {
        return <div>Thread not found</div>
    }

    return <SuggestionThread initialThread={initialThread} />
}

interface SuggestionThreadProps {
    initialThread: api.ModelsSuggestionThread
}

function SuggestionThread({ initialThread }: SuggestionThreadProps) {
    console.log('initialThread', initialThread)
    const { thread, reject, accept, loading, error } = useSuggestionThread(initialThread)
    console.log('currentThread', thread)
    return (
        <>
            <div>Suggest</div>
            <div className="text-left">{thread?.original_prompt}</div>
            {thread?.suggestions.map((suggestion, index) => (
                <div className="flex space-x-2" key={index}>
                    <div className="text-left">{suggestion.response_text}</div>
                    <div className="text-left">{suggestion.suggestion.title}</div>
                </div>
            ))}
            <div className="flex space-x-2">
                <Button onClick={() => reject()} disabled={loading}>
                    {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                    Reject
                </Button>
                <Button onClick={() => accept()} disabled={loading}>
                    {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                    Accept
                </Button>
            </div>
            <p>{error?.message}</p>
        </>
    )
}