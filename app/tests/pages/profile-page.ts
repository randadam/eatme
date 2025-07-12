import { expect, Page } from "@playwright/test";
import { cuisineOptions } from "../../src/features/auth/forms/cuisines-form"
import { dietOptions } from "../../src/features/auth/forms/diet-form"
import { equipmentOptions } from "../../src/features/auth/forms/equipment-form"
import { allergyOptions } from "../../src/features/auth/forms/allergies-form"

export class ProfilePage {
    constructor(private page: Page) {}

    async goto() {
        await this.page.goto("/profile");
    }

    async openBasicProfileSection() {
        await this.page.getByRole("button", { name: /basic profile/i }).click();
    }

    async expectName(name: string) {
        await this.openBasicProfileSection()
        await expect(this.page.getByRole("textbox", { name: /name/i })).toHaveValue(name);
    }

    async openSkillSection() {
        await this.page.getByRole("button", { name: /skill/i }).click();
    }

    async expectSkill(skill: string) {
        await this.openSkillSection()
        await expect(this.page.getByRole("radio", { name: skill })).toBeChecked();
    }

    async openCuisinesSection() {
        await this.page.getByRole("button", { name: /cuisines/i }).click();
    }

    async expectCuisines(cuisines: string[]) {
        await this.openCuisinesSection()
        const items = getMultiSelection(cuisineOptions.map(cuisine => cuisine.value), cuisines)
        for (const cuisine of items) {
            if (cuisine.checked) {
                await expect(this.page.getByTestId(cuisine.value)).toBeChecked();
            } else {
                await expect(this.page.getByTestId(cuisine.value)).not.toBeChecked();
            }
        }
    }

    async openDietSection() {
        await this.page.getByRole("button", { name: /diet/i }).click();
    }

    async expectDiets(diets: string[]) {
        await this.openDietSection()
        const items = getMultiSelection(dietOptions.map(diet => diet.value), diets)
        for (const diet of items) {
            if (diet.checked) {
                await expect(this.page.getByTestId(diet.value)).toBeChecked();
            } else {
                await expect(this.page.getByTestId(diet.value)).not.toBeChecked();
            }
        }
    }

    async openAllergiesSection() {
        await this.page.getByRole("button", { name: /allergies/i }).click();
    }

    async expectAllergies(allergies: string[]) {
        await this.openAllergiesSection()
        const items = getMultiSelection(allergyOptions.map(allergy => allergy.value), allergies)
        for (const allergy of items) {
            if (allergy.checked) {
                await expect(this.page.getByTestId(allergy.value)).toBeChecked();
            } else {
                await expect(this.page.getByTestId(allergy.value)).not.toBeChecked();
            }
        }
    }

    async openEquipmentSection() {
        await this.page.getByRole("button", { name: /equipment/i }).click();
    }

    async expectEquipment(equipments: string[]) {
        await this.openEquipmentSection()
        const items = getMultiSelection(equipmentOptions.map(equipment => equipment.value), equipments)
        for (const equipment of items) {
            if (equipment.checked) {
                await expect(this.page.getByTestId(equipment.value)).toBeChecked();
            } else {
                await expect(this.page.getByTestId(equipment.value)).not.toBeChecked();
            }
        }
    }
}

interface MultiSelectItem {
    value: string
    checked: boolean
}

function getMultiSelection(options: string[], selected: string[]): MultiSelectItem[] {
    const selectedSet = new Set(selected)
    return options.map(option => ({
        value: option,
        checked: selectedSet.has(option),
    }))
}
