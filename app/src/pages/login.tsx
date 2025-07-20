import { Form, FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form"
import { APP_NAME } from "@/constants"
import { Input } from "@/components/ui/input"
import { useForm } from "react-hook-form"
import LoaderButton from "@/components/shared/loader-button"
import { Link, useNavigate } from "react-router-dom"
import { useLogin, useUser } from "@/features/auth/hooks"
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider"

interface LoginValues {
    email: string
    password: string
}

export default function LoginPage() {
    const form = useForm<LoginValues>({
        defaultValues: {
            email: "",
            password: "",
        }
    })
    const handleFormError = useFormErrorHandler(form)
    const { isAuthenticated } = useUser()
    const { login, isPending }= useLogin()
    const nav = useNavigate()

    const handleLogin = (values: LoginValues) => {
        login(values, {
            onError: (err) => {
                handleFormError(err)
            }
        })
    }

    if (isAuthenticated) {
        nav("/")
    }

    return (
        <>
            <title>{APP_NAME} - Login</title>
            <meta name="description" content="Login to your account" />

            <div className="flex flex-col pt-16">
                <div className="space-y-6 p-6 border rounded">
                    <h2 className="text-xl font-semibold">Login</h2>
                    <Form {...form}>
                        <form onSubmit={form.handleSubmit(handleLogin)} className="space-y-4">
                            <FormField
                                control={form.control}
                                name="email"
                                render={({ field }) => (
                                    <FormItem
                                        className="space-y-2"
                                        {...field}
                                    >
                                        <FormLabel>Email</FormLabel>
                                        <FormControl>
                                            <Input type="email" placeholder="Email" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="password"
                                render={({ field }) => (
                                    <FormItem
                                        className="space-y-2"
                                        {...field}
                                    >
                                        <FormLabel>Password</FormLabel>
                                        <FormControl>
                                            <Input type="password" placeholder="Password" {...field} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <LoaderButton type="submit" className="w-full" isLoading={isPending}>
                                Login
                            </LoaderButton>
                            <FormErrorMessage form={form} />
                        </form>
                    </Form>
                </div>

                <div className="h-16 flex items-center justify-center">
                    <p className="mt-4 text-sm text-muted-foreground">
                        Don't have an account?{" "}
                        <Link to="/signup" className="underline text-primary">
                            Sign up
                        </Link>
                    </p>
                </div>

            </div>
        </>
    )
}