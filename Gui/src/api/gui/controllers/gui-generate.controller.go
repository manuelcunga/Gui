package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/manuelcunga/Gui/Gui/src/usecases/generate"
)

type GenerateGuiGPTController struct {
	GuiGPTGeneratorUsecase usecase.GPTGeneratorUsecase
}

type RequestData struct {
	Body string // Atualize o nome do campo conforme necessário
}

func NewGenerateController(generator usecase.GPTGeneratorUsecase) *GenerateGuiGPTController {
	return &GenerateGuiGPTController{
		GuiGPTGeneratorUsecase: generator,
	}
}


func (ctrl *GenerateGuiGPTController) Handle(c echo.Context) error {
	requestBody := RequestData{}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	fmt.Println("Mensagem do body:", requestBody.Body)

	gptText, err := ctrl.GuiGPTGeneratorUsecase.GenerateText(requestBody.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to generate GPT text",
		})
	}

	response := struct {
		Text string `json:"text"`
	}{
		Text: gptText,
	}

	fmt.Println("Response do gpt:", response)

	return c.JSON(http.StatusOK, response)
}
