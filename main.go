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
	currentTime := formatDatetime(current.DatetimeEpoch)
	currentWeatherFormat := `
Current Time: %s
Summary:      %s
Temperature:  %.2f
Wind Speed:   %.2f
`
	fmt.Printf(currentWeatherFormat, currentTime, current.Conditions, current.Temp, current.WindSpeed)

	next5DaysForecast := forecast.Days[1:6]

	for _, day := range next5DaysForecast {
		forecastTime := formatDatetime(day.DatetimeEpoch)
		forecastFormat := `
Date:         %s
Summary:      %s
Description:  %s
Temperature:  %.2f
Temp_max:     %.2f
Temp_min:     %.2f
Wind Speed:   %.2f
`
		fmt.Printf(forecastFormat, forecastTime, day.Conditions, day.Description, day.Temp, day.TempMax, day.TempMin, day.WindSpeed)
	}
}

func formatDatetime(datetime int64) string {
	return time.Unix(datetime, 0).Format(layout)
}
