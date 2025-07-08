import type api from "@/api"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion"
import { Separator } from "@/components/ui/separator"
import { RecipeOverview } from "./recipe-overview"

export interface RecipeAccordianProps {
    id: string
    recipe: api.ModelsRecipeBody
}

export function RecipeAccordion({ id, recipe }: RecipeAccordianProps) {
    return (
        <Accordion type="single" collapsible>
            <AccordionItem value={id}>
                <AccordionTrigger>
                    <RecipeOverview recipe={recipe} />
                </AccordionTrigger>
                <AccordionContent className="text-left space-y-2">
                    <Separator />
                    <h2 className="font-semibold">Ingredients:</h2>
                    <ul>
                        {recipe.ingredients.map(i => (
                            <li key={i.name} className="flex space-x-2">
                                <p>{i.quantity}</p>
                                <p>{i.unit}</p>
                                <p>{i.name}</p>
                            </li>
                        ))}
                    </ul>
                    <Separator />
                    <h2 className="font-semibold">Steps:</h2>
                    <ol>
                        {recipe.steps.map((s, i) => (
                            <li key={i}>{i + 1}. {s}</li>
                        ))}
                    </ol>
                </AccordionContent>
            </AccordionItem>
        </Accordion>
    )
}