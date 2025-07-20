import { Notebook, User } from "lucide-react"
import { Link, useLocation } from "react-router-dom"
import { cn } from "@/lib/utils"

const items = [
  { to: "/recipes", icon: Notebook, label: "Recipes", testId: 'nav-recipes' },
  { to: "/profile", icon: User, label: "Profile", testId: 'nav-profile' },
]

export default function BottomNav() {
  const { pathname } = useLocation()

  return (
    <nav className="fixed bottom-0 inset-x-0 bg-background border-t sm:hidden">
      <ul className="flex justify-around py-2">
        {items.map(({ to, icon: Icon, label, testId }) => (
          <li key={to}>
            <Link
              to={to}
              data-testid={testId}
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
