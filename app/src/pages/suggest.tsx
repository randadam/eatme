import { useParams } from "react-router-dom"
import { useGetSuggestionThread } from "@/features/chat/hooks"
import SuggestionThread from "@/features/chat/suggest/suggestion-thread"

export default function Suggest() {
    const threadId = useParams().threadId
    if (!threadId) {
        return <div>Missing thread ID</div>
    }

    const { thread: initialThread, fetchLoading, fetchError } = useGetSuggestionThread(threadId)

    if (fetchLoading) {
        return <div>Loading...</div>
    }
    if (fetchError) {
        return <div>Error: {fetchError.message}</div>
    }
    if (!initialThread) {
        return <div>Thread not found</div>
    }

    return <SuggestionThread initialThread={initialThread} />
}
