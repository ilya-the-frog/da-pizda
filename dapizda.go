package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"time"
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

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	time.Sleep(time.Millisecond * 500)
	upd.Clear()
	// читаем обновления из канала
	for {
		select {
		case update := <-upd:
			//проверяем, от канала или от пользователя
			if update.ChannelPost == nil && update.EditedMessage == nil {
				var reply = "" // чекаю текст

				if update.Message.Text == "да" || update.Message.Text == "Да" {
					reply = "пизда"
				}

				if update.Message.Text == "пизда" || update.Message.Text == "Пизда" {
					reply = "да"
				}

				if update.Message.Text == "нет" || update.Message.Text == "Нет" {
					reply = "пидора ответ"
				}

				if update.Message.Text == "здрасьте" || update.Message.Text == "Здрасьте" {
					reply = "забор покрасьте"
				}

				if update.Message.Text == "пидора ответ" {
					reply = "сам пидора ответ"
				}

				if reply != "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
					msg.BaseChat.ReplyToMessageID = update.Message.MessageID //добавляем реплай
					log.Printf("Send %s", reply)
					if _, err := bot.Send(msg); err != nil {
						panic(err)
					}
				}
			}
		}
	}
}
