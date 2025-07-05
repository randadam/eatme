import { lazy } from "react"
import { createBrowserRouter } from "react-router-dom"
import RootLayout from "@/layouts/root-layout"
import AccountStep from "./features/auth/signup/step-account"
import ProfileStep from "./features/auth/signup/step-profile"
import PreferencesStep from "./features/auth/signup/step-preferences"
import SignupSuccess from "./features/auth/signup/step-success"
import EquipmentStep from "./features/auth/signup/step-equipment"

const Signup = lazy(() => import("@/features/auth/signup/routes"))

export const router = createBrowserRouter([
    {
        element: <RootLayout/>,
        children: [
            {
                path: "/signup",
                element: <Signup/>,
                children: [
                    { path: "account", element: <AccountStep/> },
                    { path: "profile", element: <ProfileStep/> },
                    { path: "preferences", element: <PreferencesStep/> },
                    { path: "success", element: <SignupSuccess/> },
                    { path: "equipment", element: <EquipmentStep/> },
                ],
            }
        ]
    }
])