package conv

import (
	"encoding/base64"
)

// EncodeBase64 encodes string into a base64 string
func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
