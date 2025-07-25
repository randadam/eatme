import { useModifyRecipe, useRecipe } from "@/features/recipe/hooks"
import { useParams } from "react-router-dom"
import { NotFoundPage } from "./not-found"
import { ErrorPage } from "./error"
import RecipeSkeleton from "@/features/recipe/recipe-skeleton"
import { ChatBody } from "@/features/chat/chat-body"
import FocusedLayout from "@/layouts/focused-layout"
import BottomSheet from "@/components/shared/bottom-sheet"
import LoaderButton from "@/components/shared/loader-button"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"
import { ArrowLeft } from "lucide-react"
import { useState } from "react"
import { Separator } from "@/components/ui/separator"
import { FullRecipe } from "@/features/recipe/recipe-full"

export default function ModifyRecipe() {
    const recipeId = useParams().recipeId
    const nav = useNavigate()
    const [chatOpen, setChatOpen] = useState(false)

    const { recipe, isLoading, error, notFound } = useRecipe(recipeId!)
    const {
        modifyRecipe,
        modifyRecipePending,
        modifyRecipeError,
        proposedRecipe,
        reject,
        rejectPending,
        accept,
        acceptPending,
    } = useModifyRecipe(recipe?.id ?? "")

    const handleSend = (message: string) => {
        modifyRecipe({ prompt: message }, {
            onSuccess: () => {
                setChatOpen(false)
            }
        })
    }

    const handleReject = () => {
        reject()
    }

    const handleAccept = () => {
        accept(undefined, {
            onSuccess: () => {
                nav(`/recipes/${recipeId}`)
            }
        })
    }

    const onBack = () => {
        nav(`/recipes/${recipeId}`)
    }

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
        <FocusedLayout>
            <header className="sticky top-0 z-30 bg-background/90 backdrop-blur">
                <div className="flex items-center justify-between px-2 py-2">
                    <Button size="lg" variant="ghost" onClick={onBack} className="!pl-0 w-1/8">
                        <ArrowLeft />
                    </Button>
                    <h1 className="text-lg font-semibold wrap max-w-[calc(100%-6rem)]">Modify Recipe</h1>
                    <div className="w-1/8" />
                </div>
            </header>
            <div className="p-4">
                <FullRecipe recipe={proposedRecipe ?? recipe} />
            </div>
            <BottomSheet
                size={chatOpen ? "full" : "peek"}
                onSizeChange={s => {
                    if (s === "full") {
                        setChatOpen(true)
                    } else {
                        setChatOpen(false)
                    }
                }}
                header={
                    proposedRecipe ? (
                        <h2>Accept Changes?</h2>
                    ) : (
                        <h2>Modify Recipe</h2>
                    )
                }
                subHeader={
                    proposedRecipe ? (
                        <div className="flex justify-between pt-4 gap-2 w-full">
                            <LoaderButton
                                className="w-1/2"
                                isLoading={rejectPending}
                                onClick={handleReject}
                                disabled={rejectPending}
                            >
                                Reject
                            </LoaderButton>
                            <LoaderButton
                                className="w-1/2"
                                isLoading={acceptPending}
                                onClick={handleAccept}
                                disabled={acceptPending}
                            >
                                Accept
                            </LoaderButton>
                        </div>
                    ) : (
                        <p className="text-muted-foreground">
                            Let me know what you'd like to change
                        </p>
                    )
                }
                peekHeight={proposedRecipe ? 18 : 12}
                fullHeight={proposedRecipe ? 56 : 48}
            >
                {proposedRecipe && (
                    <>
                        <Separator className="my-2" />
                        <div className="flex flex-col justify-between text-center py-4">
                            <h3 className="text-md pb-2">
                                Not what you're looking for?
                            </h3>
                            <p className="text-muted-foreground">
                                Try asking me again.
                            </p>
                        </div>
                    </>
                )}
                <ChatBody
                    onSend={handleSend}
                    loading={modifyRecipePending}
                    error={modifyRecipeError?.message}
                />
            </BottomSheet>
        </FocusedLayout>
    )
}