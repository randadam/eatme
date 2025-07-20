import { Outlet } from "react-router-dom"
import RequireFinishedSignup from "@/features/auth/require-finished-signup"

export default function RootLayout() {
  return (
    <RequireFinishedSignup>
      <Outlet />
    </RequireFinishedSignup>
  )
}
