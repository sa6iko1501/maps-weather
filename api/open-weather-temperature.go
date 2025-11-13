package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const WEATHER_URL string = "https://api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&units=metric&appid={API key}"

type WeatherRequest struct {
	Latitude   float64
	Longtitude float64
	Api        string
}

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
}

func GetWeatherUrl(lat string, lon string, apiKey string) string {
	var result string
	result = strings.Replace(WEATHER_URL, "{lat}", lat, 1)
	result = strings.Replace(result, "{lon}", lon, 1)
	result = strings.Replace(result, "{API key}", apiKey, 1)
	return result
}

func GetWeatherNow(r WeatherRequest) (result WeatherResponse, error error) {
	lat := r.Latitude
	lon := r.Longtitude
	url := GetWeatherUrl(strconv.FormatFloat(lat, 'f', -1, 64), strconv.FormatFloat(lon, 'f', -1, 64), r.Api)
	resp, error := http.Get(url)

	if error != nil {
		return result, error
	}

	status := resp.StatusCode
	if status != 200 {
		return WeatherResponse{}, fmt.Errorf("error wrong status: expected: 200, actual: %v", status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error casting to json response: %v\n", err)
		return result, err
	}

	defer resp.Body.Close()

	return result, err
}
