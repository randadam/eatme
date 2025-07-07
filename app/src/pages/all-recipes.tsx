import { useAllRecipes, useStartSuggestionThread } from "@/features/recipe/hooks"
import type api from "@/api"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"
import { ChatDrawer } from "@/features/chat/chat-drawer"
import { useState } from "react"

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
            <h1>Recipes</h1>
            <ul>
                {recipesList.map(recipe => (
                    <>
                        <li key={recipe.id}>{recipe.title}</li>
                        <Button onClick={() => nav(`/recipe/${recipe.id}`)}>View</Button>
                    </>
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