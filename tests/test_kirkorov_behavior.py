import pytest

from dapizda_bot.behavior.kirkorov import find_reply


@pytest.mark.parametrize(
    ("text", "expected"),
    [
        ("да", "пизда"),
        ("ну да", "пизда"),
        ("ДА!!!", "пизда"),
        ("дааа", "пизда"),
        ("д" + "a", "пизда"),
        ("д а", "пизда"),
        ("пизда?", "да"),
        ("пииздаа?", "да"),
        ("нет", "пидора ответ"),
        ("н" + "e" + "т", "пидора ответ"),
        ("нееетт", "пидора ответ"),
        ("здрасьте.", "забор покрасьте"),
        ("здрааасьтее.", "забор покрасьте"),
        ("300", "отсоси у тракториста"),
        ("триста", "отсоси у тракториста"),
        ("триистааа", "отсоси у тракториста"),
        ("точно", "соси сочно"),
        ("привет", "минет"),
        ("утро", "хуютро"),
        ("шлюхи аргумент", "Аргумент не нужен, пидор обнаружен"),
        ("а", "хуй на"),
        ("где", "в пизде"),
        ("da", "pizda"),
        ("daa", "pizda"),
        ("pizda", "da"),
        ("piiizdaaa", "da"),
        ("net", "pidora otvet"),
        ("neettt", "pidora otvet"),
        ("gde", "v pizde"),
        ("lf", "gbplf"),
        ("lff", "gbplf"),
        ("gbplf", "lf"),
        ("gbbplff", "lf"),
        ("ytn", "gbljhf jndtn"),
        ("yttnn", "gbljhf jndtn"),
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
