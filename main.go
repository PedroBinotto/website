package main

import (
	"context"
	"os"

	"github.com/PedroBinotto/website/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	index := templates.Index()

	component := templates.Layout(templates.Hello("Pedro"))
	component.Render(context.Background(), os.Stdout)

	e.GET("/", func(c echo.Context) error {
		return index.Render(context.Background(), c.Response().Writer)
		// return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/static", "static")
	e.Logger.Fatal((e.Start(":3000")))
}
