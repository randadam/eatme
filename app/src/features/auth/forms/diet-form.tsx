import { FormDescription, FormField } from "@/components/ui/form"
import { FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"
import type { Control } from "react-hook-form"
import type { DietFormValues } from "./types"

const diets = [
    { name: "Vegetarian", value: "vegetarian" },
    { name: "Vegan", value: "vegan" },
    { name: "Keto", value: "keto" },
    { name: "Paleo", value: "paleo" },
    { name: "Low Carb", value: "low_carb" },
    { name: "High Protein", value: "high_protein" },
]

export interface DietFormProps {
    control: Control<DietFormValues>
}

export default function DietForm({ control }: DietFormProps) {
    return (
        <FormField
            control={control}
            name="diet"
            render={() => (
                <FormItem>
                    <FormLabel>Diet</FormLabel>
                    <FormDescription className="text-left">
                        Select any diets you follow.
                    </FormDescription>
                    <div className="pt-4">
                        <MultiSelectBadges
                            name="diet"
                            control={control}
                            options={diets}
                        />
                    </div>
                    <FormMessage/>
                </FormItem>
            )}
        />
    )
}