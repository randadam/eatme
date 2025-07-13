import StepInstructions from "./step-instructions";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "../hooks";
import WizardButtons from "./wizard-buttons";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { useDietForm } from "../forms/hooks";
import type { DietFormValues } from "../forms/types";
import DietForm from "../forms/diet-form";
import { Form } from "@/components/ui/form";

export default function DietStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useDietForm({
        diets: profile?.diets ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: DietFormValues) {
        const req = {
            setup_step: "allergies" as const,
            diets: values.diets,
        }
        saveProfile(req, {
            onSuccess: (profile) => {
                toast.success(profile.diets.length > 0 ? "We can work with that!" : "Endless possibilities await!")
                nav("/signup/allergies")
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <>
            <StepInstructions>Do you have any dietary restrictions?</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <DietForm control={form.control}/>
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
        </>
    )
}