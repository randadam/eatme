import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Checkbox } from "@/components/ui/checkbox";
import { allergiesForm } from "./schemas/forms";
import type { z } from "zod";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "../hooks";
import type { ModelsAllergy } from "@/api/client";
import WizardButtons from "./wizard-buttons";

const allergies = [
    { name: "Dairy", value: "dairy" },
    { name: "Eggs", value: "eggs" },
    { name: "Fish", value: "fish" },
    { name: "Gluten", value: "gluten" },
    { name: "Peanuts", value: "peanuts" },
    { name: "Soy", value: "soy" },
    { name: "Tree Nuts", value: "tree_nuts" },
    { name: "Wheat", value: "wheat" },
]

export default function AllergiesStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { mutate: saveProfile, isPending, error } = useSaveProfile()

    const form = useForm<z.infer<typeof allergiesForm>>({
        resolver: zodResolver(allergiesForm),
        defaultValues: {
            allergies: profile?.allergies ?? [],
        },
    })

    function onSubmit(values: z.infer<typeof allergiesForm>) {
        const req = {
            setup_step: "equipment" as const,
            allergies: values.allergies.map((allergy) => allergy) as ModelsAllergy[],
        }
        saveProfile(req, {
            onSuccess: () => nav("/signup/equipment"),
        })
    }

    return (
        <>
            <StepInstructions>Allergies</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="allergies"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Do you have any allergies?</FormLabel>
                                <FormDescription className="text-left">
                                    Select any allergies you have.
                                </FormDescription>
                                {allergies.map((allergy) => (
                                    <FormItem key={allergy.value} className="flex">
                                        <FormControl>
                                            <Checkbox
                                                checked={field.value.includes(allergy.value)}
                                                onCheckedChange={(checked) => (
                                                    field.onChange(
                                                        checked
                                                            ? [...field.value, allergy.value]
                                                            : field.value.filter((value) => value !== allergy.value)
                                                        )
                                                )}
                                            />
                                        </FormControl>
                                        <FormLabel>{allergy.name}</FormLabel>
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