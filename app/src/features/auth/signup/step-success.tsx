import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";

export default function SignupSuccess() {
  const nav = useNavigate()

  return (
    <div className="flex flex-col items-center justify-center space-y-4 pt-16">
      <h2 className="text-2xl font-semibold">Welcome aboard! 🎉</h2>
      <p className="text-muted-foreground pt-4">
        Well done! Now we're ready to get cooking!
      </p>
      <Button onClick={() => nav("/")} className="mt-4">Let's Get Started!</Button>
    </div>
  )
}
