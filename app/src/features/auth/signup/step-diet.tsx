import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Checkbox } from "@/components/ui/checkbox";
import { dietForm } from "./schemas/forms";
import type { z } from "zod";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "./hooks";
import WizardButtons from "./wizard-buttons";
import type { ModelsDiet } from "@/api/client";

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

    const { mutate: saveProfile, isPending, error } = useSaveProfile()

    const form = useForm<z.infer<typeof dietForm>>({
        resolver: zodResolver(dietForm),
        defaultValues: {
            diet: profile?.diet ?? [],
        },
    })

    function onSubmit(values: z.infer<typeof dietForm>) {
        const req = {
            setup_step: "equipment" as const,
            diet: values.diet.map((diet) => diet) as ModelsDiet[],
        }
        saveProfile(req, {
            onSuccess: () => nav("/signup/equipment"),
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
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Diet</FormLabel>
                                <FormDescription className="text-left">
                                    Select any diets you follow.
                                </FormDescription>
                                {diets.map((diet) => (
                                    <FormItem key={diet.value} className="flex">
                                        <FormControl>
                                            <Checkbox
                                                checked={field.value.includes(diet.value)}
                                                onCheckedChange={(checked) => (
                                                    field.onChange(
                                                        checked
                                                            ? [...field.value, diet.value]
                                                            : field.value.filter((value) => value !== diet.value)
                                                    )   
                                                )}
                                            />
                                        </FormControl>
                                        <FormLabel>{diet.name}</FormLabel>
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