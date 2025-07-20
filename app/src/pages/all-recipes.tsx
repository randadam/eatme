import { useAllRecipes, useDeleteRecipe } from "@/features/recipe/hooks"
import type api from "@/api"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"
import { ChatBody } from "@/features/chat/chat-body"
import { useState } from "react"
import { useStartSuggestionThread } from "@/features/chat/hooks"
import { RecipeOverview } from "@/features/recipe/recipe-overview"
import { PlusIcon } from "lucide-react"
import RecipeSkeleton from "@/features/recipe/recipe-skeleton"
import { ErrorPage } from "./error"
import { useErrorHandler } from "@/lib/error/error-provider"
import LoaderButton from "@/components/shared/loader-button"
import DefaultLayout from "@/layouts/default-layout"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from "@/components/ui/dialog"

export default function AllRecipesPage() {
    const [dialogOpen, setDialogOpen] = useState(false)

    const nav = useNavigate()
    const showError = useErrorHandler()
    const { recipes, isLoading, error } = useAllRecipes()
    const { deleteRecipe, deleteRecipePending } = useDeleteRecipe()
    const { startThread, startThreadPending } = useStartSuggestionThread()

    function handleAddRecipe() {
        setDialogOpen(true)
    }

    function handleSuggestRecipe(message: string) {
        startThread(message, {
            onSuccess: (threadData) => {
                setDialogOpen(false)
                nav(`/suggest/${threadData.id}`)
            },
            onError: (error) => {
                showError(error)
            }
        })
    }

    return (
        <div>
            <RecipeList
                recipes={recipes}
                isLoading={isLoading}
                error={error}
                deleteRecipe={deleteRecipe}
                deleteRecipePending={deleteRecipePending}
                viewRecipe={(recipeId) => nav(`/recipes/${recipeId}`)}
            />
            <AddButton onClick={handleAddRecipe} />
            <Dialog
                open={dialogOpen}
                onOpenChange={setDialogOpen}
            >
                <DialogContent className="top-[10%] translate-y-0">
                    <DialogHeader>
                        <DialogTitle>Add Recipe</DialogTitle>
                        <DialogDescription>
                            Let me know what you're in the mood for
                        </DialogDescription>
                    </DialogHeader>
                    <ChatBody
                        onSend={handleSuggestRecipe}
                        loading={startThreadPending}
                    />
                </DialogContent>
            </Dialog>
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
            <DefaultLayout>
                <div className="space-y-2">
                    <div className="p-2 border rounded-md">
                        <RecipeSkeleton />
                    </div>
                    <div className="p-2 border rounded-md">
                        <RecipeSkeleton />
                    </div>
                </div>
            </DefaultLayout>
        )
    }
    if (error) {
        return <ErrorPage title="Error loading recipes" description={error.message} />
    }

    if (!recipes || recipes.length === 0) {
        return (
            <DefaultLayout>
                <div className="flex flex-col items-center justify-center space-y-4 pt-16">
                    <h2 className="text-2xl font-semibold">No recipes yet</h2>
                    <p className="text-muted-foreground pt-4">
                        No worries, we're here to help! Press the + button below to get started.
                    </p>
                </div>
            </DefaultLayout>
        )
    }

    return (
        <DefaultLayout>
            <ul className="space-y-2 pb-2">
                {(recipes ?? []).map(recipe => (
                    <li key={recipe.id}>
                        <div className="p-2 border rounded-md">
                            <RecipeOverview recipe={recipe} />
                            <div className="flex justify-between pt-4">
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
        </DefaultLayout>
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