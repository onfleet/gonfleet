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

func IsErrorStatus(status int) bool {
	return status < 200 || status > 299
}

func Contains[T string | int](slice []T, target T) bool {
	for _, x := range slice {
		if x == target {
			return true
		}
	}
	return false
}

func UrlAttachQuery(baseUrl string, params map[string]string) string {
	URL, err := url.Parse(baseUrl)
	if err != nil {
		return baseUrl
	}
	q := URL.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	URL.RawQuery = q.Encode()
	return URL.String()
}

func UrlAttachPath(baseUrl string, pathSegments ...string) string {
	newUrl, err := url.JoinPath(baseUrl, pathSegments...)
	if err != nil {
		return baseUrl
	}
	return newUrl
}
