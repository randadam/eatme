import { Button } from "@/components/ui/button"
import { APP_NAME } from "@/constants"
import { Link, Navigate } from "react-router-dom"
import { useUser } from "@/features/auth/hooks"

export default function LandingPage() {
    const { isAuthenticated } = useUser()

    if (isAuthenticated) {
        return <Navigate to="/recipes" replace />
    }
    
    return (
        <>
            <title>{APP_NAME}</title>
            <meta name="description" content="AI Recipes Made Just for You" />
            <meta name="keywords" content="AI recipes, meal ideas, cooking, kitchen, cooking assistant" />

            <div className="min-h-screen flex flex-col items-center justify-center text-center px-4 bg-gradient-to-br from-muted/50 to-background">
                <h1 className="text-4xl md:text-5xl font-bold mb-4">
                    AI Recipes Made Just for You
                </h1>
                <p className="text-muted-foreground text-lg max-w-md mb-6">
                    Discover meal ideas you'll actually want to cook — tailored to your taste, time, and pantry.
                </p>

                <Link to="/signup">
                    <Button size="lg">Get Started</Button>
                </Link>

                <div className="mt-4">
                    <p className="mt-4 text-sm text-muted-foreground">
                        Already have an account?{" "}
                        <Link to="/login" className="underline text-primary">
                            Log in
                        </Link>
                    </p>
                </div>

                <div className="mt-12 w-full max-w-sm aspect-video bg-muted rounded-xl shadow-inner" />

                <footer className="mt-12 text-xs text-muted-foreground">
                    © {new Date().getFullYear()} {APP_NAME} — Your AI Kitchen Companion
                </footer>
            </div>
        </>
    )
}
