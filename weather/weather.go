package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GenerateURL will Construct the Visual Crossing API url from components
func GenerateURL(location, key string) string {
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"
	params := "unitGroup=metric&include=days%2Ccurrent"
	contentType := "contentType=json"
	url := fmt.Sprintf("%s/%s?%s&key=%s&%s", baseURL, location, params, key, contentType)
	return url
}

// BuildRequest will build a request with the headers
func BuildRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

// GetBody will take an HTTP response and extract the body as a string
func GetBody(res *http.Response) string {
	body, _ := io.ReadAll(res.Body)
	return string(body)
}

// ParseResponseBody will parse an API reponse into a Forecast
func ParseResponseBody(jsonData string) (Forecast, error) {
	forecast := Forecast{}
	err := json.Unmarshal([]byte(jsonData), &forecast)
	return forecast, err
}

// Forecast has only the parts we care about, currentConditions and Days
type Forecast struct {
	CurrentConditions CurrentCondition
	Days              []DailyWeather
}

// CurrentCondition represent the current weather
type CurrentCondition struct {
	DatetimeEpoch int64
	Temp          float32
	WindSpeed     float32
	Conditions    string
}

// DailyWeather represent the daily forecast for the next 15 days
type DailyWeather struct {
	DatetimeEpoch int64
	Temp          float32
	TempMin       float32
	TempMax       float32
	WindSpeed     float32
	Conditions    string
	Description   string
}
