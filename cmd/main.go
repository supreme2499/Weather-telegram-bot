package main

import (
	"log"
	"strings"
	"telegram-bot/internal"
	"telegram-bot/pkg/config"
	"telegram-bot/pkg/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	cfg := config.GetConfig()
	logger := logging.GetLogger()
	//подключение к клиенту телеграм
	bot, err := tgbotapi.NewBotAPI(cfg.TgKey)
	if err != nil {
		log.Fatal("ошибка получения бота: ", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { //проверка на получение сообщения
			logger.Infof("User:%s", update.Message.Chat.UserName)
			logger.Infof("Message:%s", update.Message.Text)
			if strings.HasPrefix(update.Message.Text, "/start") {
				internal.HandleStartCommand(bot, update.Message)
			} else if strings.HasPrefix(update.Message.Text, "/weather") {
				internal.HandleWeatherCommand(bot, update.Message)
			} else {
				internal.HandleHelpCommand(bot, update.Message)
			}
		}
	}
}
