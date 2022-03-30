package main

import (
	"fmt"
	"log"
	"os"

	w_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userkeyboard = w_bot.NewReplyKeyboard(
	w_bot.NewKeyboardButtonRow(
		w_bot.NewKeyboardButtonLocation("Your Geo"),
	),
)

func main() {
	bot, err := w_bot.NewBotAPI(os.Getenv("FUNSTORM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := w_bot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if update.Message.IsCommand() { // ignore any non-command Messages
			HandleCommand(bot, update.Message)
		}
		if update.Message.Location != nil {
			HandleLocation(bot, update.Message)
		}
	}
}
func HandleCommand(bot *w_bot.BotAPI, message *w_bot.Message) {
	// Create a new MessageConfig. We don't have text yet,
	// so we leave it empty.
	reply := w_bot.NewMessage(message.Chat.ID, "")

	// Extract the command from the Message.
	switch message.Command() {
	case "weather":
		reply.Text = "Home boy please give your location"
		reply.ReplyMarkup = userkeyboard
	default:
		reply.Text = "I don't know that command"
	}

	if _, err := bot.Send(reply); err != nil {
		log.Panic(err)
	}
}
func HandleLocation(bot *w_bot.BotAPI, message *w_bot.Message) {
	reply := w_bot.NewMessage(message.Chat.ID, "")

	reply.Text = fmt.Sprintf("Latitude: %f Longitude: %f", message.Location.Latitude, message.Location.Longitude)
	reply.ReplyMarkup = w_bot.NewRemoveKeyboard(true)
	if _, err := bot.Send(reply); err != nil {
		log.Panic(err)
	}
}
