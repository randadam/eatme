import type { Control } from "react-hook-form"
import { FormField, FormLabel, FormDescription, FormItem, FormMessage } from "@/components/ui/form"
import type { EquipmentFormValues } from "./types"
import { MultiSelectBadges } from "@/components/shared/multi-select-badge"

const equipmentList = [
    { name: "Stove", value: "stove" },
    { name: "Oven", value: "oven" },
    { name: "Microwave", value: "microwave" },
    { name: "Toaster", value: "toaster" },
    { name: "Grill", value: "grill" },
    { name: "Smoker", value: "smoker" },
    { name: "Slow Cooker", value: "slow_cooker" },
    { name: "Pressure Cooker", value: "pressure_cooker" },
    { name: "Sous Vide", value: "sous_vide" },
]

export interface EquipmentFormProps {
    control: Control<EquipmentFormValues>
    showTitle?: boolean
}

export default function EquipmentForm({ control, showTitle = true }: EquipmentFormProps) {
    return (
        <FormField
            control={control}
            name="equipment"
            render={() => (
                <FormItem>
                    {showTitle && <FormLabel>Equipment</FormLabel>}
                    <FormDescription className="text-left">
                        Select any equipment you have.
                    </FormDescription>
                    <div className="pt-4">
                        <MultiSelectBadges
                            name="equipment"
                            control={control}
                            options={equipmentList}
                        />
                    </div>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}
