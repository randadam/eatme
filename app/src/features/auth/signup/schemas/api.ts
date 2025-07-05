import { z } from "zod"
import * as a from "./atoms"

/** POST /api/users  – create account */
export const signupApi = z.object({
  email: a.email,
  password: a.password,
})

/** PATCH /api/users/me  – partial update */
export const profilePatchApi = z.object({
  name:      a.name.optional(),
  skill:     a.skill.optional(),
  allergies: a.allergies,
  cuisines:  a.cuisines,
  equipment: a.equipment,
}).partial()
.strict()

export type SignupApiPayload      = z.infer<typeof signupApi>
export type ProfilePatchApiInput  = z.infer<typeof profilePatchApi>
