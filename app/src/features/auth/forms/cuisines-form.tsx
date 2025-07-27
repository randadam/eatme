import type { CuisinesFormValues } from "./types"
import { FormField } from "@/components/ui/form"
import { FormLabel, FormDescription, FormItem, FormMessage } from "@/components/ui/form"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"
import type { Control } from "react-hook-form"

export const cuisineOptions = [
    { name: "American", value: "american" },
    { name: "BBQ", value: "bbq" },
    { name: "British", value: "british" },
    { name: "Cajun", value: "cajun" },
    { name: "Mexican", value: "mexican" },
    { name: "Tex-Mex", value: "tex_mex" },
    { name: "Caribbean", value: "caribbean" },
    { name: "Latin American", value: "latin-american" },
    { name: "Italian", value: "italian" },
    { name: "French", value: "french" },
    { name: "Spanish", value: "spanish" },
    { name: "Greek", value: "greek" },
    { name: "Mediterranean", value: "mediterranean" },
    { name: "Middle Eastern", value: "middle-eastern" },
    { name: "Indian", value: "indian"},
    { name: "Thai", value: "thai"},
    { name: "Vietnamese", value: "vietnamese"},
    { name: "Chinese", value: "chinese"},
    { name: "Japanese", value: "japanese"},
    { name: "Korean", value: "korean"},
    { name: "Filipino", value: "filipino"},
    { name: "African", value: "african"},
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