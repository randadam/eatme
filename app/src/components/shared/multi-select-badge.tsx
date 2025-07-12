import { useController, type Control } from "react-hook-form"
import { Badge } from "@/components/ui/badge"
import { cn } from "@/lib/utils"
import { Check } from "lucide-react"

type MultiSelectBadgesProps = {
  name: string
  control: Control<any>
  options: { name: string, value: string }[]
}

export function MultiSelectBadges({ name, control, options }: MultiSelectBadgesProps) {
  const { field } = useController({
    name,
    control,
  })

  const selected = field.value || []

  const toggle = (val: string) => {
    const isSelected = selected.includes(val)
    const next = isSelected
      ? selected.filter((v: string) => v !== val)
      : [...selected, val]
    field.onChange(next)
  }

  return (
    <div role="group" className="flex flex-wrap gap-4">
      {options.map((option) => (
        <Badge
          key={option.value}
          onClick={() => toggle(option.value)}
          role="checkbox"
          aria-checked={selected.includes(option.value)}
          aria-label={option.name}
          data-testid={option.value}
          className={cn(
            "p-2 rounded-full cursor-pointer select-none transition-colors",
            selected.includes(option.value)
              ? "bg-primary text-primary-foreground"
              : "px-4 bg-muted text-muted-foreground"
          )}
        >
          {selected.includes(option.value) && (
            <Check className="h-4 w-4" />
          )}
          {option.name}
        </Badge>
      ))}
    </div>
  )
}
