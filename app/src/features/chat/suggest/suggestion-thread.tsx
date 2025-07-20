import type api from "@/api"
import { useSuggestionThread } from "@/features/chat/hooks"
import SuggestionCard from "./suggestion-card"
import { useNavigate } from "react-router-dom"
import { Button } from "@/components/ui/button"
import { ArrowLeft, Undo } from "lucide-react"

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
        back,
        rejectLoading,
        acceptLoading,
    } = useSuggestionThread(initialThread)

    console.log('currentIndex', currentIndex)
    console.log('currentSuggestion', currentSuggestion)

    function handleAccept() {
        accept(recipeId => nav(`/recipes/${recipeId}`))
    }

    function handleBack() {
        nav('/recipes')
    }

    return (
        <div className="flex flex-col h-full space-y-6">
            <div className="flex justify-between items-center">
                <div className="items-center gap-2 w-1/8">
                    <Button onClick={handleBack} variant="ghost" size="icon" className="rounded-full" aria-label="Back">
                        <ArrowLeft className="h-4 w-4" />
                    </Button>
                </div>
                <div className="text-center">
                    <h2 className="font-semibold">Suggestion For:</h2>
                    <p className="text-muted-foreground">{thread.current_prompt}</p>
                </div>
                <div className="items-center gap-2 w-1/8">
                    {currentIndex > 0 && (
                        <Button variant="ghost" size="icon" className="rounded-full" aria-label="Undo" onClick={back}>
                            <Undo className="h-4 w-4" />
                        </Button>
                    )}
                </div>
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
        </div>
    )
}