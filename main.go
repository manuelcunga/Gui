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

	generatorUsecase := usecase.NewOpenAIGenerator(http.DefaultClient, token)

	generateController := controller.NewGenerateHandler(generatorUsecase)

	e.GET("/", home())
	e.POST("/api/v1/gui/generate", generateController.Handle)

	e.Logger.Fatal(e.Start(":5000"))

}

func home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(200, "Welcome to GUI-IA!")
	}
}
