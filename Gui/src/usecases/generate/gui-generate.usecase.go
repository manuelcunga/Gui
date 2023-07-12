package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gui_types "github.com/manuelcunga/Gui/Gui/src/domain/types"
)

type GPTGeneratorUsecase interface {
	GenerateText(query string) (string, error)
}

type OpenAIGenerator struct {
	Client *http.Client
	Token  string
}

func NewOpenAIGeneratorUsecase(client *http.Client, token string) *OpenAIGenerator {
	return &OpenAIGenerator{
		Client: client,
		Token:  token,
	}
}

func (o *OpenAIGenerator) GenerateText(query string) (string, error) {
	req := gui_types.Request{
		Model: "gpt-3.5-turbo",
		Messages: []gui_types.Messages{
			{
				Role:    "user",
				Content: query,
			},
		},
		MaxTokens: 300,
	}

	reqJSON, err := json.Marshal(req)
	fmt.Println("convertendo em json:", req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	token := o.Token
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	fmt.Println("Token from headers", request.Header.Get("token"))

	response, err := o.Client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// responseBody, err := ioutil.ReadAll(response.Body)
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("resposta do gpt32:", string(responseBody))

	var res gui_types.Response
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}
