"""
Anthropic-powered LLM wrapper for the recipe ML gateway.
API surface stays identical to the former OpenAI version.
"""

from __future__ import annotations
import asyncio, json, logging, random, os
from typing import Sequence, TypeVar, Type
from engines.schemas import schema_snippet
from anthropic import Anthropic, \
    RateLimitError, InternalServerError, APIStatusError
from pydantic import BaseModel, ValidationError, TypeAdapter

logger = logging.getLogger("llm")
client = Anthropic(
    # picks up ANTHROPIC_API_KEY from env by default
    api_key=os.getenv("ANTHROPIC_API_KEY"),
)

T = TypeVar("T", bound=BaseModel)

MODEL = "claude-3-5-haiku-latest"

# ------------------------------------------------------------------ #
# 1. Message helpers: system/user/assistant -> Anthropic format      #
# ------------------------------------------------------------------ #
def sys(content: str) -> dict:
    return {"role": "system", "content": content}

def usr(content: str) -> dict:
    return {"role": "user", "content": content}

def asst(content: str) -> dict:
    return {"role": "assistant", "content": content}

def _to_anthropic(messages: Sequence[dict]) -> str:
    """
    Convert our list-of-dicts format ➜ Anthropic Claude prompt text.
    Follows the <role>:::<content>\n\n pattern.
    """
    parts: list[str] = []
    for m in messages:
        role = m["role"]
        content = m["content"]
        if role == "system":
            # Claude likes a single system preamble; merge later
            parts.insert(0, f"{content.strip()}\n\n")
        else:
            parts.append(f"{role.capitalize()}: {content.strip()}\n\n")
    # Anthropic requires final "Assistant:" cue
    return "".join(parts) + "Assistant:"

# ------------------------------------------------------------------ #
# 2. Core chat with automatic retries                                #
# ------------------------------------------------------------------ #
async def chat(
    messages: Sequence[dict],
    *,
    model: str = MODEL,
    max_tokens: int | None = None,
    temperature: float = 0.0,
    retries: int = 3,
    backoff_base: float = 0.4,
) -> str:
    attempt = 0
    prompt = _to_anthropic(messages)

    while True:
        try:
            resp = client.messages.create(
                model=model,
                max_tokens=max_tokens or 4096,
                temperature=temperature,
                messages=[{"role": "user", "content": prompt}],
            )
            return resp.content[0].text  # Anthropic returns list[MessageBlock]
        except (RateLimitError, InternalServerError, APIStatusError) as err:
            attempt += 1
            if attempt > retries:
                logger.error("Anthropic call failed: %s", err)
                raise
            sleep = backoff_base * (2 ** (attempt - 1)) * random.uniform(0.8, 1.3)
            logger.warning("Anthropic error (%s). Retrying in %.2fs…", err.__class__.__name__, sleep)
            await asyncio.sleep(sleep)

# ------------------------------------------------------------------ #
# 3. JSON helper with Pydantic validation                            #
# ------------------------------------------------------------------ #
async def as_json(
    messages: Sequence[dict],
    schema: Type[T],
    *,
    model: str = MODEL,
    temperature: float = 0.0,
    **chat_kwargs,
) -> T:

    schema_text = schema_snippet(schema)
    print('schema_text', schema_text)

    system_hint = sys(
        "You are a strict JSON generator. "
        "Output MUST be valid JSON conforming to this schema, "
        "and must NOT be wrapped in markdown fences.\n\n"
        f"```json\n{schema_text}\n```"
    )

    print('system_hint', system_hint)

    raw = await chat(
        messages + [system_hint],
        model=model,
        temperature=temperature,
        **chat_kwargs,
    )

    print('raw', raw)

    try:
        data = json.loads(raw)
        validator = (
            schema.model_validate if isinstance(schema, type) and issubclass(schema, BaseModel)
            else TypeAdapter(schema).validate_python
        )
        return validator(data)  # type: ignore[arg-type]
    except (json.JSONDecodeError, ValidationError) as e:
        # One self-repair attempt
        repair = [
            asst(raw),
            usr("The JSON is invalid or doesn’t match the schema. "
                "Fix it and return ONLY corrected JSON.")
        ]
        raw = await chat(messages + [system_hint] + repair,
                         model=model, temperature=0.0)
        data = json.loads(raw)
        print('repair data', data)
        return validator(data)  # may raise, that’s fine

# ------------------------------------------------------------------ #
# 4. __all__ so other modules can `from engines.llm import …`        #
# ------------------------------------------------------------------ #
__all__ = ["chat", "as_json", "sys", "usr", "asst"]
