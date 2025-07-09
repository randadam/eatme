import type api from "@/api";
import { useCallback, useMemo, useState } from "react";

export function useCookMode(recipe: api.ModelsRecipeBody) {
    const [stepIdx, setStepIdx] = useState(0)
    const steps = recipe.steps

    const [ingredientsOpen, setIngredientsOpen] = useState(false)

    const goNext = useCallback(() => {
        setStepIdx(i => Math.min(i + 1, steps.length - 1))
    }, [steps.length])

    const goPrev = useCallback(() => {
        setStepIdx(i => Math.max(i - 1, 0))
    }, [steps.length])

    const progress = useMemo(
        () => (stepIdx + 1) / steps.length,
        [stepIdx, steps.length],
    )

    return {
        stepIdx,
        currentStep: steps[stepIdx],
        goNext,
        goPrev,
        progress,
        ingredients: recipe.ingredients,
        ingredientsOpen,
        toggleIngredients: () => setIngredientsOpen(!ingredientsOpen),
    }
}
