import { useRecipe } from "@/features/recipe/hooks";
import CookMode from "@/features/cook/cook-mode";
import { useParams } from "react-router-dom";
import RecipeSkeleton from "@/features/recipe/recipe-skeleton";
import { ErrorPage } from "@/pages/error"

export default function CookPage() {
    const recipeId = useParams().recipeId
    const { recipe, isLoading, error } = useRecipe(recipeId!)

    if (error) {
        return (
            <CookPageLayout>
                <ErrorPage title="Error loading recipe" description={error.message} />
            </CookPageLayout>
        )
    }
    if (isLoading || !recipe) {
        return (
            <CookPageLayout>
                <RecipeSkeleton/>
            </CookPageLayout>
        )
    }

    return (
        <CookPageLayout>
            <CookMode id={recipeId!} recipe={recipe}/>
        </CookPageLayout>
    )
}

interface CookPageLayoutProps {
    children: React.ReactNode
}

function CookPageLayout({ children }: CookPageLayoutProps) {
    return (
        <div className="flex flex-col h-screen bg-background">
            {children}
        </div>
    )
}