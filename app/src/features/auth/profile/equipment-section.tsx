import { AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Form } from "@/components/ui/form";
import EquipmentForm from "../forms/equipment-form";
import { useEquipmentForm } from "../forms/hooks";
import type { EquipmentFormValues } from "../forms/types";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { Button } from "@/components/ui/button";
import type { SaveProfileFn } from "../forms/types";
import LoaderButton from "@/components/shared/loader-button";

interface EquipmentSectionProps {
    profile: EquipmentFormValues
    onSave: SaveProfileFn
    isPending: boolean
}

export default function EquipmentSection({ profile, onSave, isPending }: EquipmentSectionProps) {
    const form = useEquipmentForm({
        equipment: profile?.equipment ?? [],
    })
    const handleFormError = useFormErrorHandler(form)

    const onSubmit = (values: EquipmentFormValues) => {
        const req = {
            equipment: values.equipment,
        }
        onSave(req, {
            onSuccess: () => {
                toast.success('Equipment saved')
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <AccordionItem value="equipment">
            <AccordionTrigger data-testid="equipment-trigger">Equipment</AccordionTrigger>
            <AccordionContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                        <EquipmentForm control={form.control} showTitle={false}/>
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