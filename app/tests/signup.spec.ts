import { test, expect } from "@playwright/test"
import { SignupAccountPage } from "./pages/signup-account-page"
import { SignupProfilePage } from "./pages/signup-profile-page"
import { SignupSkillPage } from "./pages/signup-skill-page"
import { SignupCuisinesPage } from "./pages/signup-cuisines-page"

test("sanity check", async ({ page }) => {
    await page.goto("/signup")
    await expect(page).toHaveTitle("Eat Me")
})

test("account step", async ({ page }) => {
    const randomSuffix = Math.random().toString(36).substring(2, 8)

    const accountPage = new SignupAccountPage(page)
    await accountPage.goto()
    await accountPage.fillEmail(`test-${randomSuffix}@example.com`)
    await accountPage.fillPassword("password")
    await accountPage.fillConfirmPassword("password")
    await accountPage.submit()

    const profilePage = new SignupProfilePage(page)
    await profilePage.expectToBeHere()
    await profilePage.fillName("Testi")
    await profilePage.submit()

    const skillsPage = new SignupSkillPage(page)
    await skillsPage.expectToBeHere()
    await skillsPage.selectSkill("chef")
    await skillsPage.submit()

    const cuisinesPage = new SignupCuisinesPage(page)
    await cuisinesPage.expectToBeHere()
})