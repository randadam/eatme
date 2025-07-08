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
export type ModelsModifyChatRequest = {
    message: string;
};
export type ModelsMeasurementUnit = "g" | "ml" | "tsp" | "tbsp" | "cup" | "oz" | "lb";
export type ModelsIngredient = {
    name: string;
    quantity: number;
    unit: ModelsMeasurementUnit;
};
export type ModelsRecipeBody = {
    description: string;
    ingredients: ModelsIngredient[];
    servings: number;
    steps: string[];
    title: string;
    total_time_minutes: number;
};
export type ModelsModifyChatResponse = {
    needs_clarification: boolean;
    new_recipe: ModelsRecipeBody;
    response_text: string;
};
export type ModelsApiError = {
    code?: string;
    details?: string;
    field?: string;
    message?: string;
};
export type ModelsGeneralChatRequest = {
    message: string;
};
export type ModelsGeneralChatResponse = {
    response_text: string;
};
export type ModelsSuggestChatRequest = {
    message: string;
};
export type ModelsSuggestChatResponse = {
    new_recipe: ModelsRecipeBody;
    response_text: string;
    thread_id: string;
};
export type ModelsRecipeSuggestion = {
    accepted: boolean;
    created_at: string;
    id: string;
    response_text: string;
    suggestion: ModelsRecipeBody;
    thread_id: string;
    updated_at: string;
};
export type ModelsSuggestionThread = {
    created_at: string;
    id: string;
    original_prompt: string;
    suggestions: ModelsRecipeSuggestion[];
    updated_at: string;
};
export type ModelsUserRecipe = {
    created_at: string;
    description: string;
    global_recipe_id?: string;
    id: string;
    ingredients: ModelsIngredient[];
    is_favorite: boolean;
    latest_version_id: string;
    servings: number;
    steps: string[];
    title: string;
    total_time_minutes: number;
    updated_at: string;
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
 * Handle modifying a recipe
 */
export function modifyRecipe(recipeId: string, modelsModifyChatRequest: ModelsModifyChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsModifyChatResponse;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/chat/modify/recipes/${encodeURIComponent(recipeId)}`, oazapfts.json({
        ...opts,
        method: "PUT",
        body: modelsModifyChatRequest
    }));
}
/**
 * Handle general chat request
 */
export function generalChat(recipeId: string, modelsGeneralChatRequest: ModelsGeneralChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsGeneralChatResponse;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/chat/question/recipes/${encodeURIComponent(recipeId)}`, oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsGeneralChatRequest
    }));
}
/**
 * Handle starting a recipe suggestion chat
 */
export function suggestRecipe(modelsSuggestChatRequest: ModelsSuggestChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsSuggestChatResponse;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>("/chat/suggest", oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsSuggestChatRequest
    }));
}
/**
 * Get suggestion thread
 */
export function getSuggestionThread(threadId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsSuggestionThread;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/chat/suggest/${encodeURIComponent(threadId)}`, {
        ...opts
    });
}
/**
 * Handle accepting a recipe suggestion
 */
export function acceptRecipeSuggestion(threadId: string, suggestionId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsUserRecipe;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/chat/suggest/${encodeURIComponent(threadId)}/accept/${encodeURIComponent(suggestionId)}`, {
        ...opts,
        method: "POST"
    });
}
/**
 * Handle getting next recipe suggestion
 */
export function nextRecipeSuggestion(threadId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsRecipeSuggestion;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/chat/suggest/${encodeURIComponent(threadId)}/next`, {
        ...opts,
        method: "POST"
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
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
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
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>("/profile", oazapfts.json({
        ...opts,
        method: "PUT",
        body: modelsProfileUpdateRequest
    }));
}
/**
 * Get all recipes for user
 */
export function getAllRecipes(opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsUserRecipe[];
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>("/recipes", {
        ...opts
    });
}
/**
 * Get recipe by ID
 */
export function getRecipe(recipeId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsUserRecipe;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/recipes/${encodeURIComponent(recipeId)}`, {
        ...opts
    });
}
/**
 * Delete recipe
 */
export function deleteRecipe(recipeId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsUserRecipe;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/recipes/${encodeURIComponent(recipeId)}`, {
        ...opts,
        method: "DELETE"
    });
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
        data: ModelsApiError;
    } | {
        status: 409;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>("/signup", oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsSignupRequest
    }));
}
