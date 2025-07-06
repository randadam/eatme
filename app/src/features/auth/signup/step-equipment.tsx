import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import StepInstructions from "./step-instructions";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Checkbox } from "@/components/ui/checkbox";
import { equipmentForm } from "./schemas/forms";
import type { z } from "zod";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "./hooks";
import type { ModelsEquipment } from "@/api/client";
import WizardButtons from "./wizard-buttons";

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
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { mutate: saveProfile, isPending, error } = useSaveProfile()

    const form = useForm<z.infer<typeof equipmentForm>>({
        resolver: zodResolver(equipmentForm),
        defaultValues: {
            equipment: profile?.equipment ?? [],
        },
    })

    function onSubmit(values: z.infer<typeof equipmentForm>) {
        const req = {
            setup_step: "done" as const,
            equipment: values.equipment.map((equipment) => equipment) as ModelsEquipment[],
        }
        saveProfile(req, {
            onSuccess: () => nav("/signup/done"),
        })
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
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
            {error && <p className="text-red-500">{JSON.parse(error.message).detail}</p>}
        </>
    )
}