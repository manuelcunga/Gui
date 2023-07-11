package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/manuelcunga/Gui/Gui/src/usecases/generate"
	"github.com/manuelcunga/Gui/Gui/src/utils"
)

// type GenerateGuiGPTController struct {
// 	GPTGenerator contract.IGPTGenerator
// }

type GenerateGuiGPTController struct {
	GuiGPTGeneratorUsecase usecase.GPTGeneratorUsecase
}

func NewGenerateHandler(generator usecase.GPTGeneratorUsecase) *GenerateGuiGPTController {
	return &GenerateGuiGPTController{
		GuiGPTGeneratorUsecase: generator,
	}
}

func (ctrl *GenerateGuiGPTController) Handle(c echo.Context) error {
	requestBody := struct {
		Body string `json:"body"`
	}{}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	text, err := utils.ParseBase64RequestData(requestBody.Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	gptText, err := ctrl.GuiGPTGeneratorUsecase.GenerateText(text)
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

	return c.JSON(http.StatusOK, response)
}