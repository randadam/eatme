import { expect, Page } from "@playwright/test";

export class SignupProfilePage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/signup\/profile/i)
    }

    async goto() {
        await this.page.goto("/signup/profile")
    }

    async fillName(name: string) {
        await this.page.getByRole("textbox", { name: /name/i }).fill(name)
    }

    async submit() {
        await this.page.getByRole("button", { name: /next/i }).click()
    }
}