import type api from "@/api";
import { useCallback, useEffect, useMemo, useState } from "react";
import { useSearchParams } from "react-router-dom";

export function useCookMode(recipe: api.ModelsRecipeBody) {
    const [searchParams, setSearchParams] = useSearchParams({ step: "1" })
    const [stepIdx, setStepIdx] = useState(Number(searchParams.get("step")) - 1)
    const steps = recipe.steps

    const [ingredientsOpen, setIngredientsOpen] = useState(false)

    const goNext = useCallback(() => {
        setStepIdx(i => Math.min(i + 1, steps.length - 1))
    }, [steps.length])

    const goPrev = useCallback(() => {
        setStepIdx(i => Math.max(i - 1, 0))
    }, [steps.length])

    useEffect(() => {
        setSearchParams({ step: String(stepIdx + 1) })
    }, [stepIdx, setSearchParams])

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
