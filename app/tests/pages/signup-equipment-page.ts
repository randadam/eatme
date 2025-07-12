import { expect, Page } from "@playwright/test";

export class SignupEquipmentPage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/signup\/equipment/i)
    }

    async goto() {
        await this.page.goto("/signup/equipment")
    }

    async selectEquipment(equipment: string) {
        await this.page.getByTestId(equipment).click()
    }

    async submit() {
        await this.page.getByRole("button", { name: /next/i }).click()
    }
}