import { lazy } from "react"
import { createBrowserRouter } from "react-router-dom"
import RootLayout from "@/layouts/root-layout"

const Landing = lazy(() => import("@/pages/landing"))
const Signup = lazy(() => import("@/pages/signup"))
const Recipe = lazy(() => import("@/pages/recipe"))
const AllRecipes = lazy(() => import("@/pages/all-recipes"))
const Suggest = lazy(() => import("@/pages/suggest"))
const Cook = lazy(() => import("@/pages/cook"))
const Profile = lazy(() => import("@/pages/profile"))
const Login = lazy(() => import("@/pages/login"))

export const router = createBrowserRouter([
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
                element: <Profile/>,
            },
            {
                path: "/recipes",
                element: <AllRecipes/>,
            },
            {
                path: "/recipes/:id",
                element: <Recipe/>,
            },
            {
                path: "/suggest/:threadId",
                element: <Suggest/>,
            },
            {
                path: "/cook/:recipeId",
                element: <Cook/>,
            },
            {
                path: "/signup/*",
                element: <Signup/>,
            }
        ]
    }
])