import { Outlet, Navigate, useLocation } from "react-router-dom"
import WizardNav from "./wizard-nav"

export const STEPS = [
  "/signup/account",
  "/signup/profile",
  "/signup/skill",
  "/signup/cuisines",
  "/signup/diet",
  "/signup/equipment",
  "/signup/allergies",
] as const

export default function SignupLayout() {
  const { pathname } = useLocation()

  // Redirect bare /signup → first step
  if (pathname === "/signup") return <Navigate to="account" replace />

  return (
    <div className="min-h-screen flex flex-col p-4 space-y-2">
      <Outlet />
      <WizardNav />
    </div>
  )
}
