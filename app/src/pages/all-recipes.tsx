import { useAllRecipes, useDeleteRecipe } from "@/features/recipe/hooks"
import type api from "@/api"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"
import { ChatDrawer } from "@/features/chat/chat-drawer"
import { useState } from "react"
import { useStartSuggestionThread } from "@/features/chat/hooks"
import { Separator } from "@/components/ui/separator"
import { RecipeOverview } from "@/features/recipe/recipe-overview"
import { PlusIcon } from "lucide-react"
import RecipeSkeleton from "@/features/recipe/recipe-skeleton"
import { ErrorPage } from "./error"
import { useErrorHandler } from "@/lib/error/error-provider"
import LoaderButton from "@/components/shared/loader-button"

export default function AllRecipesPage() { 
    const [drawerOpen, setDrawerOpen] = useState(false)

    const nav = useNavigate()
    const showError = useErrorHandler()
    const { recipes, isLoading, error } = useAllRecipes()
    const { deleteRecipe, deleteRecipePending } = useDeleteRecipe()
    const { startThread, startThreadPending } = useStartSuggestionThread()

    function handleAddRecipe() {
        setDrawerOpen(true)
    }

    function handleSuggestRecipe(message: string) {
        startThread(message, {
            onSuccess: (threadData) => {
                setDrawerOpen(false)
                nav(`/suggest/${threadData.id}`)
            },
            onError: (error) => {
                showError(error)
            }
        })
    }

    return (
        <div>
            <h1 className="text-2xl font-bold pb-6">Recipe Book</h1>
            <RecipeList 
                recipes={recipes} 
                isLoading={isLoading} 
                error={error} 
                deleteRecipe={deleteRecipe} 
                deleteRecipePending={deleteRecipePending}
                viewRecipe={(recipeId) => nav(`/recipes/${recipeId}`)}
            />
            <AddButton onClick={handleAddRecipe}/>
            <ChatDrawer
                open={drawerOpen}
                onOpenChange={(open) => setDrawerOpen(open)}
                mode="suggest"
                onSend={handleSuggestRecipe}
                loading={startThreadPending}
            />
        </div>
    )
}

interface RecipeListProps {
    recipes?: api.ModelsUserRecipe[]
    isLoading: boolean
    error: Error | null
    deleteRecipe: (recipeId: string) => void
    deleteRecipePending: boolean
    viewRecipe: (recipeId: string) => void
}

function RecipeList({ recipes, isLoading, error, deleteRecipe, deleteRecipePending, viewRecipe }: RecipeListProps) {
    if (isLoading) {
        return (
            <div className="space-y-2">
                <RecipeSkeleton />
                <Separator />
                <RecipeSkeleton />
            </div>
        )
    }
    if (error) {
        return <ErrorPage title="Error loading recipes" description={error.message} />
    }

    return (
        <ul className="space-y-2 pb-2">
            {(recipes ?? []).map(recipe => (
                <li key={recipe.id}>
                    <Separator />
                    <div className="p-2">
                        <RecipeOverview recipe={recipe} />
                        <div className="flex justify-between">
                            <LoaderButton
                                variant="destructive"
                                className="mt-2"
                                onClick={() => deleteRecipe(recipe.id)}
                                isLoading={deleteRecipePending}
                            >
                                Delete Recipe
                            </LoaderButton>
                            <Button className="mt-2" onClick={() => viewRecipe(recipe.id)}>
                                View Recipe
                            </Button>
                        </div>
                    </div>
                </li>
            ))}
        </ul>
    )
}

interface AddButtonProps {
    onClick: () => void
}

function AddButton({ onClick }: AddButtonProps) {
    return (
        <Button className="fixed bottom-18 right-4 flex items-center justify-center z-50 h-12 w-12 rounded-full" onClick={onClick}>
            <PlusIcon className="h-6 w-6" />
        </Button>
    )
}