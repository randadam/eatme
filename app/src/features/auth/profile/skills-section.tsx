import { toast } from "sonner";
import { useSkillForm } from "../forms/hooks";
import type { SkillFormValues } from "../forms/types";
import { FormErrorMessage, useFormErrorHandler } from "@/lib/error/error-provider";
import { Form } from "@/components/ui/form";
import { Button } from "@/components/ui/button";
import { AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";
import SkillForm from "../forms/skill-form";
import type { SaveProfileFn } from "../forms/types";
import LoaderButton from "@/components/shared/loader-button";

export interface SkillsSectionProps {
    profile: SkillFormValues
    onSave: SaveProfileFn
    isPending: boolean
}

export default function SkillsSection({ profile, onSave, isPending }: SkillsSectionProps) {
    const form = useSkillForm({
        skill: profile?.skill ?? "beginner",
    })
    const handleFormError = useFormErrorHandler(form)

    const onSubmit = (values: SkillFormValues) => {
        onSave({ ...values }, {
            onSuccess: () => {
                toast.success('Skills saved')
            },
            onError: (err) => handleFormError(err),
        })
    }

    return (
        <AccordionItem value="skills">
            <AccordionTrigger>Skills</AccordionTrigger>
            <AccordionContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                        <SkillForm control={form.control} showTitle={false}/>
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
