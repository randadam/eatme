import { test, expect } from "@playwright/test"
import { SignupAccountPage } from "./pages/signup-account-page"
import { SignupProfilePage } from "./pages/signup-profile-page"
import { SignupSkillPage } from "./pages/signup-skill-page"
import { SignupCuisinesPage } from "./pages/signup-cuisines-page"
import { SignupDietPage } from "./pages/signup-diet-page"
import { SignupEquipmentPage } from "./pages/signup-equipment-page"
import { SignupDonePage } from "./pages/signup-done-page"
import { SignupAllergiesPage } from "./pages/signup-allergies-page"
import { ProfilePage } from "./pages/profile-page"
import { RecipesPage } from "./pages/recipes-page"

test.use({ viewport: { width: 390, height: 844 } }) // iPhone 11

test("full signup flow", async ({ page }) => {

    const randomSuffix = Math.random().toString(36).substring(2, 8)

    const accountPage = new SignupAccountPage(page)
    await accountPage.goto()
    await accountPage.fillEmail(`test-${randomSuffix}@example.com`)
    await accountPage.fillPassword("password")
    await accountPage.fillConfirmPassword("password")
    await accountPage.submit()

    const name = "Testi"
    const profileSetupPage = new SignupProfilePage(page)
    await profileSetupPage.expectToBeHere()
    await profileSetupPage.fillName(name)
    await profileSetupPage.submit()

    const skillPage = new SignupSkillPage(page)
    await skillPage.expectToBeHere()
    await skillPage.selectSkill("chef")
    await skillPage.submit()

    const cuisines = ["italian", "mexican"]
    const cuisinesPage = new SignupCuisinesPage(page)
    await cuisinesPage.expectToBeHere()
    for (const cuisine of cuisines) {
        await cuisinesPage.selectCuisine(cuisine)
    }
    await cuisinesPage.submit()

    const diets = ["high_protein", "keto"]
    const dietPage = new SignupDietPage(page)
    await dietPage.expectToBeHere()
    for (const diet of diets) {
        await dietPage.selectDiet(diet)
    }
    await dietPage.submit()

    const allergies = ["milk", "eggs"]
    const allergiesPage = new SignupAllergiesPage(page)
    await allergiesPage.expectToBeHere()
    for (const allergy of allergies) {
        await allergiesPage.selectAllergy(allergy)
    }
    await allergiesPage.submit()

    const equipment = ["stove", "oven", "microwave"]
    const equipmentPage = new SignupEquipmentPage(page)
    await equipmentPage.expectToBeHere()
    for (const e of equipment) {
        await equipmentPage.selectEquipment(e)
    }
    await equipmentPage.submit()

    const donePage = new SignupDonePage(page)
    await donePage.expectToBeHere()
    await donePage.clickGetStarted()

    const recipesPage = new RecipesPage(page)
    await recipesPage.expectToBeHere()

    const profileNavButton = await page.getByTestId("nav-profile")
    await profileNavButton.click()

    const profilePage = new ProfilePage(page)
    await profilePage.expectToBeHere()
    await profilePage.expectName(name)
    await profilePage.expectSkill("chef")
    await profilePage.expectCuisines(cuisines)
    await profilePage.expectDiets(diets)
    await profilePage.expectAllergies(allergies)
    await profilePage.expectEquipment(equipment)
})