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
export type ModelsLoginRequest = {
    /** User's email address */
    email: string;
    /** User's password */
    password: string;
};
export type ModelsLoginResponse = {
    /** Access token for user */
    token: string;
};
export type ModelsApiError = {
    code: string;
    details?: string;
    field?: string;
    message: string;
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
    setup_step?: ModelsSetupStep;
    /** User's skill level */
    skill?: ModelsSkill;
};
export type ModelsMeasurementUnit = "g" | "ml" | "tsp" | "tbsp" | "cup" | "oz" | "lb";
export type ModelsIngredient = {
    name: string;
    quantity: number;
    unit: ModelsMeasurementUnit;
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
    thread_id: string;
    title: string;
    total_time_minutes: number;
    updated_at: string;
    user_id: string;
};
export type ModelsModifyRecipeViaChatRequest = {
    prompt: string;
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
export type ModelsStartSuggestionThreadRequest = {
    prompt: string;
};
export type ModelsChatMessage = {
    message: string;
    source: string;
};
export type ModelsRecipeBody = {
    description: string;
    ingredients: ModelsIngredient[];
    servings: number;
    steps: string[];
    title: string;
    total_time_minutes: number;
};
export type ModelsRecipeSuggestion = {
    accepted: boolean;
    created_at: string;
    id: string;
    rejected: boolean;
    response_text: string;
    suggestion: ModelsRecipeBody;
    thread_id: string;
    updated_at: string;
};
export type ModelsThreadState = {
    chat_history: ModelsChatMessage[];
    created_at: string;
    current_prompt: string;
    current_recipe?: ModelsRecipeBody;
    id: string;
    original_prompt: string;
    recipe_id?: string;
    suggestions: ModelsRecipeSuggestion[];
    updated_at: string;
};
export type ModelsAnswerCookingQuestionRequest = {
    question: string;
};
export type ModelsAnswerCookingQuestionResponse = {
    answer: string;
};
export type ModelsGetNewSuggestionsRequest = {
    prompt?: string;
};
/**
 * Log in
 */
export function login(modelsLoginRequest: ModelsLoginRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsLoginResponse;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>("/login", oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsLoginRequest
    }));
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
 * Modify a recipe via chat
 */
export function modifyRecipe(recipeId: string, modelsModifyRecipeViaChatRequest: ModelsModifyRecipeViaChatRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsUserRecipe;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/recipes/${encodeURIComponent(recipeId)}/modify/chat`, oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsModifyRecipeViaChatRequest
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
/**
 * Start a new suggestion thread
 */
export function startSuggestionThread(modelsStartSuggestionThreadRequest: ModelsStartSuggestionThreadRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsThreadState;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>("/thread/suggest", oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsStartSuggestionThreadRequest
    }));
}
/**
 * Get a thread
 */
export function getThread(threadId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsThreadState;
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
    }>(`/thread/${encodeURIComponent(threadId)}`, {
        ...opts
    });
}
/**
 * Accept a suggestion
 */
export function acceptSuggestion(threadId: string, suggestionId: string, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsUserRecipe;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/thread/${encodeURIComponent(threadId)}/accept/${encodeURIComponent(suggestionId)}`, {
        ...opts,
        method: "POST"
    });
}
/**
 * Answer a cooking question
 */
export function answerCookingQuestion(threadId: string, modelsAnswerCookingQuestionRequest: ModelsAnswerCookingQuestionRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsAnswerCookingQuestionResponse;
    } | {
        status: 400;
        data: ModelsApiError;
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/thread/${encodeURIComponent(threadId)}/question`, oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsAnswerCookingQuestionRequest
    }));
}
/**
 * Get new suggestions
 */
export function getNewSuggestions(threadId: string, modelsGetNewSuggestionsRequest: ModelsGetNewSuggestionsRequest, opts?: Oazapfts.RequestOpts) {
    return oazapfts.fetchJson<{
        status: 200;
        data: ModelsRecipeSuggestion[];
    } | {
        status: 401;
        data: ModelsApiError;
    } | {
        status: 404;
        data: ModelsApiError;
    } | {
        status: 500;
        data: ModelsApiError;
    }>(`/thread/${encodeURIComponent(threadId)}/suggest`, oazapfts.json({
        ...opts,
        method: "POST",
        body: modelsGetNewSuggestionsRequest
    }));
}
