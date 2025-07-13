import type api from "@/api"
import { useSuggestionThread } from "@/features/chat/hooks"
import SuggestionCard from "./suggestion-card"
import { useNavigate } from "react-router-dom"

interface SuggestionThreadProps {
    initialThread: api.ModelsThreadState
}

export default function SuggestionThread({ initialThread }: SuggestionThreadProps) {
    const nav = useNavigate()
    const {
        thread,
        currentSuggestion,
        currentIndex,
        reject,
        accept,
        rejectLoading,
        acceptLoading,
        error,
    } = useSuggestionThread(initialThread)

    console.log('currentIndex', currentIndex)
    console.log('currentSuggestion', currentSuggestion)

    function handleAccept() {
        accept(recipeId => nav(`/recipes/${recipeId}`))
    }

    return (
        <div className="flex flex-col h-full space-y-6">
            <div className="text-center">
                <h2 className="font-semibold">Suggestion For:</h2>
                <p className="text-muted-foreground">{thread.current_prompt}</p>
            </div>
            <div className="flex-1 overflow-y-auto">
                <SuggestionCard
                    suggestion={currentSuggestion}
                    reject={reject}
                    accept={handleAccept}
                    rejectLoading={rejectLoading}
                    acceptLoading={acceptLoading}
                />
            </div>
            <p>{error?.message}</p>
        </div>
    )
}