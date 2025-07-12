import StepInstructions from "./step-instructions";
import { Form } from "@/components/ui/form";
import WizardButtons from "./wizard-buttons";
import { useSaveProfile, useUser } from "../hooks";
import { Navigate, useNavigate } from "react-router-dom";
import type { ModelsCuisine } from "@/api/client";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { toast } from "sonner";
import CuisinesForm from "../forms/cuisines-form";
import { useCuisinesForm } from "../forms/hooks";
import type { CuisinesFormValues } from "../forms/types";

export default function CuisinesStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useCuisinesForm({
        cuisines: profile?.cuisines ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: CuisinesFormValues) {
        const req = {
            setup_step: "diet" as const,
            cuisines: values.cuisines.map((cuisine) => cuisine) as ModelsCuisine[],
        }
        saveProfile(req, {
            onSuccess: () => {
                toast.success("Getting hungry just thinking about the possibilities!")
                nav("/signup/diet")
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <>
            <StepInstructions>What cuisines do you like?</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <CuisinesForm control={form.control}/>
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
        </>
    )
}