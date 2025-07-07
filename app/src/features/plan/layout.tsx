import { useParams } from "react-router-dom"
import { useMealPlan } from "./hooks"
import MealPlan from "./meal-plan"

export default function MealPlanLayout() {
    const mealPlanId = useParams().id
    if (!mealPlanId) {
        return <h1>Missing meal plan ID</h1>
    }
    const { data: mealPlan, isLoading, error } = useMealPlan(mealPlanId)
    console.log('mealPlan', mealPlan)
    return (
        <>
            {isLoading && <p>Loading meal plan...</p>}
            {error && <p>Error loading meal plan: {error.message}</p>}
            {mealPlan && <MealPlan plan={mealPlan}/>}
        </>
    )
}