import asyncio
import logging

from aiogram import Router, types
from aiogram.filters import BaseFilter, Command

from dapizda_bot.behavior.kirkorov import find_reply

logger = logging.getLogger(__name__)


class HasPatternReply(BaseFilter):
    async def __call__(self, message: types.Message) -> bool | dict[str, str]:
        if reply := find_reply(message.text or ""):
            return {"pattern_reply": reply}

        return False


async def register_pattern_exchange(
    message: types.Message,
    sent: types.Message,
    pattern_reply: str,
) -> None:
    logger.debug(
        "Pattern exchange registration skipped",
        extra={
            "incoming_message_id": message.message_id,
            "sent_message_id": sent.message_id,
            "pattern_reply": pattern_reply,
        },
    )


async def pattern_response(message: types.Message, pattern_reply: str) -> None:
    await asyncio.sleep(1)
    sent = await message.reply(pattern_reply)
    await register_pattern_exchange(message, sent, pattern_reply)


async def edited_pattern_response(message: types.Message, pattern_reply: str) -> None:
    await pattern_response(message, pattern_reply)


async def disable_ai(message: types.Message) -> None:
    await message.reply("AI отключён для этого чата.")


async def enable_ai(message: types.Message) -> None:
    await message.reply("AI включён для этого чата.")


def make_router() -> Router:
    router = Router(name=__name__)
    router.message(Command("disable_ai"))(disable_ai)
    router.message(Command("enable_ai"))(enable_ai)
    router.message(HasPatternReply())(pattern_response)
    router.edited_message(HasPatternReply())(edited_pattern_response)
    return router

