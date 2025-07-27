import type api from "@/api";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { RecipeAccordion } from "@/features/recipe/recipe-accordion";
import LoaderButton from "@/components/shared/loader-button";
import { Skeleton } from "@/components/ui/skeleton";

export interface SuggestionCardProps {
    suggestion: api.ModelsRecipeSuggestion
    reject: () => void
    rejectLoading: boolean
    accept: () => void
    acceptLoading: boolean
}

export default function SuggestionCard({ suggestion, reject, accept, rejectLoading, acceptLoading }: SuggestionCardProps) {
    console.log('suggestion', suggestion)
    const recipe = suggestion.suggestion
    const eitherLoading = rejectLoading || acceptLoading
    return (
        <Card>
            <CardHeader>
                <CardTitle className="text-center">{recipe.title}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
                <div className="flex justify-center">
                    {recipe.image_url ? (
                        <img className="h-64 w-full" src={recipe.image_url} alt={recipe.title} />
                    ) : (
                        <Skeleton className="h-64 w-full" />
                    )}
                </div>
                <RecipeAccordion id={suggestion.id} recipe={recipe}/>
            </CardContent>
            <CardFooter>
                <div className="flex justify-between space-x-2 w-full">
                    <LoaderButton variant="outline" onClick={reject} isLoading={rejectLoading} disabled={eitherLoading}>
                        Next
                    </LoaderButton>
                    <LoaderButton onClick={accept} isLoading={acceptLoading} disabled={eitherLoading}>
                        Accept
                    </LoaderButton>
                </div>
            </CardFooter>
        </Card>
    )
}