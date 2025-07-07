import { useUser } from "@/features/auth/hooks";
import { Button } from "@/components/ui/button";
import RequireFinishedSignup from "@/features/auth/require-finished-signup";
import { useNavigate } from "react-router-dom";

export default function HomeLayout() {
    const nav = useNavigate()
    const { profile } = useUser();

    return (
        <RequireFinishedSignup>
            <main className="max-w-3xl mx-auto p-6 text-center space-y-6">
                <h1 className="text-4xl font-bold">
                    Welcome back{profile?.name ? `, ${profile.name}` : ""}!
                </h1>
                <p className="text-muted-foreground">
                    Ready to cook something amazing today?
                </p>

                <Button 
                    onClick={() => nav("/recipes")}
                >
                    View Recipes
                </Button>
            </main>
        </RequireFinishedSignup>
    );
}
