import { APP_NAME } from "@/constants";
import { useLocation, useNavigate } from "react-router-dom";
import { User } from "lucide-react";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger } from "@/components/ui/dropdown-menu"
import { useUser } from "@/features/auth/hooks";

export const BAR_HEIGHT = "h-14"

const pageTitles = [
    { pattern: new RegExp("^/recipes$"), title: "Recipe Book" },
    { pattern: new RegExp("^/recipes/.*"), title: "Recipe" },
    { pattern: new RegExp("^/profile"), title: "Profile" },
]

export default function AppBar() {
    const { isAuthenticated, logout } = useUser()
    const { pathname } = useLocation()
    const nav = useNavigate()
    const title = pageTitles.find(title => title.pattern.test(pathname))?.title

    const handleLogout = () => {
        logout()
        nav("/")
    }

    return (
        <>
            <header className={`fixed top-0 z-50 inset-x-0 w-screen ${BAR_HEIGHT} px-4 flex items-center justify-between border-b shadow-sm bg-background`}>
                <h1 className="text-xl font-semibold">{APP_NAME}</h1>
                {title && <h2 className="text-lg font-semibold">{title}</h2>}
                {isAuthenticated && (
                    <DropdownMenu>
                        <DropdownMenuTrigger>
                            <div className="w-8 h-8 rounded-full bg-muted flex items-center justify-center">
                                <User className="h-4 w-4" />
                            </div>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent>
                            <DropdownMenuItem onClick={() => nav("/profile")}>Profile</DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem onClick={handleLogout}>Logout</DropdownMenuItem>
                        </DropdownMenuContent>
                    </DropdownMenu>
                )}
            </header>
            <div className={`${BAR_HEIGHT}`} />
        </>
    )
}