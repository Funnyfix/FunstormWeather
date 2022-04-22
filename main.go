package main

import (
	"database/sql"
	"funstorm/owmhelper"
	"log"
	"os"
	"strings"
	"time"

	w_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type subscription struct {
	id     int
	chatId int64
	time   time.Time
	city   string
}

var userkeyboard = w_bot.NewReplyKeyboard(
	w_bot.NewKeyboardButtonRow(
		w_bot.NewKeyboardButtonLocation("Your Geo"),
	),
)
var bot *w_bot.BotAPI

func main() {
	IntializeBot()

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
			HandleCommand(update.Message)
			continue
		}
		if update.Message.Location != nil {
			HandleLocation(update.Message)
			continue
		}
		if update.Message.Text != "" {
			HandlePlace(update.Message)
			continue
		}
	}
}

func IntializeBot() {
	token := os.Getenv("FUNSTORM_BOT_TOKEN")
	client, err := w_bot.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot = client
}

func HandleCommand(message *w_bot.Message) {
	// Create a new MessageConfig. We don't have text yet,
	// so we leave it empty.
	reply := w_bot.NewMessage(message.Chat.ID, "")
	reply.ParseMode = "HTML"
	helptext := "<code>/geo</code> - Current weather in your location\n" +
		"<code>/city Your city</code> - Current weather in selected city\n" +
		"<code>/subscribe HH:mm Your city</code> - Your daily weather subscribtion\n" +
		"You can also type any place without command"

	// Extract the command from the Message.
	switch message.Command() {
	case "start":
		reply.Text = "Hello, i am dumb ass bot for now please type my functions:\n" + helptext

	case "geo":
		reply.Text = "Home boy please give your location"
		reply.ReplyMarkup = userkeyboard
	case "help":
		reply.Text = helptext
	case "city":
		HandlePlace(message)
		//	case "subscribe":
		//		HandleForecast(message)

	default:
		reply.Text = "I don't know that command"
	}
	if reply.Text != "" {
		if _, err := bot.Send(reply); err != nil {
			log.Panic(err)
		}
	}
}

func HandleLocation(message *w_bot.Message) {

	current_weather := owmhelper.CurrentWeatherByCoordinates(message.Location.Latitude, message.Location.Longitude)
	text := owmhelper.ParseWeather(current_weather)
	Answer(message.Chat.ID, text)

}

func HandlePlace(message *w_bot.Message) {
	var parsed_text = strings.TrimPrefix(message.Text, "/city")
	parsed_text = strings.TrimPrefix(parsed_text, " ")
	log.Println(parsed_text)
	if len(parsed_text) == 0 {
		text := "Enter place \nExample: <code>/city Воронеж</code>"
		Answer(message.Chat.ID, text)
		return
	}
	current_weather := owmhelper.CurrentWeatherByName(parsed_text)
	text := owmhelper.ParseWeather(current_weather)
	Answer(message.Chat.ID, text)

}

func HandleSubscibe(message *w_bot.Message) {
	var parsed_text = strings.TrimPrefix(message.Text, "/subscribe")
	parsed_text = strings.TrimPrefix(parsed_text, " ")
	log.Println(parsed_text)
	if len(parsed_text) < 7 {
		text := "Enter your time and city \nExample: <code>/subscribe 10:00 Воронеж</code>"
		Answer(message.Chat.ID, text)
		return
	}

	parsedTime, parsedPlace, err := ParseSubscribe(parsed_text)
	if err != nil {
		text := "Enter your time and city \nExample: <code>/subscribe 10:00 Воронеж</code>"
		Answer(message.Chat.ID, text)
		return
	}
	subscribeData := subscription{chatId: message.Chat.ID, time: parsedTime, city: parsedPlace}
	err2 := sqlAddSubscription(subscribeData)
	if err2 != nil {
		text := err2.Error()
		Answer(message.Chat.ID, text)
		return
	}
}

func ParseSubscribe(text string) (time.Time, string, error) {
	parsed_time := strings.TrimSuffix(text[:5], " ")
	parsed_city := strings.TrimPrefix(text[5:], " ")
	const layout = "15:04"
	subtime, err := time.Parse(layout, parsed_time)
	if err != nil {
		return *new(time.Time), "", err
	}
	return subtime, parsed_city, nil
}

func Answer(chatid int64, text string) {
	reply := w_bot.NewMessage(chatid, "")
	reply.Text = text
	reply.ReplyMarkup = w_bot.NewRemoveKeyboard(true)
	reply.ParseMode = "HTML"
	if _, err := bot.Send(reply); err != nil {
		log.Panic(err)
	}
}

func sqlAddSubscription(data subscription) error {
	db, err := sql.Open("sqlite3", "subscribe.db")
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec("insert into subscriptions (chatId, time, city) values (?,?,?)", data.chatId, data.time, data.city)
	db.Close()
	return err
}
