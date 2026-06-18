import pytest

from dapizda_bot.config import Settings


def test_settings_reads_bot_token(monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.setenv("BOT_TOKEN", "123:test")
    monkeypatch.setenv("LOG_LEVEL", "debug")

    settings = Settings.from_env()

    assert settings.bot_token == "123:test"
    assert settings.log_level == "DEBUG"


def test_settings_requires_bot_token(monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.delenv("BOT_TOKEN", raising=False)

    with pytest.raises(RuntimeError, match="BOT_TOKEN is required"):
        Settings.from_env()

