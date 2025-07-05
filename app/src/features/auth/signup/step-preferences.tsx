import z from "zod";
import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Checkbox } from "@/components/ui/checkbox";

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

const formSchema = z.object({
  cuisines: z.array(z.string()).min(1, "Please select at least one cuisine."),
  allergies: z.array(z.string()),
})

export default function PreferencesStep() {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            cuisines: [],
            allergies: [],
        },
    })

    function onSubmit(values: z.infer<typeof formSchema>) {
        console.log(values)
    }

    return (
        <>
            <StepInstructions>Preferences</StepInstructions>
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
                    <FormField
                        control={form.control}
                        name="allergies"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Allergies</FormLabel>
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