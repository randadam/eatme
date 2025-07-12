import { AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Form } from "@/components/ui/form";
import DietForm from "../forms/diet-form";
import { useDietForm } from "../forms/hooks";
import type { DietFormValues, SaveProfileFn } from "../forms/types";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { Button } from "@/components/ui/button";
import LoaderButton from "@/components/shared/loader-button";
import type { ModelsDiet } from "@/api/client";

export interface DietSectionProps {
    profile: DietFormValues
    onSave: SaveProfileFn
    isPending: boolean
}

export default function DietSection({ profile, onSave, isPending }: DietSectionProps) {
    const form = useDietForm({
        diet: profile?.diet ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    const onSubmit = (values: DietFormValues) => {
        const req = {
            diet: values.diet.map((diet) => diet) as ModelsDiet[],
        }
        onSave(req, {
            onSuccess: () => {
                toast.success('Diet saved')
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <AccordionItem value="diet">
            <AccordionTrigger>Diet</AccordionTrigger>
            <AccordionContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                        <DietForm control={form.control} showTitle={false}/>
                        <FormErrorMessage form={form}/>
                        <div className="flex justify-end gap-2">
                            <Button type="button" variant="ghost" onClick={() => form.reset()}>
                                Cancel
                            </Button>
                            <LoaderButton type="submit" isLoading={isPending}>
                                Save
                            </LoaderButton>
                        </div>
                    </form>
                </Form>
            </AccordionContent>
        </AccordionItem>
    )
}