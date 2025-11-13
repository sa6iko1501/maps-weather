package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const GEOLOC_URL string = "http://api.openweathermap.org/geo/1.0/direct?q={city name}&limit=1&appid={API key}"

type GeoLocRequest struct {
	City string
	Api  string
}

type GeoLocResponse struct {
	Latitude   float64 `json:"lat"`
	Longtitude float64 `json:"lon"`
	Country    string  `json:"country"`
	State      string  `json:"state"`
}

func GetGeoLocUrl(city string, apiKey string) string {
	var result string
	result = strings.Replace(GEOLOC_URL, "{city name}", city, 1)
	result = strings.Replace(result, "{API key}", apiKey, 1)
	return result
}

func GetGeoLoc(r GeoLocRequest) (result GeoLocResponse, error error) {
	city, api := r.City, r.Api
	url := GetGeoLocUrl(city, api)

	resp, err := http.Get(url)

	if err != nil {
		return GeoLocResponse{}, err
	}

	status := resp.StatusCode
	if status != 200 {
		return GeoLocResponse{}, fmt.Errorf("error wrong status: expected: 200, actual: %v", status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return result, err
	}

	var results []GeoLocResponse
	err = json.Unmarshal(body, &results)
	if err != nil {
		return result, err
	}

	if len(results) > 0 {
		result = results[0]
	}

	return result, error
}
