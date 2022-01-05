package weather_test

import (
	"forecast/weather"
	"testing"
)

func TestGenerateURL(t *testing.T) {

	if weather.GenerateURL() != "url" {
		t.Fatal("wrong")
	}
}
