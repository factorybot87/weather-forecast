package weather_test

import (
	"bytes"
	"fmt"
	"forecast/weather"
	"io"
	"net/http"
	"testing"
)

func TestGenerateURL(t *testing.T) {

	location := "Taipei"
	key := "dummyKey"

	want := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/Taipei?key=dummyKey"
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

func assertEqual(want, got string, t *testing.T) {
	if got != want {
		t.Fatal(outputMessage(want, got))
	}
}

func outputMessage(want, got string) string {
	return fmt.Sprint("\nexpect:\n", want, "\n\n got:\n", got)
}
