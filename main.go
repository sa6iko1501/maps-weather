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
		var geoReq api.GeoLocRequest = api.GeoLocRequest{City: city, Api: apiKey}
		var geoResp api.GeoLocResponse
		geoResp, err := api.GetGeoLoc(geoReq)
		if err != nil {
			fmt.Println(err)
		} else {
			if (geoResp == api.GeoLocResponse{}) {
				fmt.Printf("no city with the name : %s was found", city)
				return
			}
			var weatherReq api.WeatherRequest = api.WeatherRequest{Latitude: geoResp.Latitude, Longtitude: geoResp.Longtitude, Api: apiKey}
			weatherResp, err := api.GetWeatherNow(weatherReq)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("COUTRY: %v\nState: %v\nCITY: %v\nOVERALL_WEATHER: %v\nTEMPERATURE: %v C\nFEELS_LIKE: %v C\n", geoResp.Country, geoResp.State, city, weatherResp.Weather[0].Description, weatherResp.Main.Temp, weatherResp.Main.FeelsLike)
			}
		}
	} else {

		fmt.Println("It's allways sunny in Philadelphia!")
	}

}
