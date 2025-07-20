export interface FocusedLayoutProps {
    children: React.ReactNode
}

export default function FocusedLayout({ children }: FocusedLayoutProps) {
    return (
        <main
            className="flex-1 container mx-auto min-h-screen"
            style={{
                paddingBottom: "calc(var(--bs-peek, 0) + env(safe-area-inset-bottom))"
            }}
        >
            {children}
        </main>
    )
}