package weather_test

import (
	"bytes"
	"fmt"
	"forecast/weather"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestGenerateURL(t *testing.T) {

	location := "taipei"
	key := "dummyKey"

	want := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/taipei?unitGroup=metric&include=days%2Ccurrent&key=dummyKey&contentType=json"

	got := weather.GenerateURL(location, key)

	assertEqual(want, got, t)
}

func TestBuildRequest(t *testing.T) {

	url := "example.com"

	wantURL := url
	wantMethod := http.MethodGet

	got := weather.BuildRequest(url)

	if got != nil {
		assertEqual(wantMethod, got.Method, t)
		assertEqual(wantURL, got.URL.String(), t)
	} else {
		t.Fatal("Expect a request, not nil.")
	}
}

func TestGetBody(t *testing.T) {
	dummyResponse := http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString("OK")),
		Header:     make(http.Header),
	}

	want := "OK"
	got := weather.GetBody(&dummyResponse)

	assertEqual(want, got, t)
}

func TestParseResponseBody(t *testing.T) {

	t.Run("empty body returns empty result", func(t *testing.T) {
		jsonData := "{}"

		want := weather.Forecast{}
		got, err := weather.ParseResponseBody(jsonData)

		assertNoError(err)

		if !reflect.DeepEqual(want, got) {
			t.Fatal("did not match")
		}
	})

	t.Run("good response populates result", func(t *testing.T) {
		jsonData := `{
			"currentConditions": {
				"conditions":"Partially cloudy"
			},
			"days": [
				{"description":"Cloudy skies throughout the day with late afternoon rain."}
			]
		}`

		want := weather.Forecast{
			weather.CurrentCondition{
				Conditions: "Partially cloudy",
			},
			[]weather.DailyWeather{
				{
					Description: "Cloudy skies throughout the day with late afternoon rain.",
				},
			},
		}
		got, err := weather.ParseResponseBody(jsonData)

		assertNoError(err)

		if !reflect.DeepEqual(want, got) {
			t.Fatal("did not match")
		}
	})
}

func assertEqual(want, got string, t *testing.T) {
	if got != want {
		t.Fatal(outputMessage(want, got))
	}
}

func outputMessage(want, got string) string {
	return fmt.Sprint("\nexpect:\n", want, "\n\n got:\n", got)
}

func assertNoError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
