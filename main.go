package main

import (
  "github.com/go-telegram-bot-api/telegram-bot-api"
  "log"
  "io/ioutil"
)

func main() {
  //достанем токен из файла
  token, err := ioutil.ReadFile("token.txt")
   if err != nil {
     log.Printf("File reading error", err)
     return
     }
  // подключаемся к боту с помощью токена
  bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		log.Panic(err)
	}

  bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			bot.Send(msg)
		}

	}
}
