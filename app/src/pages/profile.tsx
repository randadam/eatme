import BasicProfileSection from "@/features/auth/profile/basic-profile-section";
import { Accordion } from "@/components/ui/accordion";
import SkillsSection from "@/features/auth/profile/skills-section";
import { useUser } from "@/features/auth/hooks";
import { useSaveProfile } from "@/features/auth/hooks";
import { Skeleton } from "@/components/ui/skeleton";

export default function ProfilePage() {
    const { profile, isLoading } = useUser()
    const { saveProfile, isPending } = useSaveProfile()

    if (isLoading) {
        return (
            <div>
                <Skeleton className="h-16 w-full" />
                <Skeleton className="h-16 w-full" />
            </div>
        )
    }

    return (
        <div>
            <h1>Profile</h1>
            {profile && (
                <Accordion type="single" className="w-full">
                    <BasicProfileSection profile={profile} onSave={saveProfile} isPending={isPending} />
                    <SkillsSection profile={profile} onSave={saveProfile} isPending={isPending} />
                </Accordion>
            )}
        </div>
    )
}