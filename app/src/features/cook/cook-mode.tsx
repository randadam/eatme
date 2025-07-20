import type api from "@/api";
import { useCookMode } from "./hooks";
import { Button } from "@/components/ui/button";
import { ArrowLeft, ChevronLeft, ChevronRight, List } from "lucide-react"
import { Progress } from "@/components/ui/progress"
import { useNavigate } from "react-router-dom"
import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle } from "@/components/ui/sheet";
import { ChatBody, type ChatItem } from "../chat/chat-body";
import FocusedLayout from "@/layouts/focused-layout";
import BottomSheet from "@/components/shared/bottom-sheet";
import { useState } from "react";

interface CookModeProps {
    id: string
    recipe: api.ModelsRecipeBody
    chatHistory: ChatItem[]
    askQuestion: (message: string) => void
    askQuestionPending: boolean
    askQuestionError?: string
}

export default function CookMode({ id, recipe, chatHistory, askQuestion, askQuestionPending, askQuestionError }: CookModeProps) {
    const nav = useNavigate()
    const [showChat, setShowChat] = useState(false)
    const { currentStep, stepIdx, goNext, goPrev, ingredients, ingredientsOpen, toggleIngredients } = useCookMode(recipe)

    return (
        <FocusedLayout>
            <div className="flex flex-col h-full">
                <CookModeHeader
                    stepIdx={stepIdx}
                    total={recipe.steps.length}
                    title={recipe.title}
                    onBack={() => nav(`/recipes/${id}`)}
                />
                <div className="h-full flex flex-col justify-between">
                    <CookModeInstructions currentStep={currentStep} stepIdx={stepIdx} />
                    <CookModeControls
                        stepIdx={stepIdx}
                        total={recipe.steps.length}
                        onBack={goPrev}
                        onNext={goNext}
                        toggleIngredients={toggleIngredients}
                    />
                </div>
                <Sheet open={ingredientsOpen} onOpenChange={toggleIngredients}>
                    <SheetContent>
                        <SheetHeader>
                            <SheetTitle>Ingredients</SheetTitle>
                            <SheetDescription>
                                <ul>
                                    {ingredients.map(i => (
                                        <li key={i.name} className="flex space-x-2">
                                            <p>{i.quantity}</p>
                                            <p>{i.unit}</p>
                                            <p>{i.name}</p>
                                        </li>
                                    ))}
                                </ul>
                            </SheetDescription>
                        </SheetHeader>
                    </SheetContent>
                </Sheet>
            </div>
            <BottomSheet
                size={showChat ? "full" : "peek"}
                onSizeChange={(size) => setShowChat(size === "full")}
                peekHeight={12}
                fullHeight={75}
                header={
                    (size) => size === "peek" ? (
                        <h3 className="text-lg font-semibold">Need Help?</h3>
                    ) : (
                        <h2 className="text-lg font-semibold">Here to Help</h2>
                    )
                }
                subHeader={
                    (size) => size === "peek" ? (
                        <p className="text-sm text-muted-foreground">Swipe up and ask away</p>
                    ) : (
                        <p className="text-sm text-muted-foreground">Swipe down to get back to cooking</p>
                    )
                }
            >
                <ChatBody
                    history={chatHistory}
                    loading={askQuestionPending}
                    onSend={askQuestion}
                    error={askQuestionError}
                    onCancel={() => setShowChat(false)}
                />
            </BottomSheet>
        </FocusedLayout>
    )
}

interface CookModeHeaderProps {
    stepIdx: number
    total: number
    title: string
    onBack: () => void
}

function CookModeHeader({ stepIdx, total, title, onBack }: CookModeHeaderProps) {
    const pct = ((stepIdx + 1) / total) * 100

    return (
        <header className="sticky top-0 z-30 bg-background/90 backdrop-blur pb-4">
            <div className="flex items-center justify-between px-2 py-2">
                <Button size="lg" variant="ghost" onClick={onBack} className="!pl-0">
                    <ArrowLeft />
                </Button>
                <h1 className="text-lg font-semibold wrap max-w-[calc(100%-6rem)]">{title}</h1>
                <span className="text-sm tabular-nums">
                    {stepIdx + 1} / {total}
                </span>
            </div>
            <Progress value={pct} className="h-1 rounded-none" />
        </header>
    )
}

interface CookModeInstructionsProps {
    stepIdx: number
    currentStep: string
}

function CookModeInstructions({ stepIdx, currentStep }: CookModeInstructionsProps) {
    return (
        <div className="flex flex-col justify-between p-4">
            <h3 className="text-left text-lg font-semibold pb-2">
                Step {stepIdx + 1}
            </h3>
            <p className="text-left pb-2">{currentStep}</p>
        </div>
    )
}

interface CookModeControlsProps {
    stepIdx: number
    total: number
    onBack: () => void
    onNext: () => void
    toggleIngredients: () => void
}

function CookModeControls({ stepIdx, total, onBack, onNext, toggleIngredients }: CookModeControlsProps) {
    return (
        <div className="flex flex-col justify-between space-y-4 p-4 pb-24">
            <div className="flex justify-center gap-2">
                <Button size="lg" variant="outline" onClick={toggleIngredients} className="flex items-center gap-2 w-full">
                    <List />
                    Ingredients
                </Button>
            </div>
            <div className="flex justify-center gap-2">
                <Button size="lg" disabled={stepIdx === 0} onClick={onBack} className="flex items-center gap-2 w-1/2">
                    <ChevronLeft />
                    Prev
                </Button>
                <Button size="lg" disabled={stepIdx === total - 1} onClick={onNext} className="flex items-center gap-2 w-1/2">
                    Next
                    <ChevronRight />
                </Button>
            </div>
        </div>
    )
}