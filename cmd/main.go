package main

import (
	"fmt"
	"telegram-bot/internal"
)

func main() {
	//logger := logging.GetLogger()
	//подключение к клиенту телеграм

	//получение обновлений

	//вычисление координат по названию города Орловская+область+малоархангельск
	//logger.Info("преобразование города в координаты")
	lat, lon := internal.Decod("Орловская+область+малоархангельск")
	fmt.Printf("Широта: %s Долгота: %s\n", lat, lon)

	//получение темпералуры
	//logger.Info("получение температуры")
	fmt.Printf("Температура: %s\n", internal.Weather(lat, lon))
}
