import { AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import { Form } from "@/components/ui/form";
import { BasicProfileForm } from "../forms/basic-profile-form";
import { useBasicProfileForm } from "../forms/hooks";
import type { BasicProfileFormValues } from "../forms/types";
import { toast } from "sonner";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { Button } from "@/components/ui/button";
import type { SaveProfileFn } from "../forms/types";
import LoaderButton from "@/components/shared/loader-button";

interface BasicProfileSectionProps {
    profile: BasicProfileFormValues
    onSave: SaveProfileFn
    isPending: boolean
}

export default function BasicProfileSection({ profile, onSave, isPending }: BasicProfileSectionProps) {
    const form = useBasicProfileForm({
        name: profile?.name ?? "",
    })
    const handleFormError = useFormErrorHandler(form)

    const onSubmit = (values: BasicProfileFormValues) => {
        onSave({ ...values }, {
            onSuccess: () => {
                toast.success('Basic profile details saved')
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <AccordionItem value="basic-profile">
            <AccordionTrigger data-testid="basic-profile-trigger">Basic Profile</AccordionTrigger>
            <AccordionContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                        <BasicProfileForm control={form.control} />
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