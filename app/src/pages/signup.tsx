import { Outlet, Navigate, useLocation } from "react-router-dom"
import { useUser } from "@/features/auth/hooks"
import { STEPS } from "@/features/auth/signup/constants"

export default function SignupPage() {
  const { pathname } = useLocation()
  const { isAuthenticated, profile, isLoading } = useUser()

  if (isLoading) return null

  if (pathname === "/signup") {
    if (!isAuthenticated || !profile) {
      return <Navigate to={STEPS.account} replace />
    }
    if (profile?.setup_step === "done") {
      return <Navigate to="/" replace />
    }
    const currentStep = STEPS[profile.setup_step]
    if (!currentStep) {
      console.warn(`Invalid setup step: ${profile.setup_step}`)
      return <Navigate to={STEPS.account} replace />
    }
    return <Navigate to={currentStep} replace />
  }

  return (
    <div className="min-h-screen flex flex-col items-center p-4 space-y-2">
      <Outlet />
    </div>
  )
}
