import type { CuisinesFormValues } from "./types"
import { FormField } from "@/components/ui/form"
import { FormLabel, FormDescription, FormItem, FormMessage } from "@/components/ui/form"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"
import type { Control } from "react-hook-form"

export const cuisineOptions = [
    { name: "American", value: "american" },
    { name: "British", value: "british" },
    { name: "Chinese", value: "chinese" },
    { name: "French", value: "french" },
    { name: "German", value: "german" },
    { name: "Indian", value: "indian" },
    { name: "Italian", value: "italian" },
    { name: "Japanese", value: "japanese" },
    { name: "Mexican", value: "mexican" },
    { name: "Spanish", value: "spanish" },
    { name: "Thai", value: "thai" },
    { name: "Vietnamese", value: "vietnamese" },
]


export interface CuisinesFormProps {
    control: Control<CuisinesFormValues>
    showTitle?: boolean
}

export default function CuisinesForm({ control, showTitle = true }: CuisinesFormProps) {
    return (
        <FormField
            control={control}
            name="cuisines"
            render={() => (
                <FormItem>
                    {showTitle && <FormLabel>Cuisines</FormLabel>}
                    <FormDescription className="text-left">
                        Select your favorite cuisines.
                    </FormDescription>
                    <div className="pt-4">
                        <MultiSelectBadges
                            name="cuisines"
                            control={control}
                            options={cuisineOptions}
                        />
                    </div>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}