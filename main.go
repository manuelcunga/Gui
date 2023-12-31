package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	controller "github.com/manuelcunga/Gui/Gui/src/api/gui/controllers"
	usecase "github.com/manuelcunga/Gui/Gui/src/usecases/generate"
)

func main() {
	e := echo.New()
	token := os.Getenv("GPT_TOKEN")
	generatorUsecase := usecase.NewOpenAIGeneratorUsecase(http.DefaultClient, token)

	generateController := controller.NewGenerateController(generatorUsecase)

	e.GET("/", home())
	e.POST("/", generateController.Handle)

	e.Logger.Fatal(e.Start(":5000"))

}

func home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to GUI-IA!")
	}
}
