package config

import (
	"telegram-bot/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	YDecod  string `yaml:"y_decoder_key"`
	WeatKey string `yaml:"weather_key"`
	TgKey   string `yaml:"telegram_key"`
}

var instance *Config

func GetConfig() *Config {
	logger := logging.GetLogger()
	instance = &Config{}
	if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		logger.Info("ошибка чтения конфига: ", help)
		logger.Fatal(err)
	}
	return instance
}
