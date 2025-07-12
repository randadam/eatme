import { FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { z } from "zod";
import { profileForm } from "./schemas/forms";
import type { Control } from "react-hook-form";

export interface BasicProfileFormProps {
    control: Control<z.infer<typeof profileForm>>
}

export function BasicProfileForm({ control }: BasicProfileFormProps) {
    return (
        <FormField
            control={control}
            name="name"
            render={({ field }) => (
                <FormItem>
                    <FormLabel>Name</FormLabel>
                    <FormDescription className="text-left">
                        What should we call you?
                        This can be anything you want, we're not the IRS.
                    </FormDescription>
                    <FormControl>
                        <Input placeholder="Name" autoComplete="nickname" {...field} />
                    </FormControl>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}