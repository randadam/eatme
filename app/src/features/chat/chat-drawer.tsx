import type api from "@/api";
import LoaderButton from "@/components/shared/loader-button";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import { Textarea } from "@/components/ui/textarea";
import { useState } from "react";

interface ChatDrawerProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  mode: "suggest" | "modify" | "question";
  recipe?: api.ModelsUserRecipe;
  onSend: (message: string) => void;
  loading: boolean;
  error?: string;
}

export function ChatDrawer({
  open,
  onOpenChange,
  mode,
  recipe,
  onSend,
  loading,
  error,
}: ChatDrawerProps) {
  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent side="bottom" className="flex flex-col w-full sm:w-[380px] h-[50vh] p-4">
        <SheetHeader>
          <SheetTitle>
            {mode === "suggest" && "Suggest a Recipe"}
            {mode === "modify" && `Modify: ${recipe?.title}`}
            {mode === "question" && "Ask a Question"}
          </SheetTitle>
        </SheetHeader>

        <ChatInput
          onSend={onSend}
          loading={loading}
        />
        {error && <p>{error}</p>}
      </SheetContent>
    </Sheet>
  );
}

interface ChatInputProps {
    onSend: (message: string) => void;
    loading: boolean;
}

function ChatInput({ onSend, loading }: ChatInputProps) {
    const [inputValue, setInputValue] = useState("")

    return (
        <div className="flex flex-col space-y-2">
            <Textarea value={inputValue} onChange={(e) => setInputValue(e.target.value)} />
            <LoaderButton onClick={() => onSend(inputValue)} isLoading={loading}>
                Send
            </LoaderButton>
        </div>
    )
}