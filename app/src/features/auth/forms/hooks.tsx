import { useForm } from "react-hook-form";
import { profileForm } from "./schemas/forms";
import { zodResolver } from "@hookform/resolvers/zod";
import type { BasicProfileFormValues } from "./types";

export const useBasicProfileForm = (initialValues: BasicProfileFormValues) => {
    const form = useForm<BasicProfileFormValues>({
        resolver: zodResolver(profileForm),
        defaultValues: initialValues,
    })
    return form
}
