import type z from "zod";
import type { allergiesForm, cuisinesForm, dietForm, equipmentForm, profileForm, skillForm } from "./schemas/forms";
import type { useSaveProfile } from "../hooks";

export type BasicProfileFormValues = z.infer<typeof profileForm>
export type SkillFormValues = z.infer<typeof skillForm>
export type CuisinesFormValues = z.infer<typeof cuisinesForm>
export type DietFormValues = z.infer<typeof dietForm>
export type AllergiesFormValues = z.infer<typeof allergiesForm>
export type EquipmentFormValues = z.infer<typeof equipmentForm>

export type SaveProfileFn = ReturnType<typeof useSaveProfile>["saveProfile"]