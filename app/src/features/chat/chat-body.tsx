import LoaderButton from "@/components/shared/loader-button";
import { Skeleton } from "@/components/ui/skeleton";
import { Textarea } from "@/components/ui/textarea";
import { useEffect, useRef, useState } from "react";

export interface ChatItem {
  source: 'user' | 'assistant';
  message: string;
}

interface ChatBodyProps {
  history?: ChatItem[];
  onSend: (message: string) => void;
  loading: boolean;
  error?: string;
}

export function ChatBody({
  history,
  onSend,
  loading,
  error,
}: ChatBodyProps) {

  const [input, setInput] = useState("")
  const messageEndRef = useRef<HTMLDivElement | null>(null)

  useEffect(() => {
    messageEndRef.current?.scrollIntoView({ behavior: "smooth" })
    setInput("")
  }, [history])

  const handleSend = (message: string) => {
    onSend(message)
  }

  return (
    <div className="flex flex-col h-full px-1 pb-1 gap-2">
      {history && (
        <div className="flex flex-col flex-1 space-y-4 overflow-y-scroll border border-gray-200 rounded-md p-4">
          {history?.map((item, idx) => (
            item.source === "user" ? (
              <div key={idx} className="flex flex-col gap-2 bg-primary text-primary-foreground p-2 rounded-md mb-2 ml-4">
                <p className="text-right">{item.message}</p>
                <p className="text-right text-xs">Me</p>
              </div>
            ) : (
              <div key={idx} className="flex flex-col gap-2 bg-gray-200 p-2 rounded-md mb-2 mr-4">
                <p className="text-left">{item.message}</p>
                <p className="text-right text-xs">Assistant</p>
              </div>
            )
          ))}
          {loading && (
            <div className="flex flex-col gap-2 bg-gray-200 p-2 rounded-md mb-2 mr-4">
              <Skeleton className="h-8 w-full" />
            </div>
          )}
          <div ref={messageEndRef} />
        </div>
      )}
      <div className="flex flex-col space-y-2">
        <ChatInput
          onSend={handleSend}
          loading={loading}
          input={input}
          setInput={setInput}
          disabled={loading}
        />
        {error && <p>{error}</p>}
      </div>
    </div>
  );
}

interface ChatInputProps {
  onSend: (message: string) => void;
  loading: boolean;
  input: string;
  setInput: (input: string) => void;
  disabled?: boolean;
}

function ChatInput({ onSend, loading, input, setInput, disabled }: ChatInputProps) {

  return (
    <div className="flex flex-col space-y-2">
      <Textarea
        className="min-h-[15vh]"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        disabled={disabled}
        autoFocus={true}
      />
      <LoaderButton
        onClick={() => onSend(input)}
        isLoading={loading}
        disabled={input.trim() === ""}
      >
        Send
      </LoaderButton>
    </div>
  )
}