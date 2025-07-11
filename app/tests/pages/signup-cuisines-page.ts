import { expect, Page } from "@playwright/test";

export class SignupCuisinesPage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/signup\/cuisines/i)
    }

    async goto() {
        await this.page.goto("/signup/cuisines")
    }

    async selectCuisine(cuisine: string) {
        await this.page.getByRole("checkbox", { name: cuisine }).click()
    }

    async submit() {
        await this.page.getByRole("button", { name: /next/i }).click()
    }
}