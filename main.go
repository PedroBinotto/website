package main

import (
	"context"

	"github.com/PedroBinotto/website/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		index := templates.Layout(templates.Index())
		return index.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/hello", func(c echo.Context) error {
		component := templates.Layout(templates.Hello("Pedro"))
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Static("/static", "static")
	e.Static("/css", "css")

	e.Logger.Fatal((e.Start(":3000")))
}
