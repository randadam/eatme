import type api from "@/api";
import BottomSheet from "@/components/shared/bottom-sheet";
import { ChatBody } from "@/features/chat/chat-body";
import FocusedLayout from "@/layouts/focused-layout";
import { useState } from "react";
import type { SheetSize } from "@/components/shared/bottom-sheet";

const recipe: api.ModelsUserRecipe = {
    id: "1",
    user_id: "1",
    thread_id: "1",
    title: "Test Recipe",
    description: "Test Recipe",
    ingredients: [],
    steps: [],
    servings: 1,
    total_time_minutes: 1,
    is_favorite: false,
    latest_version_id: "1",
    created_at: "2025-07-19T13:36:52.000Z",
    updated_at: "2025-07-19T13:36:52.000Z",
}

export default function ChatSandbox() {
    const [chatSize, setChatSize] = useState<SheetSize>("peek")

    return (
        <FocusedLayout>
            <BottomSheet
                header={<h2>Modify Recipe</h2>}
                subHeader={<p className="text-muted-foreground">Let me know what you'd like to change</p>}
                peekHeight={12}
                fullHeight={90}
                size={chatSize}
                onSizeChange={setChatSize}
            >
                <ChatBody
                    onSend={() => { }}
                    loading={false}
                    error={undefined}
                />
            </BottomSheet>
        </FocusedLayout>
    )
}