package main

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	var token = os.Getenv("TG_TOKEN")
	if token == "" {
		log.Println("Token cannot be empty")
		return
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func notify(bot *tgbotapi.BotAPI) {
	var chatId = os.Getenv("CHAT_ID")
	if chatId == "" {
		log.Println("CHAT_ID cannot be empty")
		return
	}

	id, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		log.Println("CHAT_ID cannot be empty")
		return
	}

	msg := tgbotapi.NewMessage(id, "Hello")
	bot.Send(msg)
}
