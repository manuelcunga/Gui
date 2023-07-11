package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	controller "github.com/manuelcunga/Gui/Gui/src/api/gui/controllers"
	usecase "github.com/manuelcunga/Gui/Gui/src/usecases/generate"
)

func main() {
	e := echo.New()

	generatorUsecase := usecase.NewOpenAIGenerator(http.DefaultClient, "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")

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
