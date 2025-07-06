async def answer(message: str) -> str:
    messages = [
        sys("You are a helpful cooking assistant. Answer concisely but clearly."),
        usr(f"User: \"{message}\"")
    ]
    return await as_json(messages, schema=str)
