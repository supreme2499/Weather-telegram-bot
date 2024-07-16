package internal

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	msgStart = "Привет это простенький бот с одной командой, с его помощью ты можешь узнать погоду.\nЧтобы воспользоваться ботом, отправь ему /weather + город, и он скажет солько сейчас градусов на улице"
	msgHelp  = "Неизвестная команда, попробуй ещё раз. Отправь мне /weather + город погода которого интересует"
)

// команда /start
func HandleStartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, msgStart)
	bot.Send(msg)
}

// команда /weather
func HandleWeatherCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	parts := strings.Fields(message.Text)
	if len(parts) < 2 {
		// проверка на то что город указан
		msg := tgbotapi.NewMessage(message.Chat.ID, "Пожалуйста, укажите город после команды /weather.")
		bot.Send(msg)
		return
	}
	town := parts[1]

	lat, lon, country, result := Decod(town)
	if result {
		temp, desc := Weather(lat, lon)
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%s:\nПогода: %s\nТемпература: %s", country, desc, temp))
		bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Город не найден, попробуй ещё раз:(")
		bot.Send(msg)
	}

}

// команда /help или неизвестная комманда
func HandleHelpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(message.Chat.ID, msgHelp)
	bot.Send(msg)
}
