"use client"

import StepInstructions from "./step-instructions";
import { useForm } from "react-hook-form";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { profileForm } from "./schemas/forms";
import type { z } from "zod";
import { Navigate, useNavigate } from "react-router-dom";
import { useSaveProfile, useUser } from "./hooks";
import WizardButtons from "./wizard-buttons";

export default function ProfileStep() {
    const nav = useNavigate()
    const user = useUser()
    if (!user.isAuthenticated) {
        return <Navigate to="/" replace />
    }

    const form = useForm<z.infer<typeof profileForm>>({
        resolver: zodResolver(profileForm),
        defaultValues: {
            name: user.profile?.name ?? "",
        },
    })

    const { mutate: saveProfile, isPending, error } = useSaveProfile()

    function onSubmit(values: z.infer<typeof profileForm>) {
        saveProfile({ ...values, setup_step: "skill" }, {
            onSuccess: () => nav("/signup/skill"),
        })
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
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
            {error && <p className="text-red-500">{JSON.parse(error.message).detail}</p>}
        </>
    )
}