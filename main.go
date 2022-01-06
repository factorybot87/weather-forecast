package main

import (
	"forecast/weather"
	"net/http"
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

}
