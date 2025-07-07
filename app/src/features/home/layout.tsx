import { useUser } from "@/features/auth/hooks";
import { Button } from "@/components/ui/button";
import RequireFinishedSignup from "../auth/require-finished-signup";
import { useNewMealPlan } from "../plan/hooks";
import { Loader2 } from "lucide-react";
import { useNavigate } from "react-router-dom";

export default function HomeLayout() {
    const nav = useNavigate()
    const { profile } = useUser();
    const { mutate: newMealPlan, isPending, error } = useNewMealPlan();

    function handleCreateMealPlan() {
        newMealPlan(void 0, {
            onSuccess: (resp) => nav(`/plan/${resp.id}`),
        })
    }

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
                    onClick={handleCreateMealPlan}
                    disabled={isPending}
                >
                    {isPending && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                    Create Meal Plan
                </Button>
                {error && <p>Error creating meal plan: {error.message}</p>}
            </main>
        </RequireFinishedSignup>
    );
}
