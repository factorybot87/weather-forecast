package weather_test

import (
	"fmt"
	"forecast/weather"
	"testing"
)

func TestGenerateURL(t *testing.T) {

	location := "Taipei"
	key := "dummyKey"

	want := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/Taipei?key=dummyKey"
	got := weather.GenerateURL(location, key)

	assertEqual(want, got, t)
}

func assertEqual(want, got string, t *testing.T) {
	if got != want {
		t.Fatal(outputMessage(want, got))
	}
}

func outputMessage(want, got string) string {
	return fmt.Sprint("\nexpect:\n", want, "\n\n got:\n", got)
}
