package main

import (
	"context"

	"database/sql"

	generated "github.com/PedroBinotto/website/sqlc-generated"
	"github.com/PedroBinotto/website/templates"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
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

	e.GET("/resume", func(c echo.Context) error {
		component := templates.Layout(templates.Resume())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/blogs", func(c echo.Context) error {

		db, err := sql.Open("sqlite3", "./app.db")
		if err != nil {
			println("Unable to connect to application database.")
		}

		// Testing sql connection
		queries := generated.New(db)
		users, err := queries.GetUsers(context.Background())
		println(len(users))

		component := templates.Layout(templates.Blogs())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/playground", func(c echo.Context) error {
		component := templates.Layout(templates.Playground())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Static("/static", "static")
	e.Static("/css", "css")

	e.Logger.Fatal((e.Start(":8081")))
}
