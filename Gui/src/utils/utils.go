package utils

import (
	"encoding/base64"
	"errors"
	"net/url"
)

func ParseBase64RequestData(response string) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(response)

	if err != nil {
		return "", err
	}

	data, err := url.ParseQuery(string(dataBytes))
	if data.Has("Body") {
		return data.Get("Body"), nil
	}

	return "", errors.New("Body not found")
}