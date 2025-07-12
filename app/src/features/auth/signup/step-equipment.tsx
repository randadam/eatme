import StepInstructions from "./step-instructions";
import { Form } from "@/components/ui/form";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "../hooks";
import type { ModelsEquipment } from "@/api/client";
import WizardButtons from "./wizard-buttons";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { toast } from "sonner";
import EquipmentForm from "../forms/equipment-form";
import { useEquipmentForm } from "../forms/hooks";
import type { EquipmentFormValues } from "../forms/types";

export default function EquipmentStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useEquipmentForm({
        equipment: profile?.equipment ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: EquipmentFormValues) {
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
                    <EquipmentForm control={form.control}/>
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
        </>
    )
}