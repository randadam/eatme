import { NotebookPen } from "lucide-react"
import { Link, useLocation } from "react-router-dom"
import { cn } from "@/lib/utils"

const items = [
  { to: "/plans", icon: NotebookPen, label: "Plans" },
]

const hiddenRoutes: RegExp[] = [
    new RegExp("^/signup"),
]

export default function BottomNav() {
  const { pathname } = useLocation()
  if (hiddenRoutes.some(route => route.test(pathname))) {
    return null
  }

  return (
    <nav className="fixed bottom-0 inset-x-0 bg-background border-t sm:hidden">
      <ul className="flex justify-around py-2">
        {items.map(({ to, icon: Icon, label }) => (
          <li key={to}>
            <Link
              to={to}
              className={cn(
                "flex flex-col items-center gap-1 text-muted-foreground",
                pathname === to && "text-primary"
              )}
            >
              <Icon className="h-6 w-6" />
              <span className="text-xs">{label}</span>
            </Link>
          </li>
        ))}
      </ul>
    </nav>
  )
}
