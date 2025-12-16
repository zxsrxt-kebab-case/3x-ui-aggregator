package Utils

import (
	"encoding/base64"
	"errors"
)

func DecodeBase64(s string) (string, error) {
	if d, err := base64.StdEncoding.DecodeString(s); err == nil {
		return string(d), nil
	}
	if d, err := base64.RawStdEncoding.DecodeString(s); err == nil {
		return string(d), nil
	}
	if d, err := base64.URLEncoding.DecodeString(s); err == nil {
		return string(d), nil
	}
	return "", errors.New("invalid base64")
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
