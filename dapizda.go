package main

import (
	"io/ioutil"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

var answers = map[string]string{
	"да":       "пизда",
	"da":       "пизда",
	"lf":       "пизда",
	"пизда":    "да",
	"pizda":    "да",
	"gbplf":    "да",
	"нет":      "пидора ответ",
	"net":      "пидора ответ",
	"ytn":      "пидора ответ",
	"здрасьте": "забор покрасьте",
	"300":      "отсоси у тракториста",
	"триста":   "отсоси у тракториста",
}

func main() {
	//достанем токен из файла
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.WithError(err).Error("File reading error")
		return
	}

	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		log.WithError(err).Fatal("Error connecting to the bot")
	}

	bot.Debug = false
	log.WithField("account", bot.Self.UserName).Print("Authorized on account")

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.WithError(err).Fatal("Error getting updates channel")
	}
	time.Sleep(time.Millisecond * 500)
	upd.Clear()
	// читаем обновления из канала
	for {
		if update, ok := <-upd; ok {
			//проверяем, от канала или от пользователя
			if update.ChannelPost == nil && update.EditedMessage == nil {
				if reply, ok := answers[strings.ToLower(update.Message.Text)]; ok {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
					msg.BaseChat.ReplyToMessageID = update.Message.MessageID //добавляем реплай
					log.WithField("reply", reply).Print("Sending reply")
					_, err := bot.Send(msg)
					if err != nil {
						log.WithError(err).Fatal("Error sending message")
					}
				}
			}
		}
	}
}
