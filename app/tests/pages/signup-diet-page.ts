import { expect, Page } from "@playwright/test";

export class SignupDietPage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/signup\/diet/i)
    }

    async goto() {
        await this.page.goto("/signup/diet")
    }

    async selectDiet(diet: string) {
        await this.page.getByTestId(diet).click()
    }

    async submit() {
        await this.page.getByRole("button", { name: /next/i }).click()
    }
}