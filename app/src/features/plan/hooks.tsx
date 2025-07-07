import { useMutation, useQuery } from "@tanstack/react-query"
import api from "@/api"

const mealPlanKeys = {
    all: ["meal-plan"] as const,
    byId: (mealPlanId: string) => [...mealPlanKeys.all, mealPlanId] as const,
}

export function useNewMealPlan() {
    return useMutation({
        mutationFn: async () => {
            const res = await api.createMealPlan()
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
            return res.data as api.ModelsMealPlan
        },
    })
}

export function useMealPlan(mealPlanId: string) {
    return useQuery({
        queryKey: mealPlanKeys.byId(mealPlanId),
        queryFn: async () => {
            const res = await api.getMealPlan(mealPlanId)
            if (res.status > 299) {
                throw new Error(JSON.stringify(res.data))
            }
            return res.data as api.ModelsMealPlan
        },
    })
}
