import { test, expect } from "@playwright/test"

test("sanity check", async ({ page }) => {
    await page.goto("/signup")
    await expect(page).toHaveTitle("Eat Me")
})