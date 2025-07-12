import { useForm } from "react-hook-form";
import { allergiesForm, cuisinesForm, dietForm, equipmentForm, profileForm, skillForm } from "./schemas/forms";
import { zodResolver } from "@hookform/resolvers/zod";
import type { AllergiesFormValues, BasicProfileFormValues, CuisinesFormValues, DietFormValues, EquipmentFormValues, SkillFormValues } from "./types";

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
    
export const useCuisinesForm = (initialValues: CuisinesFormValues) => {
    const form = useForm<CuisinesFormValues>({
        resolver: zodResolver(cuisinesForm),
        defaultValues: initialValues,
    })
    return form
}

export const useDietForm = (initialValues: DietFormValues) => {
    const form = useForm<DietFormValues>({
        resolver: zodResolver(dietForm),
        defaultValues: initialValues,
    })
    return form
}
    
export const useAllergiesForm = (initialValues: AllergiesFormValues) => {
    const form = useForm<AllergiesFormValues>({
        resolver: zodResolver(allergiesForm),
        defaultValues: initialValues,
    })
    return form
}

export const useEquipmentForm = (initialValues: EquipmentFormValues) => {
    const form = useForm<EquipmentFormValues>({
        resolver: zodResolver(equipmentForm),
        defaultValues: initialValues,
    })
    return form
}
