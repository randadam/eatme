import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-react";
import { useEffect, useRef, useState, type ComponentPropsWithoutRef } from "react";

export interface LoaderButtonProps extends ComponentPropsWithoutRef<typeof Button> {
    isLoading: boolean
    delay?: number
    minDuration?: number
}

export default function LoaderButton({ children, isLoading, delay = 150, minDuration = 400, ...rest }: LoaderButtonProps) {
    const [visible, setVisible] = useState(false)
    const startRef = useRef<number | null>(null)

    useEffect(() => {
        let t: ReturnType<typeof setTimeout>
    
        if (isLoading) {
            t = setTimeout(() => {
                startRef.current = Date.now()
                setVisible(true)
            }, delay)
        } else if (visible) {
            const elapsed = Date.now() - (startRef.current ?? 0)
            const remaining = Math.max(0, minDuration - elapsed)
            t = setTimeout(() => setVisible(false), remaining)
        }

        return () => clearTimeout(t)
    }, [isLoading, delay, minDuration, visible])

    return (
        <Button disabled={isLoading || visible} {...rest}>
            {visible && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
            {children}
        </Button>
    )
}