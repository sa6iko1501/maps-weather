package main

import (
	"flag"
	"fmt"
	"sa6iko1501/maps-weather/api"
)

func main() {
	var city string
	var apiKey string
	flag.StringVar(&city, "city", "Default", "Geographical Data")
	flag.StringVar(&apiKey, "api_key", "Default", "API_KEY for OpenWeather")
	flag.Parse()

	if city != "Default" && apiKey != "Default" {
		var req api.Request = api.Request{Latitude: 43.2, Longtitude: 23.2, Api: apiKey}
		var response api.Response
		response, err := api.GetWeatherNow(req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("CITY: %v\nOVERALL_WEATHER: %v\nTEMPERATURE: %v C\nFEELS_LIKE: %v C\n", city, response.Weather[0].Description, response.Main.Temp, response.Main.FeelsLike)
		}
	} else {

		fmt.Println("It's allways sunny in Philadelphia!")
	}

}
