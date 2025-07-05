import { Link, useLocation } from "react-router-dom"
import { STEPS } from "./routes"
import { Button } from "@/components/ui/button"

export default function WizardNav() {
  const { pathname } = useLocation()
  const index = STEPS.indexOf(pathname as (typeof STEPS)[number])

  // If we're outside the wizard (e.g. success page) render nothing
  if (index === -1) return null

  const prev = STEPS[index - 1]
  const next = STEPS[index + 1]

  return (
    <div className="flex justify-between gap-4 pt-6">
      {prev ? (
        <Button variant="secondary" asChild>
          <Link to={prev}>Prev</Link>
        </Button>
      ) : <span />}

      {next ? (
        <Button asChild>
          <Link to={next}>Next</Link>
        </Button>
      ) : (
        <Button asChild>
          <Link to="/signup/success">Finish</Link>
        </Button>
      )}
    </div>
  )
}
