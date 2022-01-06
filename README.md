# Forecast

A CLI tool to get the current weather information and the forecast of the next 6 days.

## Configuration

You need a Visual Crossing account to access the Timeline API.

Replace the `key` constant in the `main.go` with your API key.

`location` is default to Taipei.

## Usage

`go run ./main.go`

## Project structure

`weather/` contains a http client and a json parser to interacts with the Timeline API.

## Screenshot