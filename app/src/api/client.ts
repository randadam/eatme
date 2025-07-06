/**
 * EatMe API
 * 1.0
 * DO NOT MODIFY - This file has been generated using oazapfts.
 * See https://www.npmjs.com/package/oazapfts
 */
import * as Oazapfts from "@oazapfts/runtime";
import * as QS from "@oazapfts/runtime/query";
export const defaults: Oazapfts.Defaults<Oazapfts.CustomHeaders> = {
    headers: {},
    baseUrl: "//localhost:8080",
};
const oazapfts = Oazapfts.runtime(defaults);
export const servers = {
    server1: "//localhost:8080"
};
export type ModelsAllergy = "dairy" | "eggs" | "fish" | "gluten" | "peanuts" | "soy" | "tree_nuts" | "wheat";
export type ModelsCuisine = "american" | "british" | "chinese" | "french" | "german" | "indian" | "italian" | "japanese" | "mexican" | "spanish" | "thai" | "vietnamese";
export type ModelsDiet = "vegetarian" | "vegan" | "keto" | "paleo" | "low_carb" | "high_protein";
export type ModelsEquipment = "stove" | "oven" | "microwave" | "toaster" | "grill" | "smoker" | "slow_cooker" | "pressure_cooker" | "sous_vide";
export type ModelsSetupStep = "profile" | "skill" | "cuisines" | "diet" | "equipment" | "allergies" | "done";
export type ModelsSkill = "beginner" | "intermediate" | "advanced" | "chef";
export type ModelsProfile = {
    /** User's allergies */
    allergies?: ModelsAllergy[];
    /** User's cuisines */
    cuisines?: ModelsCuisine[];
    /** User's diet restrictions */
    diet?: ModelsDiet[];
    /** User's equipment */
    equipment?: ModelsEquipment[];
    /** User's name */
    name?: string;
    /** Setup Step */
    setup_step: ModelsSetupStep;
    /** User's skill level */
    skill?: ModelsSkill;
};
export type ModelsUnauthorizedResponse = {
    /** Error message */
    error: string;
};
export type ModelsInternalServerErrorResponse = {
    /** Error message */
    error: string;
};
export type ModelsBadRequestResponse = {
    /** Error message */
    error: string;
};
export type ModelsSignupRequest = {
    /** User's email address */
    email: string;
    /** User's password */
    password: string;
};
export type ModelsSignupResponse = {
    /** Access token for user */
    token: string;
};
/**
 * Generate a recipe
 */
export function getGenerate(opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchText("/generate", {
        ...opts
    });
}
/**
 * Get user profile
 */
export function getProfile(opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsProfile;
    } | {
        status: 401;
        data: ModelsUnauthorizedResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>("/profile", {
        ...opts
    });
}
/**
 * Save user profile
 */
export function saveProfile(modelsProfile: ModelsProfile, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsProfile;
    } | {
        status: 400;
        data: ModelsBadRequestResponse;
    } | {
        status: 401;
        data: ModelsUnauthorizedResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>("/profile", oazapfts.json({
        ...opts,
        method: "PUT",
        body: modelsProfile
    }));
}
/**
 * Create a new user account
 */
export function signup(modelsSignupRequest: ModelsSignupRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsSignupResponse;
    } | {
        status: 400;
        data: ModelsBadRequestResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>("/signup", oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsSignupRequest
    }));
}
