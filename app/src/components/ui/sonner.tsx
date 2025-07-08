import { useTheme } from "next-themes"
import { Toaster as Sonner, type ToasterProps } from "sonner"

const Toaster = ({ ...props }: ToasterProps) => {
  const { theme = "system" } = useTheme()

  return (
    <Sonner
      theme={theme as ToasterProps["theme"]}
      className="toaster group"
      toastOptions={{
        classNames: {
          toast: "text-left bg-popover border-popover text-popover-foreground",
          success: "text-left !bg-green-500 !border-green-500 !text-green-foreground",
          error: "text-left !bg-destructive !border-destructive !text-destructive-foreground",
          warning: "text-left !bg-yellow-500 !border-yellow-500 !text-yellow-foreground",
          info: "text-left !bg-blue-500 !border-blue-500 !text-blue-foreground",
        }
      }}
      {...props}
    />
  )
}

export { Toaster }
