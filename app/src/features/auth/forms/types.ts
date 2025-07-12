import type z from "zod";
import type { profileForm, skillForm } from "./schemas/forms";
import type { useSaveProfile } from "../hooks";

export type BasicProfileFormValues = z.infer<typeof profileForm>
export type SkillFormValues = z.infer<typeof skillForm>

export type SaveProfileFn = ReturnType<typeof useSaveProfile>["saveProfile"]