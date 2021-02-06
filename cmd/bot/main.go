package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/sqlite3"
	tgbotapi "gopkg.in/telegram-bot-api.v4"

	"github.com/agneum/scheduler-bot/internal/scheduler"
	"github.com/agneum/scheduler-bot/pkg/storage"
)

var (
//port = flag.Int("port", 8080, "a default server port")
)

type Config struct {
	Telegram Telegram `yaml:"telegram"`
}

type Telegram struct {
	Token string `yaml:"token" env:"SB_TG_TOKEN"`
}

func main() {
	//flag.Parse()

	conn, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db := reform.NewDB(conn, sqlite3.Dialect, nil)
	templateSvc := storage.NewTemplateRepo(db.Querier)
	eventSvc := storage.NewEventRepo(db.Querier)

	schedulerSvc := scheduler.NewScheduler(templateSvc, eventSvc)
	//if err := schedulerSvc.Schedule(); err != nil {
	//	log.Fatal(err)
	//}
	//return

	var cfg Config

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		Files: []string{"config.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
		EnvPrefix: "SB",
	})

	if err := loader.Load(); err != nil {
		log.Fatalf("fail to load config: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatalf("fail to init Telegram API: %v", err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "/register":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "TBD: User has been registered")
			bot.Send(msg)

		case "/templates":
			templates := schedulerSvc.ShowTemplates()
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, templates))
			//fmt.Println(templates)

		case "/schedule":
			events := schedulerSvc.ShowEvents()
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, events))
			//fmt.Println(events)
		}
	}
}

func cronScheduler() {
	//cronJob := cron.New()
	//
	//if err := cronJob.AddFunc("0 0 0 * * 1", schedulePosts); err != nil {
	//	log.Fatalln(err)
	//}
	//
	//if err := cronJob.AddFunc("0 10 * * * *", checkScheduledPosts); err != nil {
	//	log.Fatalln(err)
	//}
	//cronJob.Start()
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
