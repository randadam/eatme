import { lazy } from "react"
import { createBrowserRouter } from "react-router-dom"
import RootLayout from "@/layouts/root-layout"
import AccountStep from "./features/auth/signup/step-account"
import ProfileStep from "./features/auth/signup/step-profile"
import PreferencesStep from "./features/auth/signup/step-cuisines"
import SignupSuccess from "./features/auth/signup/step-success"
import EquipmentStep from "./features/auth/signup/step-equipment"
import AllergiesStep from "./features/auth/signup/step-allergies"
import SkillStep from "./features/auth/signup/step-skill"
import DietStep from "./features/auth/signup/step-diet"

const Home = lazy(() => import("@/features/home/layout"))
const Signup = lazy(() => import("@/features/auth/signup/layout"))

export const router = createBrowserRouter([
    {
        element: <RootLayout/>,
        children: [
            {
                path: "/",
                element: <Home/>,
            },
            {
                path: "/signup",
                element: <Signup/>,
                children: [
                    { path: "account", element: <AccountStep/> },
                    { path: "profile", element: <ProfileStep/> },
                    { path: "skill", element: <SkillStep/> },
                    { path: "cuisines", element: <PreferencesStep/> },
                    { path: "diet", element: <DietStep/> },
                    { path: "equipment", element: <EquipmentStep/> },
                    { path: "allergies", element: <AllergiesStep/> },
                    { path: "done", element: <SignupSuccess/> },
                ],
            }
        ]
    }
])