// src/auth/RequireFinishedSignup.tsx
import { Navigate, useLocation } from "react-router-dom";
import { useUser } from "@/features/auth/hooks"

interface Props {
  children: React.ReactNode
}

export default function RequireFinishedSignup({
  children,
}: Props) {
  const { isAuthenticated } = useUser();
  const { pathname } = useLocation()

  if (pathname === "/" || pathname.includes("/login")) return children;

  if (!isAuthenticated && !pathname.startsWith("/signup")) return <Navigate to="/signup" replace />;

  return children;
}
