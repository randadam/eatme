import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import StepInstructions from "./step-instructions";
import { accountForm } from "../forms/schemas/forms";
import type { z } from "zod";
import { useSignup, useUser } from "../hooks";
import { Link, Navigate, useNavigate } from "react-router-dom";
import WizardButtons from "./wizard-buttons";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { toast } from "sonner";

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
    const user = useUser()
    const { signup, isPending } = useSignup()
    const handleFormError = useFormErrorHandler(form)

    function onSubmit(values: z.infer<typeof accountForm>) {
        signup(values, {
            onSuccess: () => {
                toast.success("Account created successfully")
                nav("/signup/profile")
            },
            onError: (err) => {
                handleFormError(err)
            },
        })
    }

    if (user.isAuthenticated) {
        return <Navigate to="/" replace />
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
                    <FormErrorMessage form={form}/>
                    <WizardButtons loading={isPending}/>
                </form>
            </Form>
            <div className="mt-4">
                <p className="mt-4 text-sm text-muted-foreground">
                    Already have an account?{" "}
                    <Link to="/login" className="underline text-primary">
                        Log in
                    </Link>
                </p>
            </div>
        </>
    );
}