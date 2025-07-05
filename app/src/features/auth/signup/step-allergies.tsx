import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Checkbox } from "@/components/ui/checkbox";
import { allergiesForm } from "./schemas/forms";
import type { z } from "zod";

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
    const form = useForm<z.infer<typeof allergiesForm>>({
        resolver: zodResolver(allergiesForm),
        defaultValues: {
            allergies: [],
        },
    })

    function onSubmit(values: z.infer<typeof allergiesForm>) {
        console.log(values)
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
                </form>
            </Form>
        </>
    )
}