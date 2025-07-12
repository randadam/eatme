import { AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Form } from "@/components/ui/form";
import AllergiesForm from "../forms/allergies-form";
import { useAllergiesForm } from "../forms/hooks";
import type { AllergiesFormValues } from "../forms/types";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { Button } from "@/components/ui/button";
import type { SaveProfileFn } from "../forms/types";
import LoaderButton from "@/components/shared/loader-button";
import type { ModelsAllergy } from "@/api/client";

interface AllergiesSectionProps {
    profile: AllergiesFormValues
    onSave: SaveProfileFn
    isPending: boolean
}

export default function AllergiesSection({ profile, onSave, isPending }: AllergiesSectionProps) {
    const form = useAllergiesForm({
        allergies: profile?.allergies ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    const onSubmit = (values: AllergiesFormValues) => {
        const req = {
            allergies: values.allergies as ModelsAllergy[],
        }
        onSave(req, {
            onSuccess: () => {
                toast.success('Allergies saved')
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <AccordionItem value="allergies">
            <AccordionTrigger>Allergies</AccordionTrigger>
            <AccordionContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                        <AllergiesForm control={form.control} showTitle={false}/>
                        <FormErrorMessage form={form} />
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