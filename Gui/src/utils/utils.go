package utils

import (
	"encoding/base64"
	"errors"
	"net/url"
)

func ParseBase64RequestData(response string) (string, error) {
	data, err := url.ParseQuery(response)
	if err != nil {
		return "", err
	}

	body := data.Get("Body")
	if body == "" {
		return "", errors.New("Body not found")
	}

	decoded, err := base64.URLEncoding.DecodeString(body)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
