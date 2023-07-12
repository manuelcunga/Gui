package utils

import (
	"encoding/base64"
	"errors"
	"net/url"
)

func ParseBase64RequestData(request string) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(request)

	if err != nil {
		return "", err
	}

	// content body

	data, err := url.ParseQuery(string(dataBytes))
	if data.Has("Body") {
		return data.Get("Body"), nil
	}

	return "", errors.New("Body not found")
}
