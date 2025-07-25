import { Button } from "@/components/ui/button";
import { useNavigate, useParams } from "react-router-dom";
import { useRecipe } from "../features/recipe/hooks";
import { NotFoundPage } from "./not-found";
import { ErrorPage } from "./error";
import RecipeSkeleton from "@/features/recipe/recipe-skeleton";
import DefaultLayout from "@/layouts/default-layout";
import { FullRecipe } from "@/features/recipe/recipe-full";
import { Separator } from "@/components/ui/separator";

export default function RecipePage() {
    const recipeId = useParams().id
    return (
        <DefaultLayout>
            <Recipe recipeId={recipeId!}/>
        </DefaultLayout>
    )
}

interface Props {
    recipeId: string
}

export function Recipe({ recipeId }: Props) {
    const nav = useNavigate()
    const { recipe, isLoading, error, notFound } = useRecipe(recipeId)

    if (notFound) {
        return <NotFoundPage />
    }

    if (error) {
        return <ErrorPage title="Error loading recipe" description={error.message} />
    }

    if (isLoading || !recipe) {
        return <RecipeSkeleton />
    }

    return (
        <>
            <div className="border rounded p-2">
                <div className="p-2">
                    <FullRecipe recipe={recipe} />
                </div>
                <Separator/>
                <div className="flex justify-between items-center p-2 pt-4">
                    <Button variant="outline" onClick={() => nav(`/modify-recipe/${recipe.id}`)}>Modify</Button>
                    <Button onClick={() => nav(`/cook/${recipe.id}`)}>Start Cooking</Button>
                </div>
            </div>
        </>
    )
}