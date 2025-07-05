"use client"

import z from "zod";
import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Label } from "@/components/ui/label";

const skills: SkillLevelProps[] = [
  {
    value: "beginner",
    name: "Beginner",
    description: "You can boil water and heat up canned soup — and that's totally fine.",
  },
  {
    value: "intermediate",
    name: "Intermediate",
    description: "You're comfortable following recipes and can whip up simple meals, often using canned or packaged ingredients.",
  },
  {
    value: "advanced",
    name: "Advanced",
    description: "You cook from scratch often and feel confident experimenting with techniques and flavors.",
  },
  {
    value: "master",
    name: "Master",
    description: "You're a serious home cook or even work in food — nothing in the kitchen intimidates you.",
  },
];


const formSchema = z.object({
    name: z.string().nonempty("Please enter a name."),
    skillLevel: z.enum(["beginner", "intermediate", "advanced", "master"]),
})

export default function ProfileStep() {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: "",
            skillLevel: "beginner",
        },
    })

    function onSubmit(values: z.infer<typeof formSchema>) {
        console.log(values)
    }

    return (
        <>
            <StepInstructions>Tell us about yourself</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="name"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Name</FormLabel>
                                <FormDescription className="text-left">
                                    What should we call you?
                                    This can be anything you want, we're not the IRS.
                                </FormDescription>
                                <FormControl>
                                    <Input placeholder="Name" autoComplete="nickname" {...field}/>
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="skillLevel"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Current Skill Level</FormLabel>
                                <FormDescription className="text-left">
                                    How experienced are you in the kitchen? Don't worry, no matter your current skill level,
                                    we'll help you level up.
                                </FormDescription>
                                <FormControl>
                                    <RadioGroup
                                        value={field.value}
                                        onValueChange={field.onChange}
                                    >
                                        {skills.map(skill => (
                                            <FormItem key={skill.value}>
                                                <FormControl>
                                                    <SkillLevel {...skill}/>
                                                </FormControl>
                                            </FormItem>
                                        ))}
                                    </RadioGroup>
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />
                </form>
            </Form>
        </>
    )
}

interface SkillLevelProps {
    value: string
    name: string
    description: string
}

function SkillLevel({ value, name, description }: SkillLevelProps) {
    return (
        <div className="flex space-x-2 border rounded p-2">
            <RadioGroupItem value={value} id={value}/>
            <Label htmlFor={value}>
                <div className="flex flex-col items-start">
                    <span className="pb-2">{name}:</span>
                    <span className="text-left text-xs text-muted-foreground">
                        {description}
                    </span>
                </div>
            </Label>
        </div>
    )
}