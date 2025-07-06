import { useUser } from "@/features/auth/hooks";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
import RequireFinishedSignup from "../auth/require-finished-signup";

export default function HomeLayout() {
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

                <div className="flex justify-center gap-4">
                    <Button asChild>
                        <Link to="/recipes">Explore Recipes</Link>
                    </Button>
                    <Button variant="secondary" asChild>
                        <Link to="/profile">Edit Preferences</Link>
                    </Button>
                </div>
            </main>
        </RequireFinishedSignup>
    );
}
