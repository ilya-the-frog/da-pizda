# Да? Пизда!

Telegram bot на aiogram.

Подписывайтесь на канал автора бота: https://t.me/travelasproduct

## Оживление бота

Спасибо ребятам, которые вернули бота к жизни: зачистили старую реализацию,
перенесли поведение на aiogram, добавили Docker Compose, Taskfile и тесты.

## Локальный запуск

1. Скопируйте `.env.example` в `.env`.
2. Укажите токен Telegram-бота в `BOT_TOKEN`.
3. Запустите:

```bash
task run
```

## Проверки

```bash
task test
task lint
```

## Расширение

Новые aiogram-роутеры добавляются в `src/dapizda_bot/handlers/` и подключаются в
`src/dapizda_bot/handlers/__init__.py`.
