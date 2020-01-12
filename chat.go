package main

import (
	"log"
	"os"

	"github.com/themoah/go-web-tools/core"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// InitTGBot start to listen for telegram messages
func InitTGBot() {

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

		userText := update.Message.Text
		go handleMsg(update, bot, userText)
	}

}

func handleMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI, userText string) {

	chatID := update.Message.Chat.ID
	log.Printf("[%v] %v on %v", update.Message.From.UserName, userText, chatID)

	if userText == "usage" {
		msg := tgbotapi.NewMessage(chatID, printUsage())
		bot.Send(msg)
	} else if userText == "foo" {
		msg := tgbotapi.NewMessage(chatID, core.Foo())
		bot.Send(msg)
	} else if userText[0:2] == "ip" {
		msg := tgbotapi.NewMessage(chatID, core.WhoIS(userText[2:]))
		bot.Send(msg)
	} else if userText == "random" {
		msg := tgbotapi.NewMessage(chatID, core.Random())
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(chatID, userText)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}

func printUsage() string {

	usage := `
	foo 		returns bar
	ip 1.2.3.4  returns WhoIS for ipv4
	random		returns random float between 0 and 1
	`

	return usage
}
