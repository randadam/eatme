import type api from "@/api";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
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
                <div className="text-left space-y-2">
                    <p className="font-semibold">Description:</p>
                    <p>{recipe.description}</p>
                </div>
                <div className="text-left flex space-x-2">
                    <p className="font-semibold">Servings:</p>
                    <p>{recipe.servings}</p>
                </div>
                <div className="text-left flex space-x-2">
                    <p className="font-semibold">Total time:</p>
                    <p>{recipe.total_time_minutes} minutes</p>
                </div>
            </CardContent>
            <CardFooter>
                <div className="flex justify-between space-x-2 w-full">
                    <Button onClick={reject} disabled={eitherLoading}>
                        {rejectLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                        Reject
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