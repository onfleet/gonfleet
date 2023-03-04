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
