package weather

import (
	"fmt"
	"io"
	"net/http"
)

// GenerateURL will Construct the Visual Crossing API url from components
func GenerateURL(location, key string) string {
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"
	url := fmt.Sprintf("%s/%s?key=%s", baseURL, location, key)
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
