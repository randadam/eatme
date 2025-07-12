import type api from "@/api";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { RecipeAccordion } from "@/features/recipe/recipe-accordion";
import LoaderButton from "@/components/shared/loader-button";

export interface SuggestionCardProps {
    suggestion: api.ModelsRecipeSuggestion
    reject: () => void
    rejectLoading: boolean
    accept: () => void
    acceptLoading: boolean
}

export default function SuggestionCard({ suggestion, reject, accept, rejectLoading, acceptLoading }: SuggestionCardProps) {
    const recipe = suggestion.suggestion
    const eitherLoading = rejectLoading || acceptLoading
    return (
        <Card>
            <CardHeader>
                <CardTitle className="text-center">{recipe.title}</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
                <RecipeAccordion id={suggestion.id} recipe={recipe}/>
            </CardContent>
            <CardFooter>
                <div className="flex justify-between space-x-2 w-full">
                    <LoaderButton variant="outline" onClick={reject} isLoading={eitherLoading}>
                        Next
                    </LoaderButton>
                    <LoaderButton onClick={accept} isLoading={eitherLoading}>
                        Accept
                    </LoaderButton>
                </div>
            </CardFooter>
        </Card>
    )
}