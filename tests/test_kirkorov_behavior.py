import pytest

from dapizda_bot.behavior.kirkorov import find_reply


@pytest.mark.parametrize(
    ("text", "expected"),
    [
        ("да", "пизда"),
        ("ну да", "пизда"),
        ("ДА!!!", "пизда"),
        ("дааа", "пизда"),
        ("пизда?", "да"),
        ("нет", "пидора ответ"),
        ("здрасьте.", "забор покрасьте"),
        ("300", "отсоси у тракториста"),
        ("триста", "отсоси у тракториста"),
        ("точно", "соси сочно"),
        ("шлюхи аргумент", "Аргумент не нужен, пидор обнаружен"),
        ("а", "хуй на"),
        ("где", "в пизде"),
        ("da", "pizda"),
        ("pizda", "da"),
        ("net", "pidora otvet"),
        ("gde", "v pizde"),
        ("lf", "gbplf"),
        ("gbplf", "lf"),
        ("ytn", "gbljhf jndtn"),
        ("ult", "d gbplt"),
    ],
)
def test_find_reply_matches_original_patterns(text: str, expected: str) -> None:
    assert find_reply(text) == expected


def test_find_reply_uses_random_reply_for_sosal() -> None:
    assert find_reply("сосал", choose=lambda replies: replies[1]) == (
        "Спрашиваешь так, будто сам хочешь мастер-класс провести"
    )


@pytest.mark.parametrize(
    "text",
    [
        "",
        "мимо",
        "давай",
        "нету",
        "пиздатый",
        "где-то",
    ],
)
def test_find_reply_ignores_non_matching_text(text: str) -> None:
    assert find_reply(text) is None

