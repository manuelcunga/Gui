package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type RequestData struct {
	Content string `json:"content"`
}

func ParseBase64RequestData(response string) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(response)
	if err != nil {
		return "", err
	}

	var reqData RequestData
	if err := json.Unmarshal(dataBytes, &reqData); err != nil {
		return "", err
	}

	if reqData.Content != "" {
		return reqData.Content, nil
	}

	return "", errors.New("Content not found")
}
