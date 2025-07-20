import AppBar from "@/components/shared/app-bar"
import { Suspense } from "react"
import { Loader2 } from "lucide-react"
import BottomNav from "@/components/shared/bottom-nav"

interface DefaultLayoutProps {
    children: React.ReactNode
}

export default function DefaultLayout({ children }: DefaultLayoutProps) {
    return (
        <>
            <AppBar />

            <main className="flex-1 container mx-auto p-4 pb-24 min-h-screen">
                <Suspense fallback={
                    <Loader2 className="w-6 h-6 animate-spin" />
                }>
                    {children}
                </Suspense>
            </main>

            <BottomNav />
        </>
    )
}