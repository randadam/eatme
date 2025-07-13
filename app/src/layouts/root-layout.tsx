import { Outlet } from "react-router-dom"
import BottomNav from "@/components/shared/bottom-nav"
import RequireFinishedSignup from "@/features/auth/require-finished-signup"
import AppBar from "@/components/shared/app-bar"
import { Suspense } from "react"
import { Loader2 } from "lucide-react"

export default function RootLayout() {
  return (
    <RequireFinishedSignup>
      <AppBar />

      <main className="flex-1 container mx-auto p-4 pb-24">
        <Suspense fallback={
          <Loader2 className="w-6 h-6 animate-spin" />
        }>
          <Outlet />
        </Suspense>
      </main>

      <BottomNav />
    </RequireFinishedSignup>
  )
}
