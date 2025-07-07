import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";

export default function SignupSuccess() {
  const nav = useNavigate()

  return (
    <div className="flex flex-col items-center justify-center h-full space-y-4">
      <h2 className="text-2xl font-semibold">Welcome aboard! 🎉</h2>
      <p className="text-muted-foreground">
        You can start browsing recipes right away.
      </p>
      <Button onClick={() => nav("/")}>Let's Get Started!</Button>
    </div>
  )
}
