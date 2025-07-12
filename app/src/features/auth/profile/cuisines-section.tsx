import { AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Form } from "@/components/ui/form";
import CuisinesForm from "../forms/cuisines-form";
import { useCuisinesForm } from "../forms/hooks";
import type { CuisinesFormValues } from "../forms/types";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { Button } from "@/components/ui/button";
import type { SaveProfileFn } from "../forms/types";
import LoaderButton from "@/components/shared/loader-button";
import type { ModelsCuisine } from "@/api/client";

interface CuisinesSectionProps {
    profile: CuisinesFormValues
    onSave: SaveProfileFn
    isPending: boolean
}

export default function CuisinesSection({ profile, onSave, isPending }: CuisinesSectionProps) {
    const form = useCuisinesForm({
        cuisines: profile?.cuisines ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    const onSubmit = (values: CuisinesFormValues) => {
        const req = {
            cuisines: values.cuisines as ModelsCuisine[],
        }
        onSave(req, {
            onSuccess: () => {
                toast.success('Cuisines saved')
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <AccordionItem value="cuisines">
            <AccordionTrigger>Cuisines</AccordionTrigger>
            <AccordionContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                        <CuisinesForm control={form.control} showTitle={false}/>
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