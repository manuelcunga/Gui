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

	data, err := url.ParseQuery(string(dataBytes))
	if err != nil {
		return "", err
	}

	if body := data.Get("Body"); body != "" {
		return body, nil
	}

	return "", errors.New("Body not found")
}
