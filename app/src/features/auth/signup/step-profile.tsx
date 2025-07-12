import StepInstructions from "./step-instructions";
import { Form } from "@/components/ui/form";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "../hooks";
import WizardButtons from "./wizard-buttons";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { BasicProfileForm } from "../forms/basic-profile-form";
import { useBasicProfileForm } from "../forms/hooks";
import type { BasicProfileFormValues } from "../forms/types";

export default function ProfileStep() {
    const nav = useNavigate()
    const user = useUser()
    if (!user.isAuthenticated) {
        return <Navigate to="/" replace />
    }

    const form = useBasicProfileForm({
        name: user.profile?.name ?? "",
    })

    const { saveProfile, isPending } = useSaveProfile()
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: BasicProfileFormValues) {
        saveProfile({ ...values, setup_step: "skill" }, {
            onSuccess: (profile) => {
                toast.success(`Welcome aboard ${profile.name}!`)
                nav("/signup/skill")
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <>
            <StepInstructions>Tell us about yourself</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <BasicProfileForm control={form.control}/>
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
        </>
    )
}