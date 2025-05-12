package main

import (
	"context"
	"os"
	"strconv"

	"database/sql"

	generated "github.com/PedroBinotto/website/sqlc-generated"
	"github.com/PedroBinotto/website/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logFile, logFileErr := os.OpenFile("server.log.jsonl", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logFileErr != nil {
		panic(logFileErr)
	}

	e := echo.New()
	e.Logger.SetOutput(logFile)
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))

	sqliteDB := os.Getenv("SQLITE_DB")
	db, dbErr := sql.Open("sqlite3", sqliteDB)

	e.Logger.Info("sqliteDB")
	e.Logger.Info(sqliteDB)

	if dbErr != nil {
		e.Logger.Fatal(dbErr)
		panic("Unable to connect to application database.")
	}

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
		queries := generated.New(db)

		insertedBlog, err := queries.CreateBlog(context.Background())

		c.Echo().Logger.Info(insertedBlog)
		c.Echo().Logger.Info(err)

		blogs, _ := queries.GetBlogs(context.Background())

		c.Echo().Logger.Info("DB entries: ")
		c.Echo().Logger.Info(strconv.Itoa(len(blogs)))

		component := templates.Layout(templates.Blogs())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/playground", func(c echo.Context) error {
		component := templates.Layout(templates.Playground())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Static("/", "public")

	e.Logger.Info("Server starting...")
	e.Logger.Fatal((e.Start(":8080")))
}
