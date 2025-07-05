import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Checkbox } from "@/components/ui/checkbox";
import { cuisinesForm } from "./schemas/forms";
import type { z } from "zod";

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
    const form = useForm<z.infer<typeof cuisinesForm>>({
        resolver: zodResolver(cuisinesForm),
        defaultValues: {
            cuisines: [],
        },
    })

    function onSubmit(values: z.infer<typeof cuisinesForm>) {
        console.log(values)
    }

    return (
        <>
            <StepInstructions>What cuisines do you like?</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="cuisines"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Cuisines</FormLabel>
                                <FormDescription className="text-left">
                                    Select your favorite cuisines.
                                </FormDescription>
                                {cuisines.map((cuisine) => (
                                    <FormItem key={cuisine.value} className="flex">
                                        <FormControl>
                                            <Checkbox
                                                checked={field.value.includes(cuisine.value)}
                                                onCheckedChange={(checked) => (
                                                    field.onChange(
                                                        checked
                                                            ? [...field.value, cuisine.value]
                                                            : field.value.filter((value) => value !== cuisine.value)
                                                    )   
                                                )}
                                            />
                                        </FormControl>
                                        <FormLabel>{cuisine.name}</FormLabel>
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