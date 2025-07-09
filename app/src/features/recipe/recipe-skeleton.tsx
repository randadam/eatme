import { Skeleton } from "@/components/ui/skeleton"

export default function RecipeSkeleton() {
    return (
        <div className="space-y-2">
            <Skeleton className="h-12 w-full" />
            <Skeleton className="h-32 w-full" />
            <div className="flex justify-between space-x-2">
                <Skeleton className="h-12 w-1/2" />
                <Skeleton className="h-12 w-1/2" />
            </div>
        </div>
    )
}