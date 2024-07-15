package main

import (
	"fmt"
	"telegram-bot/internal"
)

//const mytoken = "7319163777:AAHEm4eC21EwOr5Zjuq72m844rgkGrYrZBM"

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
