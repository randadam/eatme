import { test, expect } from "@playwright/test"
import { SignupAccountPage } from "./pages/signup-account-page"
import { SignupProfilePage } from "./pages/signup-profile-page"
import { SignupSkillPage } from "./pages/signup-skill-page"
import { SignupCuisinesPage } from "./pages/signup-cuisines-page"
import { SignupDietPage } from "./pages/signup-diet-page"
import { SignupEquipmentPage } from "./pages/signup-equipment-page"
import { SignupDonePage } from "./pages/signup-done-page"
import { SignupAllergiesPage } from "./pages/signup-allergies-page"

test("sanity check", async ({ page }) => {
    await page.goto("/signup")
    await expect(page).toHaveTitle("Eat Me")
})

test("full signup flow", async ({ page }) => {
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
    await cuisinesPage.selectCuisine("italian")
    await cuisinesPage.selectCuisine("mexican")
    await cuisinesPage.submit()

    const dietPage = new SignupDietPage(page)
    await dietPage.expectToBeHere()
    await dietPage.selectDiet("high_protein")
    await dietPage.submit()

    const allergiesPage = new SignupAllergiesPage(page)
    await allergiesPage.expectToBeHere()
    await allergiesPage.selectAllergy("dairy")
    await allergiesPage.submit()

    const equipmentPage = new SignupEquipmentPage(page)
    await equipmentPage.expectToBeHere()
    await equipmentPage.selectEquipment("stove")
    await equipmentPage.selectEquipment("oven")
    await equipmentPage.selectEquipment("microwave")
    await equipmentPage.submit()

    const donePage = new SignupDonePage(page)
    await donePage.expectToBeHere()

    // TODO(adam): go to profile page and verify it matches selections
})