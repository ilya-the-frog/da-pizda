import random
import re
from collections.abc import Callable, Sequence

_TRAILING_PUNCT = re.compile(r"[!?.,…;:—–]+$")

_SOSAL_REPLIES = (
    "сам сосал",
    "Спрашиваешь так, будто сам хочешь мастер-класс провести",
    "Сосал. Чупа-чупс. Кислый, кстати, попался",
    "Конечно! Умею, практикую — особенно если это касается вкусного леденца",
    "Хочешь поговорить об этом?",
    "Что за вопрос? У тебя детская травма?",
    "Да, леденцы от боли в горле. Тебе тоже отсыпать?",
    "Только сок из этой жизни, пока ты тратишь мое время",
    "Конечно, прямо сейчас высасываю остатки твоего IQ",
    "Да, гранит науки. Тебе явно не помешало бы приобщиться",
    "Сосал. Палец. Пока думал, стоит ли вообще тебе отвечать",
)

PatternReply = str | tuple[str, ...]

RESPONSE_PATTERNS: tuple[tuple[re.Pattern[str], PatternReply], ...] = (
    # Cyrillic
    (re.compile(r"\bда+$", re.IGNORECASE), "пизда"),
    (re.compile(r"\bпи+зда+$", re.IGNORECASE), "да"),
    (re.compile(r"\bне+т+$", re.IGNORECASE), "пидора ответ"),
    (re.compile(r"\bздра+сьте+$", re.IGNORECASE), "забор покрасьте"),
    (re.compile(r"\b300$"), "отсоси у тракториста"),
    (re.compile(r"\bтри+ста+$", re.IGNORECASE), "отсоси у тракториста"),
    (re.compile(r"\bточно$", re.IGNORECASE), "соси сочно"),
    (re.compile(r"\bшлюхи аргумент$", re.IGNORECASE), "Аргумент не нужен, пидор обнаружен"),
    (re.compile(r"\bа$", re.IGNORECASE), "хуй на"),
    (re.compile(r"\bгде$", re.IGNORECASE), "в пизде"),
    (re.compile(r"\bсосал$", re.IGNORECASE), _SOSAL_REPLIES),
    # Latin phonetic translit
    (re.compile(r"\bda+$", re.IGNORECASE), "pizda"),
    (re.compile(r"\bpi+zda+$", re.IGNORECASE), "da"),
    (re.compile(r"\bne+t+$", re.IGNORECASE), "pidora otvet"),
    (re.compile(r"\bgde$", re.IGNORECASE), "v pizde"),
    # Russian keyboard layout translit
    (re.compile(r"\blf+$", re.IGNORECASE), "gbplf"),
    (re.compile(r"\bgb+plf+$", re.IGNORECASE), "lf"),
    (re.compile(r"\byt+n+$", re.IGNORECASE), "gbljhf jndtn"),
    (re.compile(r"\bult$", re.IGNORECASE), "d gbplt"),
)


def find_reply(
    text: str,
    choose: Callable[[Sequence[str]], str] = random.choice,
) -> str | None:
    normalized = _TRAILING_PUNCT.sub("", text.strip())
    for pattern, reply in RESPONSE_PATTERNS:
        if pattern.search(normalized):
            return choose(reply) if isinstance(reply, tuple) else reply

    return None
