import { z } from "zod"

export const email     = z.string().email()
export const password  = z.string().min(6)
export const name      = z.string().min(1)
export const skill     = z.enum(["beginner", "intermediate", "advanced", "chef"])
export const cuisines  = z.array(z.string()).min(1)
export const diets     = z.array(z.string())
export const equipment = z.array(z.string()).min(1)
export const allergies = z.array(z.string())
