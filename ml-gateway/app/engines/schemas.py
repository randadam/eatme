import json
from typing import Any, Type
from pydantic import BaseModel, TypeAdapter

def schema_snippet(tp: Type[Any]) -> str:
    """
    Return a compact JSON Schema string for a Pydantic model OR typing-hint
    (e.g. list[Recipe]).
    """
    if isinstance(tp, type) and issubclass(tp, BaseModel):
        sch = tp.model_json_schema()
    else:
        sch = TypeAdapter(tp).json_schema()
    # keep prompt short: drop titles/descriptions
    sch.pop("title", None)
    sch.pop("description", None)
    return json.dumps(sch, ensure_ascii=False, indent=2)
