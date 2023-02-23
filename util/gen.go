package util

import "encoding/json"

func PrettyPrint(data any) (string, error) {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
