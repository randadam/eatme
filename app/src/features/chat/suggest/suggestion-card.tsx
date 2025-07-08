import type api from "@/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { RecipeAccordion } from "@/features/recipe/recipe-accordion";
import { Loader2 } from "lucide-react";

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
                    <Button variant="outline" onClick={reject} disabled={eitherLoading}>
                        {rejectLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                        Next
                    </Button>
                    <Button onClick={accept} disabled={eitherLoading}>
                        {acceptLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                        Accept
                    </Button>
                </div>
            </CardFooter>
        </Card>
    )
}