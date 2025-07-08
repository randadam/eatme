import { useAllRecipes } from "@/features/recipe/hooks"
import type api from "@/api"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"
import { ChatDrawer } from "@/features/chat/chat-drawer"
import { useState } from "react"
import { useStartSuggestionThread } from "@/features/chat/hooks"
import { Separator } from "@/components/ui/separator"
import { RecipeOverview } from "@/features/recipe/recipe-overview"

export default function AllRecipesPage() { 
    const [drawerOpen, setDrawerOpen] = useState(false)

    const nav = useNavigate()
    const { recipes, isLoading, error } = useAllRecipes()
    const { startThread, startThreadPending, startThreadError } = useStartSuggestionThread()

    if (isLoading) {
        return <p>Loading...</p>
    }
    if (error) {
        return <p>Error: {error.message}</p>
    }

    function handleAddRecipe() {
        setDrawerOpen(true)
    }

    function handleSuggestRecipe(message: string) {
        startThread(message, {
            onSuccess: (resp) => {
                const threadData = resp.data as api.ModelsSuggestChatResponse
                console.log(threadData)
                setDrawerOpen(false)
                nav(`/suggest/${threadData.thread_id}`)
            },
        })
    }

    const recipesList = (recipes?.data ?? []) as api.ModelsUserRecipe[]

    return (
        <div>
            <h1 className="text-2xl font-bold pb-6">Recipe Book</h1>
            <ul className="space-y-2 pb-2">
                {recipesList.map(recipe => (
                    <li key={recipe.id}>
                        <Separator />
                        <div className="p-2">
                            <RecipeOverview recipe={recipe} />
                            <Button variant="outline" className="mt-2" onClick={() => nav(`/recipes/${recipe.id}`)}>
                                View Recipe
                            </Button>
                        </div>
                    </li>
                ))}
            </ul>
            <Button onClick={handleAddRecipe}>Add Recipe</Button>
            <ChatDrawer
                open={drawerOpen}
                onOpenChange={(open) => setDrawerOpen(open)}
                mode="suggest"
                onSend={handleSuggestRecipe}
                loading={startThreadPending}
                error={startThreadError?.message}
            />
        </div>
    )
}