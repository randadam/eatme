import { Navigate, useLocation, useRoutes } from "react-router-dom"
import { useUser } from "@/features/auth/hooks"
import { STEPS } from "@/features/auth/signup/constants"
import { Loader2 } from "lucide-react"
import AccountStep from "@/features/auth/signup/step-account"
import ProfileStep from "@/features/auth/signup/step-profile"
import SkillStep from "@/features/auth/signup/step-skill"
import PreferencesStep from "@/features/auth/signup/step-cuisines"
import DietStep from "@/features/auth/signup/step-diet"
import EquipmentStep from "@/features/auth/signup/step-equipment"
import AllergiesStep from "@/features/auth/signup/step-allergies"
import SignupSuccess from "@/features/auth/signup/step-success"

export default function SignupPage() {
  const { pathname } = useLocation()
  const { isAuthenticated, profile, isLoading } = useUser()

  const element = useRoutes([
    { path: "account", element: <AccountStep /> },
    { path: "profile", element: <ProfileStep /> },
    { path: "skill", element: <SkillStep /> },
    { path: "cuisines", element: <PreferencesStep /> },
    { path: "diet", element: <DietStep /> },
    { path: "equipment", element: <EquipmentStep /> },
    { path: "allergies", element: <AllergiesStep /> },
    { path: "done", element: <SignupSuccess /> },
  ])

  if (isLoading) {
    return (
      <div className="h-screen flex items-center justify-center">
        <Loader2 className="w-6 h-6 animate-spin text-muted-foreground" />
      </div>
    )
  }

  if (pathname === "/signup") {
    if (!isAuthenticated || !profile) {
      return <Navigate to={STEPS.account} replace />
    }
    if (profile?.setup_step === "done") {
      return <Navigate to="/" replace />
    }
    const currentStep = STEPS[profile.setup_step]
    if (!currentStep) {
      return <Navigate to={STEPS.account} replace />
    }
    return <Navigate to={currentStep} replace />
  }

  return (
    <div className="min-h-screen flex flex-col items-center p-4 space-y-2">
      {element}
    </div>
  )
}
