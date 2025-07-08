import type api from "@/api"
import { useSuggestionThread } from "@/features/chat/hooks"
import SuggestionCard from "./suggestion-card"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"

interface SuggestionThreadProps {
    initialThread: api.ModelsSuggestionThread
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
        back,
        forward,
    } = useSuggestionThread(initialThread)

    console.log('currentIndex', currentIndex)
    console.log('currentSuggestion', currentSuggestion)

    function handleAccept() {
        accept(recipeId => nav(`/recipes/${recipeId}`))
    }

    return (
        <div className="space-y-6">
            <div className="text-center">
                <h2 className="font-semibold">Suggestion For:</h2>
                <p className="text-muted-foreground">{thread.original_prompt}</p>
            </div>
            <SuggestionCard
                suggestion={currentSuggestion}
                reject={reject}
                accept={handleAccept}
                rejectLoading={rejectLoading}
                acceptLoading={acceptLoading}
            />
            <p>{error?.message}</p>
            <div className="flex justify-between">
                <Button onClick={back} disabled={currentIndex === 0}>
                    Back
                </Button>
                <Button onClick={forward} disabled={currentIndex === thread.suggestions.length - 1}>
                    Forward
                </Button>
            </div>
        </div>
    )
}