// src/auth/RequireFinishedSignup.tsx
import { Navigate } from "react-router-dom";
import { useUser } from "@/features/auth/hooks"

interface Props {
  children: React.ReactNode
}

export default function RequireFinishedSignup({
  children,
}: Props) {
  const { isAuthenticated, profile, isLoading } = useUser();

  if (isLoading) return null;

  if (!isAuthenticated) return <Navigate to="/signup" replace />;

  if (profile?.setup_step !== "done") {
    return <Navigate to={`/signup/${profile?.setup_step}`} replace />;
  }

  return children;
}
