import { FormDescription, FormField } from "@/components/ui/form"
import { FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"
import type { Control } from "react-hook-form"
import type { DietFormValues } from "./types"

export const dietOptions = [
    { name: "Vegetarian", value: "vegetarian" },
    { name: "Vegan", value: "vegan" },
    { name: "Pescatarian", value: "pescatarian" },
    { name: "Keto", value: "keto" },
    { name: "Low-Carb", value: "low_carb" },
    { name: "High-Protein", value: "high_protein" },
    { name: "Paleo", value: "paleo" },
    { name: "Whole 30", value: "whole_30" },
    { name: "Mediterranean Diet", value: "mediterranean_diet" },
    { name: "DASH", value: "dash" },
    { name: "Low-FODMAP", value: "low_fodmap" },
    { name: "Gluten-Free", value: "gluten_free" },
    { name: "Dairy-Free", value: "dairy_free" },
    { name: "Low Sodium", value: "low_sodium" },
    { name: "Heart-Healthy", value: "heart_healthy" },
    { name: "Diabetic-Friendly", value: "diabetic_friendly" },
]

export interface DietFormProps {
    control: Control<DietFormValues>
    showTitle?: boolean
}

export default function DietForm({ control, showTitle = true }: DietFormProps) {
    return (
        <FormField
            control={control}
            name="diets"
            render={() => (
                <FormItem>
                    {showTitle && <FormLabel>Diet</FormLabel>}
                    <FormDescription className="text-left">
                        Select any diets you follow.
                    </FormDescription>
                    <div className="pt-4">
                        <MultiSelectBadges
                            name="diets"
                            control={control}
                            options={dietOptions}
                        />
                    </div>
                    <FormMessage/>
                </FormItem>
            )}
        />
    )
}