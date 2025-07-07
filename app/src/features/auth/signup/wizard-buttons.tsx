import { useLocation, Link } from "react-router-dom"
import { Button } from "@/components/ui/button"
import { STEPS } from "./constants"
import { Loader2 } from "lucide-react"

interface Props {
  submitLabel?: string
  loading?: boolean
}

export default function WizardButtons({ submitLabel = "Next", loading }: Props) {
  const { pathname } = useLocation()
  const index = Object.values(STEPS).indexOf(pathname)
  const prev = index > 1 ? Object.values(STEPS)[index - 1] : undefined

  return (
    <div className="flex justify-between gap-4 pt-6">
      {prev ? (
        <Button variant="secondary" asChild>
          <Link to={prev}>Back</Link>
        </Button>
      ) : (
        <span />
      )}

      <Button type="submit" disabled={loading}>
        {loading && (
            <Loader2 className="mr-2 h-4 w-4 animate-spin" />
        )}
        {submitLabel}
      </Button>
    </div>
  )
}
