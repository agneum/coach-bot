package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	var token = os.Getenv("TG_TOKEN")
	if token == "" {
		log.Println("Token cannot be empty")
		return
	}

	cronJob := cron.New()
	err := cronJob.AddFunc("0 0 0 * * 1", schedulePosts)
	if err != nil {
		log.Fatalln(err)
	}
	err = cronJob.AddFunc("0 10 * * * *", checkScheduledPosts)
	if err != nil {
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

func initialize() {
	os.Remove("./data.db")

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open a DB connect: %s", err)
	}

	defer db.Close()

	creationQuery := `
CREATE TABLE autopost_plans(
id integer not null primary key AUTOINCREMENT,
chatid BIGINT NOT NULL,
type VARCHAR(16) NOT NULL,
texttemplate TEXT DEFAULT '',
lastscheduled TIMESTAMP,
intervals TEXT NOT NULL,
startdate TIMESTAMP NOT NULL,
enddate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
isactive integer NOT NULL DEFAULT 0
);

CREATE TABLE scheduled_posts (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	chatid BIGINT NOT NULL,
	senddate TIMESTAMP NOT NULL,
	message TEXT NOT NULL,
	done INTEGER NOT NULL DEFAULT 0
);
`
	_, err = db.Exec(creationQuery)
	if err != nil {
		log.Printf("%q: %s\n", err, creationQuery)
		return
	}
}

func schedulePosts() {
	fmt.Println("Schedule posts")
}

func checkScheduledPosts() {
	fmt.Println("Check scheduled posts")
}
