import { expect, Page } from "@playwright/test";

export class SignupDonePage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/signup\/done/i)
    }
}