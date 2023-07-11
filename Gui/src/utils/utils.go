package utils

import (
	"encoding/base64"
	"errors"
	"strings"
)

func ParseBase64RequestData(response string) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(response)
	if err != nil {
		return "", err
	}

	body := strings.TrimSpace(string(dataBytes))
	if body != "" {
		return body, nil
	}

	return "", errors.New("Body not found")
}
