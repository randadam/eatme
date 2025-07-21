import { useRecipe } from "@/features/recipe/hooks";
import CookMode from "@/features/cook/cook-mode";
import { useParams } from "react-router-dom";
import RecipeSkeleton from "@/features/recipe/recipe-skeleton";
import { ErrorPage } from "@/pages/error"
import { useAnswerQuestion } from "@/features/chat/hooks";

export default function CookPage() {
    const recipeId = useParams().recipeId
    const { recipe, isLoading: recipeLoading, error: recipeError } = useRecipe(recipeId!)
    const { answerQuestion, answerQuestionPending, answerQuestionError, history } = useAnswerQuestion(recipe?.thread_id ?? "")

    if (recipeError) {
        return (
            <CookPageLayout>
                <ErrorPage title="Error loading recipe" description={recipeError?.message} />
            </CookPageLayout>
        )
    }
    if (recipeLoading || !recipe) {
        return (
            <CookPageLayout>
                <RecipeSkeleton/>
            </CookPageLayout>
        )
    }

    return (
        <CookPageLayout>
            <CookMode
                id={recipeId!}
                recipe={recipe}
                chatHistory={history}
                askQuestion={answerQuestion}
                askQuestionPending={answerQuestionPending}
                askQuestionError={answerQuestionError?.message}
            />
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