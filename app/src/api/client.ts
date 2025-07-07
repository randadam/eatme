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
export type ModelsGeneralChatRequest = {
    message: string;
};
export type ModelsGeneralChatResponse = {
    response_text: string;
};
export type ModelsBadRequestResponse = {
    /** Error message */
    error: string;
};
export type ModelsInternalServerErrorResponse = {
    /** Error message */
    error: string;
};
export type ModelsSuggestChatRequest = {
    message: string;
};
export type ModelsMeasurementUnit = "g" | "ml" | "tsp" | "tbsp" | "cup" | "oz" | "lb";
export type ModelsIngredient = {
    name: string;
    quantity: number;
    unit: ModelsMeasurementUnit;
};
export type ModelsRecipe = {
    description: string;
    id: string;
    ingredients: ModelsIngredient[];
    servings: number;
    steps: string[];
    title: string;
    total_time_minutes: number;
};
export type ModelsSuggestChatResponse = {
    new_recipe?: ModelsRecipe;
    response_text: string;
};
export type ModelsModifyChatRequest = {
    message: string;
};
export type ModelsModifyChatResponse = {
    needs_clarification?: boolean;
    new_recipe?: ModelsRecipe;
    response_text: string;
};
export type ModelsMealPlan = {
    id: string;
    recipes: ModelsRecipe[];
    user_id: string;
};
export type ModelsAllergy = "dairy" | "eggs" | "fish" | "gluten" | "peanuts" | "soy" | "tree_nuts" | "wheat";
export type ModelsCuisine = "american" | "british" | "chinese" | "french" | "german" | "indian" | "italian" | "japanese" | "mexican" | "spanish" | "thai" | "vietnamese";
export type ModelsDiet = "vegetarian" | "vegan" | "keto" | "paleo" | "low_carb" | "high_protein";
export type ModelsEquipment = "stove" | "oven" | "microwave" | "toaster" | "grill" | "smoker" | "slow_cooker" | "pressure_cooker" | "sous_vide";
export type ModelsSetupStep = "profile" | "skill" | "cuisines" | "diet" | "equipment" | "allergies" | "done";
export type ModelsSkill = "beginner" | "intermediate" | "advanced" | "chef";
export type ModelsProfile = {
    /** User's allergies */
    allergies: ModelsAllergy[];
    /** User's cuisines */
    cuisines: ModelsCuisine[];
    /** User's diet restrictions */
    diet: ModelsDiet[];
    /** User's equipment */
    equipment: ModelsEquipment[];
    /** User's name */
    name: string;
    /** Setup Step */
    setup_step: ModelsSetupStep;
    /** User's skill level */
    skill: ModelsSkill;
};
export type ModelsUnauthorizedResponse = {
    /** Error message */
    error: string;
};
export type ModelsProfileUpdateRequest = {
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
 * Handle general chat request
 */
export function generalChat(mealPlanId: string, modelsGeneralChatRequest: ModelsGeneralChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsGeneralChatResponse;
    } | {
        status: 400;
        data: ModelsBadRequestResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>(`/chat/plan/${encodeURIComponent(mealPlanId)}`, oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsGeneralChatRequest
    }));
}
/**
 * Handle recipe suggestion chat request
 */
export function suggestRecipe(mealPlanId: string, modelsSuggestChatRequest: ModelsSuggestChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsSuggestChatResponse;
    } | {
        status: 400;
        data: ModelsBadRequestResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>(`/chat/plan/${encodeURIComponent(mealPlanId)}/recipe`, oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsSuggestChatRequest
    }));
}
/**
 * Handle recipe modification chat request
 */
export function modifyRecipe(mealPlanId: string, recipeId: string, modelsModifyChatRequest: ModelsModifyChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsModifyChatResponse;
    } | {
        status: 400;
        data: ModelsBadRequestResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>(`/chat/plan/${encodeURIComponent(mealPlanId)}/recipe/${encodeURIComponent(recipeId)}`, oazapfts.json({
        ...opts,
        method: "PUT",
        body: modelsModifyChatRequest
    }));
}
/**
 * Create new meal plan
 */
export function createMealPlan(opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsMealPlan;
    } | {
        status: 400;
        data: ModelsBadRequestResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>("/meal/plan", {
        ...opts,
        method: "POST"
    });
}
/**
 * Get meal plan by ID
 */
export function getMealPlan(mealPlanId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsMealPlan;
    } | {
        status: 400;
        data: ModelsBadRequestResponse;
    } | {
        status: 500;
        data: ModelsInternalServerErrorResponse;
    }>(`/meal/plan/${encodeURIComponent(mealPlanId)}`, {
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
export function saveProfile(modelsProfileUpdateRequest: ModelsProfileUpdateRequest, opts?: Oazapfts.RequestOpts) {
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
        body: modelsProfileUpdateRequest
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
