import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { cuisinesForm } from "../forms/schemas/forms";
import type { z } from "zod";
import WizardButtons from "./wizard-buttons";
import { useSaveProfile, useUser } from "../hooks";
import { Navigate, useNavigate } from "react-router-dom";
import type { ModelsCuisine } from "@/api/client";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { toast } from "sonner";
import { MultiSelectBadges } from "@/components/shared/multi-select-badge";

const cuisines = [
    { name: "American", value: "american" },
    { name: "British", value: "british" },
    { name: "Chinese", value: "chinese" },
    { name: "French", value: "french" },
    { name: "German", value: "german" },
    { name: "Indian", value: "indian" },
    { name: "Italian", value: "italian" },
    { name: "Japanese", value: "japanese" },
    { name: "Mexican", value: "mexican" },
    { name: "Spanish", value: "spanish" },
    { name: "Thai", value: "thai" },
    { name: "Vietnamese", value: "vietnamese" },
]

export default function CuisinesStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useForm<z.infer<typeof cuisinesForm>>({
        resolver: zodResolver(cuisinesForm),
        defaultValues: {
            cuisines: profile?.cuisines ?? [],
        },
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: z.infer<typeof cuisinesForm>) {
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
                    <FormField
                        control={form.control}
                        name="cuisines"
                        render={() => (
                            <FormItem>
                                <FormLabel>Cuisines</FormLabel>
                                <FormDescription className="text-left">
                                    Select your favorite cuisines.
                                </FormDescription>
                                <div className="pt-4">
                                    <MultiSelectBadges
                                        name="cuisines"
                                        control={form.control}
                                        options={cuisines}
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