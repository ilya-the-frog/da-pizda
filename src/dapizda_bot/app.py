import asyncio
import logging

from aiogram import Bot, Dispatcher

from dapizda_bot.config import Settings
from dapizda_bot.handlers import get_routers


async def run_bot(settings: Settings | None = None) -> None:
    settings = settings or Settings.from_env()

    logging.basicConfig(
        level=settings.log_level,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )

    bot = Bot(token=settings.bot_token)
    dispatcher = Dispatcher()
    for router in get_routers():
        dispatcher.include_router(router)

    await dispatcher.start_polling(bot)


def main() -> None:
    asyncio.run(run_bot())

