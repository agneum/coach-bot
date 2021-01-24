package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func main() {
	token := os.Getenv("TG_TOKEN")
	if token == "" {
		log.Println("Token cannot be empty")
		return
	}

	cronJob := cron.New()

	if err := cronJob.AddFunc("0 0 0 * * 1", schedulePosts); err != nil {
		log.Fatalln(err)
	}

	if err := cronJob.AddFunc("0 10 * * * *", checkScheduledPosts); err != nil {
		log.Fatalln(err)
	}
	cronJob.Start()

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
	chatID := os.Getenv("CHAT_ID")
	if chatID == "" {
		log.Println("CHAT_ID cannot be empty")
		return
	}

	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		log.Println("CHAT_ID cannot be empty")
		return
	}

	msg := tgbotapi.NewMessage(id, "Hello")
	message, err := bot.Send(msg)
	if err != nil {
		log.Println("failed to send message")
		return
	}

	bot.PinChatMessage(tgbotapi.PinChatMessageConfig{
		ChatID:              id,
		DisableNotification: true,
		MessageID:           message.MessageID,
	})
}

func schedulePosts() {
	fmt.Println("Schedule posts")
}

func checkScheduledPosts() {
	fmt.Println("Check scheduled posts")
}
