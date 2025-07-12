import { FormDescription, FormField } from "@/components/ui/form"
import { FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"
import type { Control } from "react-hook-form"
import type { DietFormValues } from "./types"

export const dietOptions = [
    { name: "Vegetarian", value: "vegetarian" },
    { name: "Vegan", value: "vegan" },
    { name: "Keto", value: "keto" },
    { name: "Paleo", value: "paleo" },
    { name: "Low Carb", value: "low_carb" },
    { name: "High Protein", value: "high_protein" },
]

export interface DietFormProps {
    control: Control<DietFormValues>
    showTitle?: boolean
}

export default function DietForm({ control, showTitle = true }: DietFormProps) {
    return (
        <FormField
            control={control}
            name="diet"
            render={() => (
                <FormItem>
                    {showTitle && <FormLabel>Diet</FormLabel>}
                    <FormDescription className="text-left">
                        Select any diets you follow.
                    </FormDescription>
                    <div className="pt-4">
                        <MultiSelectBadges
                            name="diet"
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