import asyncio
import importlib
from dataclasses import dataclass

import pytest

pytest.importorskip("aiogram")

kirkorov = importlib.import_module("dapizda_bot.handlers.kirkorov")


@dataclass
class FakeSentMessage:
    text: str
    message_id: int = 2


class FakeMessage:
    message_id = 1

    def __init__(self) -> None:
        self.replies: list[str] = []

    async def reply(self, text: str) -> FakeSentMessage:
        self.replies.append(text)
        return FakeSentMessage(text=text)


def test_disable_ai_is_reply_only_stub() -> None:
    message = FakeMessage()

    asyncio.run(kirkorov.disable_ai(message))

    assert message.replies == ["AI отключён для этого чата."]


def test_enable_ai_is_reply_only_stub() -> None:
    message = FakeMessage()

    asyncio.run(kirkorov.enable_ai(message))

    assert message.replies == ["AI включён для этого чата."]


def test_pattern_response_replies_and_registers_noop(monkeypatch: pytest.MonkeyPatch) -> None:
    message = FakeMessage()

    async def no_sleep(_seconds: float) -> None:
        return None

    monkeypatch.setattr(kirkorov.asyncio, "sleep", no_sleep)

    asyncio.run(kirkorov.pattern_response(message, "пизда"))

    assert message.replies == ["пизда"]
