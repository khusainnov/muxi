package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"test/internal/weather"

	"github.com/gorilla/mux"
)



type weatherLocal struct {
	we weather.Weather
}

func (s *Service) GetWeather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]

	API := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("WEATHER_API_TOKEN"), city)

	resp, err := http.Get(API)
	if err != nil {
		log.Fatalln(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	weather1 := &weatherLocal{}

	err = json.Unmarshal(body, weather1)
	if err != nil {
		log.Fatalln(err.Error())
	}

	weatherLocation := weather1.we.Location

	fmt.Fprintf(w, "<b>Country: %v\n<br>City: %v\n<br>Airport: %v\n<br>Local Time: %v\n<br>Temperature: %v,C / %v,F\n</b>",
		weatherLocation.Country, weatherLocation.Region, weatherLocation.Name, weatherLocation.Localtime, weather1.we.Current.TempC, weather1.we.Current.TempF)
	log.Printf("\nCountry: %v\nCity: %v\nAirport: %v\nLocal Time: %v\nTemperature: %v,C / %v,F\n",
		weatherLocation.Country, weatherLocation.Region, weatherLocation.Name, weatherLocation.Localtime, weather1.we.Current.TempC, weather1.we.Current.TempF)

	return
}
