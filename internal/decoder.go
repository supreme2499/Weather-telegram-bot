package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"telegram-bot/pkg/config"
)

//вот пример запроса https://geocode-maps.yandex.ru/1.x/?apikey={key}&geocode=москваformat=json
//мне нужен Points{pos}

type Response struct {
	Response struct {
		GeoObjectCollection struct {
			FeatureMember []struct {
				GeoObject struct {
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

func Decod(Town string) (lat, lon string) {
	cfg := *config.GetConfig()
	nTown := url.QueryEscape(Town)
	url := fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?apikey=%s&geocode=%s&format=json", cfg.YDecod, nTown)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("ошибка выполнения запроса декодированияЖ:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ошибка чтения ответа: ", err)
	}

	var response Response

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Fatal("ошибка парсинга: ", err)
	}

	coordinaties := response.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos

	nCord := strings.Split(coordinaties, " ")
	lat = nCord[1]
	lon = nCord[0]

	return lat[:5], lon[:5]
}
