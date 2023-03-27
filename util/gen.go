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

func Contains[T string | int](slice []T, target T) bool {
	for _, x := range slice {
		if x == target {
			return true
		}
	}
	return false
}
