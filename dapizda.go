package main

import (
	"io/ioutil"
	"log"
	"strings"
	"time"
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
	"точно":    "соси сочно",
}

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
		log.Fatalf("Error connecting to the bot: %v", err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.Fatalf("Error getting updates channel: %v", err)
	}
	time.Sleep(time.Millisecond * 500)
	upd.Clear()
	// читаем обновления из канала
	for {
		select {
		case update := <-upd:
			//проверяем, от канала или от пользователя
			if update.ChannelPost == nil && update.EditedMessage == nil {
				if reply := answers[strings.ToLower(update.Message.Text)]; reply != "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
					msg.BaseChat.ReplyToMessageID = update.Message.MessageID //добавляем реплай
					log.Printf("Sending %s", reply)
					_, err := bot.Send(msg)
					if err != nil {
						log.Fatalf("Error sending message: %v", err)
					}
				}
			}
		}
	}
}
