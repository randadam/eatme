import StepInstructions from "./step-instructions";
import { Form } from "@/components/ui/form";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "../hooks";
import WizardButtons from "./wizard-buttons";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { toast } from "sonner";
import { useAllergiesForm } from "../forms/hooks";
import type { AllergiesFormValues } from "../forms/types";
import AllergiesForm from "../forms/allergies-form";

export default function AllergiesStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useAllergiesForm({
        allergies: profile?.allergies ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: AllergiesFormValues) {
        const req = {
            setup_step: "equipment" as const,
            allergies: values.allergies,
        }
        saveProfile(req, {
            onSuccess: () => {
                // TODO: Confirmation that user is still responsible for allergies
                toast.success("We'll keep that in mind!")
                nav("/signup/equipment")
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <>
            <StepInstructions>Do you have any allergies?</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <AllergiesForm control={form.control}/>
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
        </>
    )
}