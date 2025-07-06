import { Outlet, Navigate, useLocation, useNavigate } from "react-router-dom"
import { useUser } from "../hooks"
import { useEffect, useRef } from "react"

export const STEPS = {
  account: "/signup/account",
  profile: "/signup/profile",
  skill: "/signup/skill",
  cuisines: "/signup/cuisines",
  diet: "/signup/diet",
  equipment: "/signup/equipment",
  allergies: "/signup/allergies",
}

export default function SignupLayout() {
  const { pathname } = useLocation()
  const nav = useNavigate()
  const { isAuthenticated, profile } = useUser()

  const hasRedirected = useRef(false)

  useEffect(() => {
    if (!isAuthenticated || !profile || hasRedirected.current) {
      return
    }
    if (profile.setup_step === "done") {
      hasRedirected.current = true
      nav("/", { replace: true })
    }
    const target = STEPS[profile.setup_step as keyof typeof STEPS]
    if (target && target !== pathname) {
      hasRedirected.current = true
      nav(target, { replace: true })
    }
  }, [isAuthenticated, profile, pathname, nav])

  if (pathname === "/signup") return <Navigate to="account" replace />

  return (
    <div className="min-h-screen flex flex-col items-center p-4 space-y-2">
      <Outlet />
    </div>
  )
}
