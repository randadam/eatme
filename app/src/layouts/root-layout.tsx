import { Outlet } from "react-router-dom"
import BottomNav from "@/components/shared/bottom-nav"

export default function RootLayout() {
  return (
    <div className="min-h-screen flex flex-col">
      <header className="hidden sm:flex h-14 items-center px-4 border-b">
        <h1 className="font-semibold text-lg">My Recipes</h1>
      </header>

      <main className="flex-1 container mx-auto p-4 pb-24">
        <Outlet />
      </main>

      <BottomNav />
    </div>
  )
}
