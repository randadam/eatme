import { lazy } from "react"
import { createBrowserRouter } from "react-router-dom"
import RootLayout from "@/layouts/root-layout"
import ProtectedRoute from "@/components/shared/protected-route"

const Landing = lazy(() => import("@/pages/landing"))
const Signup = lazy(() => import("@/pages/signup"))
const Recipe = lazy(() => import("@/pages/recipe"))
const AllRecipes = lazy(() => import("@/pages/all-recipes"))
const Suggest = lazy(() => import("@/pages/suggest"))
const Cook = lazy(() => import("@/pages/cook"))
const Profile = lazy(() => import("@/pages/profile"))
const Login = lazy(() => import("@/pages/login"))
const ModifyRecipe = lazy(() => import("@/pages/modify-recipe"))
const ChatSandbox = lazy(() => import("@/sandbox/chat-sandbox"))

const routes = [
    {
        element: <RootLayout/>,
        children: [
            {
                path: "/",
                element: <Landing/>,
            },
            {
                path: "/login",
                element: <Login/>,
            },
            {
                path: "/profile",
                element: (
                    <ProtectedRoute>
                        <Profile/>
                    </ProtectedRoute>
                ),
            },
            {
                path: "/recipes",
                element: (
                    <ProtectedRoute>
                        <AllRecipes/>
                    </ProtectedRoute>
                ),
            },
            {
                path: "/recipes/:id",
                element: (
                    <ProtectedRoute>
                        <Recipe/>
                    </ProtectedRoute>
                ),
            },
            {
                path: "/suggest/:threadId",
                element: (
                    <ProtectedRoute>
                        <Suggest/>
                    </ProtectedRoute>
                ),
            },
            {
                path: "/cook/:recipeId",
                element: (
                    <ProtectedRoute>
                        <Cook/>
                    </ProtectedRoute>
                ),
            },
            {
                path: "/modify-recipe/:recipeId",
                element: (
                    <ProtectedRoute>
                        <ModifyRecipe/>
                    </ProtectedRoute>
                ),
            },
            {
                path: "/signup/*",
                element: <Signup/>,
            }
        ]
    }
]

if (import.meta.env.DEV) {
    routes[0].children?.push({
        path: "/sandbox/chat",
        element: <ChatSandbox/>,
    })
}

export const router = createBrowserRouter(routes)