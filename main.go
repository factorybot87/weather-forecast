package main

import (
	"fmt"
	"forecast/weather"
	"net/http"
	"time"
)

const (
	key      = ""
	location = "taipei"
	layout   = "January 2, 2006"
)

func main() {
	forecast, _ := getForecast(location, key)
	formatOutput(forecast)
}

func getForecast(location, key string) (weather.Forecast, error) {
	client := http.Client{}
	url := weather.GenerateURL(location, key)
	req := weather.BuildRequest(url)

	res, _ := client.Do(req)
	// TODO: Check reponse status code

	jsonData := weather.GetBody(res)
	forecast, err := weather.ParseResponseBody(jsonData)

	return forecast, err
}

func formatOutput(forecast weather.Forecast) {
	current := forecast.CurrentConditions
	currentTime := time.Unix(current.DatetimeEpoch, 0).Format(layout)
	currentWeatherFormat := `
		Current Time: %s
		Summary:      %s
		Temperature:  %.2f
		Wind Speed:   %.2f
	`
	fmt.Printf(currentWeatherFormat, currentTime, current.Conditions, current.Temp, current.WindSpeed)
}
