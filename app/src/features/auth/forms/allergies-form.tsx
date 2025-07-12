import { FormField } from "@/components/ui/form"
import { FormLabel, FormDescription, FormItem, FormMessage } from "@/components/ui/form"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"
import type { Control } from "react-hook-form"
import type { AllergiesFormValues } from "./types"

const allergies = [
    { name: "Dairy", value: "dairy" },
    { name: "Eggs", value: "eggs" },
    { name: "Fish", value: "fish" },
    { name: "Gluten", value: "gluten" },
    { name: "Peanuts", value: "peanuts" },
    { name: "Soy", value: "soy" },
    { name: "Tree Nuts", value: "tree_nuts" },
    { name: "Wheat", value: "wheat" },
]

export interface AllergiesFormProps {
    control: Control<AllergiesFormValues>
    showTitle?: boolean
}

export default function AllergiesForm({ control, showTitle = true }: AllergiesFormProps) {
    return (
        <FormField
            control={control}
            name="allergies"
            render={() => (
                <FormItem>
                    {showTitle && <FormLabel>Do you have any allergies?</FormLabel>}
                    <FormDescription className="text-left">
                        Select any allergies you have.
                    </FormDescription>
                    <div className="pt-4">
                        <MultiSelectBadges
                            name="allergies"
                            control={control}
                            options={allergies}
                        />
                    </div>
                    <FormMessage/>
                </FormItem>
            )}
        />
    )
}