import { createPortal } from "react-dom";
import { useEffect, useState } from "react";
import { motion, useAnimate, useDragControls } from "framer-motion";
import { FocusTrap } from "focus-trap-react";
import { cn } from "@/lib/utils"; // shadcn helper for classnames

export type SheetSize = "peek" | "full";

interface BottomSheetProps {
    children: React.ReactNode;
    header?: React.ReactNode | ((size: SheetSize) => React.ReactNode);
    subHeader?: React.ReactNode | ((size: SheetSize) => React.ReactNode);
    actions?: React.ReactNode | ((size: SheetSize) => React.ReactNode);
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
    actions,
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
    const [scope, animate] = useAnimate();

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
                initialFocus: () => scope.current as HTMLElement,
                fallbackFocus: () => scope.current as HTMLElement,
                allowOutsideClick: true,
                setReturnFocus: false,
                tabbableOptions: { displayCheck: "none" },
                onDeactivate: () => { }, // noop
            }}
        >
            <motion.div
                role="dialog"
                aria-modal="false"
                ref={scope}
                className={cn(
                    "fixed inset-x-0 bottom-0 z-50 bg-background border-t flex flex-col",
                    "select-none touch-none", // prevents unwanted highlighting
                )}
                style={{
                    height: fullHeightWithPadding,
                    boxShadow: "0px -4px 24px rgba(0, 0, 0, 0.1)",
                }}
                variants={{
                    peek: { y: peekOffsetPx },
                    full: { y: bottomPadding },
                }}
                animate={size}
                transition={{ type: "spring", stiffness: 320, damping: 24 }}
                // drag to expand/collapse
                drag="y"
                dragListener={false}               // manual start
                dragControls={dragControls}
                dragConstraints={{ top: 2 * bottomPadding, bottom: 0 }}
                dragSnapToOrigin
                onDragEnd={(_, info) => {
                    console.log('info', info)
                    if (info.offset.y > 100 || info.velocity.y > 0) {
                        console.log("setting peek")
                        setSize("peek");
                        animate(scope.current, { y: peekOffsetPx }, { type: "spring", stiffness: 320, damping: 24 })
                    } else if (info.offset.y < -100 || info.velocity.y < 0) {
                        console.log("setting full")
                        setSize("full");
                        animate(scope.current, { y: bottomPadding }, { type: "spring", stiffness: 320, damping: 24 })
                    }
                }}
            >
                {/* grab handle */}
                <div>
                    <div
                        onPointerDown={(e) => dragControls.start(e)}
                        onClick={() => {
                            setSize(size === "peek" ? "full" : "peek")
                            animate(scope.current, { y: size === "peek" ? bottomPadding : peekOffsetPx }, { type: "spring", stiffness: 320, damping: 24 })
                        }}
                        className="flex items-center justify-center pt-4 pb-1 cursor-pointer"
                    >
                        <div className="h-1.5 w-10 rounded-full bg-muted-foreground/40" />
                    </div>

                    <div
                        onPointerDown={(e) => dragControls.start(e)}
                        className={`h-[calc(${peekHeight}vh-4rem)] py-2`}
                    >
                        {header && (
                            <header className="flex items-center justify-center px-4 w-full">
                                {typeof header === "function" ? header(size) : header}
                            </header>
                        )}
                        {subHeader && (
                            <div className="flex items-center justify-center px-4 w-full">
                                {typeof subHeader === "function" ? subHeader(size) : subHeader}
                            </div>
                        )}
                    </div>
                    {actions && (
                        <div className="flex items-center justify-center px-4 w-full">
                            {typeof actions === "function" ? actions(size) : actions}
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
