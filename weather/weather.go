package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Main struct {
	Temp float64 `json:"temp"`
}
type Weather struct {
	Main string `json:"main"`
}

type CurrentWeatherResponse struct {
	Weather []Weather `json:"weather"`
	Main    `json:"main"`
}

const (
	// url is in the format https://api.openweathermap.org/data/2.5/weather?lat=HEADER&lon=HEADER&appid=HEADER then you can add &units=metric to get the temperature in celsius
	url           = "ADD YOUR URL HERE"
	file_to_write = "/tmp/weather.txt"
)

func main() {
	var cwr CurrentWeatherResponse
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		jsonErr := json.Unmarshal(body, &cwr)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

	}
	current_temp := strconv.FormatFloat(cwr.Temp, 'f', 0, 64)
	current_weather := cwr.Weather[0].Main
	weather_data := current_temp + "C " + current_weather
	// write weather_data to file_to_write
	file, err := os.Create(file_to_write)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(weather_data)

}
