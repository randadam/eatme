import { expect, Page } from "@playwright/test"

export class SignupAllergiesPage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/signup\/allergies/i)
    }

    async selectAllergy(allergy: string) {
        await this.page.getByTestId(allergy).click()
    }

    async submit() {
        await this.page.getByRole("button", { name: /next/i }).click()
    }
}