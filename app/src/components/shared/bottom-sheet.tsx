import { createPortal } from "react-dom";
import { useEffect, useRef, useState } from "react";
import { motion, useDragControls } from "framer-motion";
import { FocusTrap } from "focus-trap-react";
import { cn } from "@/lib/utils"; // shadcn helper for classnames

export type SheetSize = "peek" | "full";

interface BottomSheetProps {
    children: React.ReactNode;
    header?: React.ReactNode;
    subHeader?: React.ReactNode;
    initialSize?: SheetSize;
    peekHeight?: number;      // in vh
    fullHeight?: number;      // in vh
    size?: SheetSize;         // controlled mode (optional)
    onSizeChange?: (s: SheetSize) => void;
}

export default function BottomSheet({
    children,
    header,
    subHeader,
    initialSize = "peek",
    peekHeight = 26,
    fullHeight = 90,
    size: controlledSize,
    onSizeChange,
}: BottomSheetProps) {
    const [internalSize, setInternalSize] = useState<SheetSize>(initialSize);
    const size = controlledSize ?? internalSize;          // uncontrolled by default
    const bottomPadding = 50;
    const vh = window.innerHeight / 100;
    const fullHeightWithPadding = (fullHeight * vh) + bottomPadding;
    const peekOffsetPx = (fullHeight - peekHeight) * vh + bottomPadding;

    const setSize = (s: SheetSize) => {
        controlledSize ? onSizeChange?.(s) : setInternalSize(s);
    };

    const dragControls = useDragControls();
    const sheetRef = useRef<HTMLDivElement>(null);

    // collapse on Esc when full
    useEffect(() => {
        const onKey = (e: KeyboardEvent) => {
            if (e.key === "Escape" && size === "full") {
                setSize("peek");
            }
        };
        window.addEventListener("keydown", onKey);
        return () => window.removeEventListener("keydown", onKey);
    }, [size]);

    // manage main content size to avoid peek mode blocking content
    useEffect(() => {
        // Apply on mount; clean up on unmount
        const root = document.documentElement
        root.style.setProperty("--bs-peek", `${peekHeight}vh`)
        // optional: let iOS bottom safe area add itself
        return () => {
            root.style.removeProperty("--bs-peek")
        }
    }, [peekHeight])

    // portal target
    const root = document.body;

    return createPortal(
        <FocusTrap
            focusTrapOptions={{
                // don't trap when in peek mode
                escapeDeactivates: false,
                clickOutsideDeactivates: false,
                initialFocus: () => sheetRef.current as HTMLElement,
                fallbackFocus: () => sheetRef.current as HTMLElement,
                allowOutsideClick: true,
                setReturnFocus: false,
                tabbableOptions: { displayCheck: "none" },
                onDeactivate: () => { }, // noop
            }}
        >
            <motion.div
                role="dialog"
                aria-modal="false"
                ref={sheetRef}
                className={cn(
                    "fixed inset-x-0 bottom-0 z-50 bg-background border-t shadow-lg flex flex-col",
                    "select-none touch-none", // prevents unwanted highlighting
                )}
                style={{
                    height: fullHeightWithPadding,
                }}
                variants={{
                    peek: {
                        y: peekOffsetPx,
                    },
                    full: {
                        y: bottomPadding,
                    },
                }}
                animate={size}
                // layout animation between heights
                layout="size"
                transition={{ type: "spring", stiffness: 320, damping: 24 }}
                // drag to expand/collapse
                drag="y"
                dragListener={false}               // manual start
                dragControls={dragControls}
                dragConstraints={{ top: 0, bottom: 0 }}
                onDragEnd={(_, info) => {
                    if (info.velocity.y > 300 || info.offset.y > 120) {
                        setSize("peek");
                    } else if (info.velocity.y < -300 || info.offset.y < -120) {
                        setSize("full");
                    }
                }}
            >
                {/* grab handle */}
                <div
                    onPointerDown={(e) => dragControls.start(e)}
                    onClick={() => setSize(size === "peek" ? "full" : "peek")}
                    className="flex items-center justify-center pt-4 pb-1 cursor-pointer"
                >
                    <div className="h-1.5 w-10 rounded-full bg-muted-foreground/40" />
                </div>

                <div className={`h-[calc(${peekHeight}vh-4rem)] py-2`}>
                    {header && (
                        <header className="flex items-center justify-center px-4 w-full">
                            {header}
                        </header>
                    )}
                    {subHeader && (
                        <div className="flex items-center justify-center px-4 w-full">
                            {subHeader}
                        </div>
                    )}
                </div>

                {size === 'full' && (
                    <div
                        className="flex-1 overflow-y-auto"
                        style={{
                            paddingBottom: `calc(${bottomPadding}px + env(safe-area-inset-bottom))`
                        }}
                    >
                        {children}
                    </div>
                )}
                <div className={`h-[${bottomPadding}px]`} />
            </motion.div>
        </FocusTrap>,
        root,
    );
}
