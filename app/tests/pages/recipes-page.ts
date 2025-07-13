import { expect, Page } from "@playwright/test";

export class RecipesPage {
    constructor(private page: Page) {}

    async expectToBeHere() {
        await expect(this.page).toHaveURL(/\/recipes/i)
    }
}
    