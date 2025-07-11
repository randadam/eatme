import { expect, Page } from "@playwright/test";

export class SignupSkillPage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/signup\/skill/i)
    }

    async goto() {
        await this.page.goto("/signup/skill")
    }

    async selectSkill(skill: string) {
        await this.page.getByRole("radio", { name: skill }).click()
    }

    async submit() {
        await this.page.getByRole("button", { name: /next/i }).click()
    }
}