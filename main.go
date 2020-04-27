package main

import (
  "github.com/go-telegram-bot-api/telegram-bot-api"
  "log"
  "io/ioutil"
  "time"
  "fmt"
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
      if update.ChannelPost == nil {
        fmt.Println(update)
        // Пользователь, который написал боту
        UserName := update.Message.From.UserName
        // ID чата/диалога.
        // Может быть идентификатором как чата с пользователем
        // (тогда он равен UserID) так и публичного чата/канала
			  ChatID := update.Message.Chat.ID
			  // Текст сообщения
			  Text := update.Message.Text
			  log.Printf("[%s] %d %s", UserName, ChatID, Text)
			  // Ответим пользователю его же отредаченным сообщением
			  reply := Text + "\n \n" + "Автор поста: @" + UserName
			  // Созадаем сообщение
			  msg := tgbotapi.NewMessage(ChatID, reply)
			  // и отправляем его
        //time.Sleep(time.Millisecond * 5000000)
			  bot.Send(msg)
    }
		}
	}
} 
