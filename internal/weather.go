package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"telegram-bot/pkg/config"
)

type Temperature struct {
	Weather []struct {
		Desc string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
}

// lat - долгота lon - широта
func Weather(lat, lon string) (string, string) {
	cfg := config.GetConfig()
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric&lang=ru", lat, lon, cfg.WeatKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("ошибка выполннения запроса:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ошибка чтения ответа: ", err)
	}

	var temperature Temperature

	err = json.Unmarshal([]byte(body), &temperature)
	if err != nil {
		log.Fatal("ошибка парсинга: ", err)
	}
	desc := temperature.Weather[0].Desc
	temp := temperature.Main.Temp
	return fmt.Sprintf("%.2f℃", temp), desc
}
