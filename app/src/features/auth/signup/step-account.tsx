"use client"

import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import StepInstructions from "./step-instructions";
import { accountForm } from "./schemas/forms";
import type { z } from "zod";
import { useSignup } from "../hooks";
import { useNavigate } from "react-router-dom";
import WizardButtons from "./wizard-buttons";

export default function AccountStep() {
    const form = useForm<z.infer<typeof accountForm>>({
        resolver: zodResolver(accountForm),
        defaultValues: {
            email: "",
            password: "",
            confirmPassword: "",
        },
    });

    const nav = useNavigate()
    const { signup, isPending, error } = useSignup()

    function onSubmit(values: z.infer<typeof accountForm>) {
        signup(values, {
            onSuccess: () => nav("/signup/profile"),
        })
    }

    return (
        <>
            <StepInstructions>Create your account</StepInstructions>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                    <FormField
                        control={form.control}
                        name="email"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Email</FormLabel>
                                <FormControl>
                                    <Input placeholder="Email" {...field}/>
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input type="password" placeholder="Password" {...field}/>
                                </FormControl>
                                <FormMessage/>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="confirmPassword"
                        render={({field}) => (
                            <FormItem>
                                <FormLabel>Confirm Password</FormLabel>
                                <FormControl>
                                    <Input type="password" placeholder="Confirm Password" {...field}/>
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
    );
}