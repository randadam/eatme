import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import StepInstructions from "./step-instructions";
import { Form, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { equipmentForm } from "./schemas/forms";
import type { z } from "zod";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "../hooks";
import type { ModelsEquipment } from "@/api/client";
import WizardButtons from "./wizard-buttons";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { toast } from "sonner";
import { MultiSelectBadges } from "@/components/shared/multi-select-badge";

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

    const { saveProfile, isPending } = useSaveProfile()

    const form = useForm<z.infer<typeof equipmentForm>>({
        resolver: zodResolver(equipmentForm),
        defaultValues: {
            equipment: profile?.equipment ?? [],
        },
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: z.infer<typeof equipmentForm>) {
        const req = {
            setup_step: "done" as const,
            equipment: values.equipment.map((equipment) => equipment) as ModelsEquipment[],
        }
        saveProfile(req, {
            onSuccess: () => {
                toast.success("We can work with that!")
                nav("/signup/done")
            },
            onError: (err) => handleFormError(err),
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
                        render={() => (
                            <FormItem>
                                <FormLabel>Equipment</FormLabel>
                                <FormDescription className="text-left">
                                    Select any equipment you have.
                                </FormDescription>
                                <div className="pt-4">
                                    <MultiSelectBadges
                                        name="equipment"
                                        control={form.control}
                                        options={equipmentList}
                                    />
                                </div>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
        </>
    )
}