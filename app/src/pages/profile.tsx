import BasicProfileSection from "@/features/auth/profile/basic-profile-section";
import { Accordion } from "@/components/ui/accordion";
import SkillSection from "@/features/auth/profile/skill-section";
import { useUser } from "@/features/auth/hooks";
import { useSaveProfile } from "@/features/auth/hooks";
import { Skeleton } from "@/components/ui/skeleton";
import CuisinesSection from "@/features/auth/profile/cuisines-section";
import DietSection from "@/features/auth/profile/diet-section";
import AllergiesSection from "@/features/auth/profile/allergies-section";
import EquipmentSection from "@/features/auth/profile/equipment-section";
import DefaultLayout from "@/layouts/default-layout";

export default function ProfilePage() {
    const { profile, isLoading } = useUser()
    const { saveProfile, isPending } = useSaveProfile()

    if (isLoading) {
        return (
            <DefaultLayout>
                <Skeleton className="h-16 w-full" />
                <Skeleton className="h-16 w-full" />
            </DefaultLayout>
        )
    }

    return (
        <DefaultLayout>
            {profile && (
                <Accordion type="single" collapsible className="w-full">
                    <BasicProfileSection profile={profile} onSave={saveProfile} isPending={isPending} />
                    <SkillSection profile={profile} onSave={saveProfile} isPending={isPending} />
                    <CuisinesSection profile={profile} onSave={saveProfile} isPending={isPending} />
                    <DietSection profile={profile} onSave={saveProfile} isPending={isPending} />
                    <AllergiesSection profile={profile} onSave={saveProfile} isPending={isPending} />
                    <EquipmentSection profile={profile} onSave={saveProfile} isPending={isPending} />
                </Accordion>
            )}
        </DefaultLayout>
    )
}