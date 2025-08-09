package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

var answers = map[string]string{
    `(?i)(?:^|[^\wа-я])да+[^\wа-я]*$`:          "пизда",
    `(?i)(?:^|[^\wа-я])da+[^\wа-я]*$`:          "пизда",
    `(?i)(?:^|[^\wа-я])lf+[^\wа-я]*$`:          "пизда",
    `(?i)(?:^|[^\wа-я])пи+зда+[^\wа-я]*$`:      "да",
    `(?i)(?:^|[^\wа-я])pi+zda+[^\wа-я]*$`:      "да",
    `(?i)(?:^|[^\wа-я])gb+plf+[^\wа-я]*$`:      "да",
    `(?i)(?:^|[^\wа-я])не+т+[^\wа-я]*$`:        "пидора ответ",
    `(?i)(?:^|[^\wа-я])ne+t+[^\wа-я]*$`:        "пидора ответ",
    `(?i)(?:^|[^\wа-я])yt+n+[^\wа-я]*$`:        "пидора ответ",
    `(?i)(?:^|[^\wа-я])здра+сьте+[^\wа-я]*$`:   "забор покрасьте",
    `(?:^|[^\wа-я])300[^\wа-я]*$`:              "отсоси у тракториста",
    `(?i)(?:^|[^\wа-я])три+ста+[^\wа-я]*$`:     "отсоси у тракториста"}

var compiledAnswers map[*regexp.Regexp]string

func init() {
	compiledAnswers = make(map[*regexp.Regexp]string)
	for pattern, reply := range answers {
		re := regexp.MustCompile(pattern)
		compiledAnswers[re] = reply
	}
}

func checkInput(input string) string {
	for re, reply := range compiledAnswers {
		if re.MatchString(input) {
			return reply
		}
	}
	return ""
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
                if reply := checkInput(update.Message.Text); reply != "" {
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
