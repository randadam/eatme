import z from "zod"
import * as a from "./atoms"

// STEP 1 – account credentials
export const accountForm = z.object({
  email: a.email,
  password: a.password,
  confirmPassword: a.password,
}).refine(
  d => d.password === d.confirmPassword,
  { path: ["confirmPassword"], message: "Passwords don't match" },
)

export type AccountFormValues = z.infer<typeof accountForm>

// STEP 2 – profile
export const profileForm = z.object({
  name: a.name,
})
export type ProfileFormValues = z.infer<typeof profileForm>

// STEP 3 - skill
export const skillForm = z.object({
  skill: a.skill,
})
export type SkillFormValues = z.infer<typeof skillForm>

// STEP 4 – cuisines
export const cuisinesForm = z.object({
  cuisines: a.cuisines,
})
export type CuisinesFormValues = z.infer<typeof cuisinesForm>

// STEP 5 – diet
export const dietForm = z.object({
  diet: a.diets,
})
export type DietFormValues = z.infer<typeof dietForm>

// STEP 6 – equipment
export const equipmentForm = z.object({
  equipment: a.equipment,
})
export type EquipmentFormValues = z.infer<typeof equipmentForm>

// STEP 7 – allergies
export const allergiesForm = z.object({
  allergies: a.allergies,
})
export type AllergiesFormValues = z.infer<typeof allergiesForm>