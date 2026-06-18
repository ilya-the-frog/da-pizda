import os
from dataclasses import dataclass


@dataclass(frozen=True)
class Settings:
    bot_token: str
    log_level: str = "INFO"

    @classmethod
    def from_env(cls) -> "Settings":
        bot_token = os.environ.get("BOT_TOKEN")
        if not bot_token:
            raise RuntimeError("BOT_TOKEN is required. Set it in .env.")

        return cls(
            bot_token=bot_token,
            log_level=os.environ.get("LOG_LEVEL", "INFO").upper(),
        )

