import StepInstructions from "./step-instructions";
import WizardButtons from "./wizard-buttons";
import { useSaveProfile, useUser } from "../hooks";
import { Navigate, useNavigate } from "react-router-dom";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import SkillForm from "../forms/skill-form";
import { useSkillForm } from "../forms/hooks";
import { Form } from "@/components/ui/form";
import type { SkillFormValues } from "../forms/types";

export default function SkillStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useSkillForm({
        skill: profile.skill,
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: SkillFormValues) {
        saveProfile({ ...values, setup_step: "cuisines" }, {
            onSuccess: (profile) => {
                let skillLevelMessage = "Skill level set successfully"
                switch (profile.skill) {
                    case "beginner":
                        skillLevelMessage = "We're going to turn you into a master chef!"
                        break
                    case "intermediate":
                        skillLevelMessage = "You're on your way to becoming a master chef!"
                        break
                    case "advanced":
                        skillLevelMessage = "You're well on your way to becoming a master chef!"
                        break
                    case "chef":
                        skillLevelMessage = "Can't wait to help you hone your skills!"
                        break
                }
                toast.success(skillLevelMessage)
                nav("/signup/cuisines")
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <>
            <StepInstructions>What is your skill level?</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <SkillForm control={form.control}/>
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
        </>
    )
}
