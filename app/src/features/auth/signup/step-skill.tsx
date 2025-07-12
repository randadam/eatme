import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { zodResolver } from "@hookform/resolvers/zod";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Label } from "@/components/ui/label";
import { skillForm } from "../forms/schemas/forms";
import type { z } from "zod";
import WizardButtons from "./wizard-buttons";
import { useSaveProfile, useUser } from "../hooks";
import { Navigate, useNavigate } from "react-router-dom";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";

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
    value: "chef",
    name: "Chef",
    description: "You're a serious home cook or even work in the industry — nothing in the kitchen intimidates you.",
  },
];


export default function SkillStep() {
    const nav = useNavigate()
    const { isAuthenticated, profile } = useUser()
    if (!isAuthenticated || !profile) {
        return <Navigate to="/" replace />
    }

    const { saveProfile, isPending } = useSaveProfile()

    const form = useForm<z.infer<typeof skillForm>>({
        resolver: zodResolver(skillForm),
        defaultValues: {
            skill: (profile?.skill ?? "beginner") as z.infer<typeof skillForm>['skill'],
        },
    })
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: z.infer<typeof skillForm>) {
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
                    <FormField
                        control={form.control}
                        name="skill"
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
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
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