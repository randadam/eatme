import type api from "@/api";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Separator } from "@/components/ui/separator";

interface Props {
    plan: api.ModelsMealPlan
}

export default function MealPlan({ plan }: Props) {
    if (plan.recipes.length === 0) {
        return <h1>No recipes found</h1>
    }
    return (
        <div>
            <h1>Meal Plan</h1>
            {plan.recipes.map(r => (
                <RecipeAccordion key={r.id} recipe={r} />
            ))}
        </div>
    )
}

interface RecipeProps {
    recipe: api.ModelsRecipe
}

function RecipeAccordion({ recipe }: RecipeProps) {
    return (
        <Accordion className="border rounded p-2 mt-2" type="single" collapsible>
            <AccordionItem value={recipe.id}>
                <AccordionTrigger>
                    <div className="flex flex-col space-y-2">
                        <h2>{recipe.title}</h2>
                        <p className="text-muted-foreground">{recipe.description}</p>
                        <div className="flex space-x-2">
                            <div className="flex space-x-2">
                                <p>Servings:</p>
                                <p className="text-muted-foreground">{recipe.servings}</p>
                            </div>
                            <div className="flex space-x-2">
                                <p>Time:</p>
                                <p className="text-muted-foreground">{recipe.total_time_minutes} minutes</p>
                            </div>
                        </div>
                    </div>
                </AccordionTrigger>
                <AccordionContent className="text-left space-y-2">
                    <Separator/>
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
                    <Separator/>
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