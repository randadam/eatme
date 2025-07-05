import type { ReactNode } from "react"

export default function StepInstructions({ children }: { children: ReactNode }) {
  return (
    <div className="space-y-2 pb-6">
      <h2 className="text-xl font-semibold">Sign Up</h2>
      <p className="text-muted-foreground">{children}</p>
    </div>
  )
}
