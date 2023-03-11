package util

import (
	"encoding/json"
	"fmt"
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

func Contains[T string | int](slice []T, target any) bool {
	for _, x := range slice {
		if x == target {
			return true
		}
	}
	return false
}
