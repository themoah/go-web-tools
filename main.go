package main

import (
	"log"
	"net/http"
	"os"

	"github.com/themoah/go-web-tools/routes"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func initTGBot() {

	botToken := os.Getenv("TG_TOKEN")
	if botToken == "" {
		log.Panic("token undefined")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	// bot.Debug = true

	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			log.Println("bummber, got empty message")
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.Text == "foo" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "bar")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}

func main() {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
		log.Printf("Defaulting to port %s", serverPort)
	}

	log.Println("starting server, listening on port 0.0.0.0:" + serverPort)

	initTGBot()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.IndexHandler)
	r.HandleFunc("/healthz", routes.HealthCheckHanlder).Methods("GET")
	r.HandleFunc("/echo", routes.EchoHandler).Methods("GET")
	r.HandleFunc("/foo", routes.FooHandler).Methods("GET")
	r.HandleFunc("/random", routes.RandomHandler).Methods("GET")
	r.HandleFunc("/secure", routes.SecureHandler).Methods("GET").Schemes("https")
	r.HandleFunc("/whois/{ip}", routes.WhoisHandler).Methods("GET")

	log.Fatal(http.ListenAndServe("0.0.0.0:"+serverPort, r))

}
