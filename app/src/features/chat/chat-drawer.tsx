import type api from "@/api";
import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import { Textarea } from "@/components/ui/textarea";
import { Loader2 } from "lucide-react";
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
      <SheetContent side="right" className="flex flex-col w-full sm:w-[380px]">
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
            <Button onClick={() => onSend(inputValue)} disabled={loading}>
                {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                Send
            </Button>
        </div>
    )
}