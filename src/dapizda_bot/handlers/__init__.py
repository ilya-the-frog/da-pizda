from aiogram import Router

from dapizda_bot.handlers.kirkorov import make_router as make_kirkorov_router


def get_routers() -> list[Router]:
    return [
        make_kirkorov_router(),
    ]

