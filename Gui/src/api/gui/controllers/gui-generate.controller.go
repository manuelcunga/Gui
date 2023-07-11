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

func NewGenerateController(generator usecase.GPTGeneratorUsecase) *GenerateGuiGPTController {
	return &GenerateGuiGPTController{
		GuiGPTGeneratorUsecase: generator,
	}
}

func (ctrl *GenerateGuiGPTController) Handle(c echo.Context) error {
	var requestBody struct {
		Body string `json:"body"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	fmt.Println("mensagem vindo do user:", requestBody.Body)

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

	fmt.Println("Resposta do GPT:", response)
	return c.JSON(http.StatusOK, response)
}
