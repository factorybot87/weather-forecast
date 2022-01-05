package weather

import (
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
