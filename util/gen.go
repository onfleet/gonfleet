package util

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func PrettyPrint(data any) {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return
	}
	fmt.Println(string(b))
}

func Contains[T string | int](slice []T, target T) bool {
	for _, x := range slice {
		if x == target {
			return true
		}
	}
	return false
}

// Stomp converts a struct to a map[string]any
func Stomp(v any) (map[string]any, error) {
	m := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return m, err
	}
	err = json.Unmarshal(b, &m)
	return m, err
}

// UrlAttachQuery sets query parameters on the provided baseUrl.
func UrlAttachQuery(baseUrl string, v any) string {
	URL, err := url.Parse(baseUrl)
	if err != nil {
		return baseUrl
	}
	q := URL.Query()
	params, err := Stomp(v)
	if err != nil {
		return baseUrl
	}
	for k, v := range params {
		q.Set(k, fmt.Sprintf("%v", v))
	}
	URL.RawQuery = q.Encode()
	return URL.String()
}

// UrlAttachPath appends path segments onto provided baseUrl.
func UrlAttachPath(baseUrl string, pathSegments ...string) string {
	newUrl, err := url.JoinPath(baseUrl, pathSegments...)
	if err != nil {
		return baseUrl
	}
	return newUrl
}
