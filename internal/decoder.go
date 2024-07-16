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

type Response struct {
	Response struct {
		GeoObjectCollection struct {
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							AddressDetails struct {
								Country struct {
									AddressLine string `json:"AddressLine"`
								} `json:"Country"`
							} `json:"AddressDetails"`
						} `json:"GeocoderMetaData"`
					} `json:"metaDataProperty"`
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

func Decod(Town string) (lat, lon, country string, result bool) {
	cfg := *config.GetConfig()
	nTown := url.QueryEscape(Town)
	url := fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?apikey=%s&geocode=%s&format=json&results=1", cfg.YDecod, nTown)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("ошибка выполнения запроса декодирования:", err)
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

	if len(response.Response.GeoObjectCollection.FeatureMember) > 0 {
		result = true
		country = response.Response.GeoObjectCollection.FeatureMember[0].GeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AddressLine
		coordinaties := response.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos

		nCord := strings.Split(coordinaties, " ")
		lat = nCord[1]
		lon = nCord[0]

		return lat[:5], lon[:5], country, result
	} else {
		return "", "", "", result
	}
}
