import type api from "@/api";
import { useCookMode } from "./hooks";
import { Button } from "@/components/ui/button";
import { ArrowLeft, ChevronLeft, ChevronRight, CircleQuestionMark, List } from "lucide-react"
import { Progress } from "@/components/ui/progress"
import { useNavigate } from "react-router-dom"
import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle } from "@/components/ui/sheet";

interface CookModeProps {
    id: string
    recipe: api.ModelsRecipeBody
}

export default function CookMode({ id, recipe }: CookModeProps) {
    const nav = useNavigate()
    const { currentStep, stepIdx, goNext, goPrev, ingredients, ingredientsOpen, toggleIngredients } = useCookMode(recipe)

    return (
        <div className="flex flex-col h-screen bg-background">
            <CookModeHeader stepIdx={stepIdx} total={recipe.steps.length} title={recipe.title} onBack={() => nav(`/recipes/${id}`)} />
            <CookModeInstructions currentStep={currentStep} stepIdx={stepIdx}/>
            <CookModeControls stepIdx={stepIdx} total={recipe.steps.length} onBack={goPrev} onNext={goNext} toggleIngredients={toggleIngredients}/>
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
                <Button size="lg" variant="ghost" onClick={onBack}>
                    <ArrowLeft />
                </Button>
                <h1 className="text-lg font-semibold truncate">{title}</h1>
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
        <div className="flex flex-col justify-between">
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
        <div className="fixed bottom-2 inset-x-0 px-4 flex flex-col justify-between space-y-4 pb-4">
            <div className="flex justify-center gap-2">
                <Button size="lg" disabled={stepIdx === 0} onClick={onBack} className="flex items-center gap-2 w-1/2">
                    <ChevronLeft/>
                    Prev
                </Button>
                <Button size="lg" disabled={stepIdx === total - 1} onClick={onNext} className="flex items-center gap-2 w-1/2">
                    Next
                    <ChevronRight/>
                </Button>
            </div>
            <div className="flex justify-center gap-2">
                <Button size="lg" variant="outline" onClick={toggleIngredients} className="flex items-center gap-2 w-1/2">
                    <List/>
                    Ingredients
                </Button>
                <Button size="lg" variant="outline" className="flex items-center gap-2 w-1/2">
                    <CircleQuestionMark/>
                    Help
                </Button>
            </div>
        </div>
    )
}