package main

import (
	"context"
	"net/http"
	"os"

	"github.com/PedroBinotto/website/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	component := templates.Hello("Pedro")
	component.Render(context.Background(), os.Stdout)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/static", "static")
	e.Logger.Fatal((e.Start(":3000")))
}
