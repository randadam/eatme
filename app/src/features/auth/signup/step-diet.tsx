import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { dietForm } from "./schemas/forms";
import type { z } from "zod";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "../hooks";
import WizardButtons from "./wizard-buttons";
import type { ModelsDiet } from "@/api/client";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { MultiSelectBadges } from "@/components/shared/multi-select-badge";

const diets = [
    { name: "Vegetarian", value: "vegetarian" },
    { name: "Vegan", value: "vegan" },
    { name: "Keto", value: "keto" },
    { name: "Paleo", value: "paleo" },
    { name: "Low Carb", value: "low_carb" },
    { name: "High Protein", value: "high_protein" },
]

export default function DietStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useForm<z.infer<typeof dietForm>>({
        resolver: zodResolver(dietForm),
        defaultValues: {
            diet: profile?.diet ?? [],
        },
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: z.infer<typeof dietForm>) {
        const req = {
            setup_step: "equipment" as const,
            diet: values.diet.map((diet) => diet) as ModelsDiet[],
        }
        saveProfile(req, {
            onSuccess: (profile) => {
                toast.success(profile.diet.length > 0 ? "We can work with that!" : "Endless possibilities await!")
                nav("/signup/equipment")
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <>
            <StepInstructions>Do you have any dietary restrictions?</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="diet"
                        render={() => (
                            <FormItem>
                                <FormLabel>Diet</FormLabel>
                                <FormDescription className="text-left">
                                    Select any diets you follow.
                                </FormDescription>
                                <div className="pt-4">
                                    <MultiSelectBadges
                                        name="diet"
                                        control={form.control}
                                        options={diets}
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