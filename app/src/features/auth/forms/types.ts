import type z from "zod";
import type { profileForm } from "./schemas/forms";

export type BasicProfileFormValues = z.infer<typeof profileForm>