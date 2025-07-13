import { FormField } from "@/components/ui/form"
import { FormLabel, FormDescription, FormItem, FormMessage } from "@/components/ui/form"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"
import type { Control } from "react-hook-form"
import type { AllergiesFormValues } from "./types"

export const allergyOptions = [
    { name: "Peanuts", value: "peanuts" },
    { name: "Tree Nuts", value: "tree_nuts" },
    { name: "Milk / Dairy", value: "milk" },
    { name: "Eggs", value: "eggs" },
    { name: "Wheat / Gluten", value: "wheat" },
    { name: "Soy", value: "soy" },
    { name: "Fish", value: "fish" },
    { name: "Shellfish", value: "shellfish" },
    { name: "Sesame", value: "sesame" },
    { name: "Sulfites", value: "sulfites" },
    { name: "Mustard", value: "mustard" },
    { name: "Celery", value: "celery" },
    { name: "Lupin", value: "lupin" },
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
                    {showTitle && <FormLabel>Allergies</FormLabel>}
                    <FormDescription className="text-left">
                        Select any allergies you have.
                    </FormDescription>
                    <div className="pt-4">
                        <MultiSelectBadges
                            name="allergies"
                            control={control}
                            options={allergyOptions}
                        />
                    </div>
                    <FormMessage/>
                </FormItem>
            )}
        />
    )
}