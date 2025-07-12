import { useForm } from "react-hook-form";
import { profileForm, skillForm } from "./schemas/forms";
import { zodResolver } from "@hookform/resolvers/zod";
import type { BasicProfileFormValues, SkillFormValues } from "./types";

export const useBasicProfileForm = (initialValues: BasicProfileFormValues) => {
    const form = useForm<BasicProfileFormValues>({
        resolver: zodResolver(profileForm),
        defaultValues: initialValues,
    })
    return form
}

export const useSkillForm = (initialValues: SkillFormValues) => {
    const form = useForm<SkillFormValues>({
        resolver: zodResolver(skillForm),
        defaultValues: initialValues,
    })
    return form
}
    