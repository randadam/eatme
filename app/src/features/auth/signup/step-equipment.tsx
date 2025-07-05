import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import StepInstructions from "./step-instructions";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Checkbox } from "@/components/ui/checkbox";
import { equipmentForm } from "./schemas/forms";
import type { z } from "zod";

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

export default function EquipmentStep() {
    const form = useForm<z.infer<typeof equipmentForm>>({
        resolver: zodResolver(equipmentForm),
        defaultValues: {
            equipment: [],
        },
    })

    function onSubmit(values: z.infer<typeof equipmentForm>) {
        console.log(values)
    }

    return (
        <>
            <StepInstructions>What equipment do you have?</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="equipment"
                        render={({field}) => (
                            <FormItem>
                                {equipmentList.map((equipment) => (
                                    <FormItem key={equipment.value} className="flex">
                                        <FormControl>
                                            <Checkbox
                                                checked={field.value.includes(equipment.value)}
                                                onCheckedChange={(checked) => (
                                                    field.onChange(
                                                        checked
                                                            ? [...field.value, equipment.value]
                                                            : field.value.filter((value) => value !== equipment.value)
                                                    )   
                                                )}
                                            />
                                        </FormControl>
                                        <FormLabel>{equipment.name}</FormLabel>
                                    </FormItem>
                                ))}
                                <FormMessage/>
                            </FormItem>
                        )}
                    />
                </form>
            </Form>
        </>
    )
}