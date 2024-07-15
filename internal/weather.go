package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"telegram-bot/pkg/config"
)

//service: https://openweathermap.org/api

// пример запроса https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid={API key}ы
// lat - широта; lon - долгота

type Temperature struct {
	Main struct {
		Temp float32 `json:"temp"`
	} `json:"main"`
}

func Weather(lat, lon string) string {
	cfg := config.GetConfig()
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, cfg.WeatKey)
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

	temp := temperature.Main.Temp

	return fmt.Sprintf("%.2f℃", temp)
}
