package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/PedroBinotto/website/shared"
	generated "github.com/PedroBinotto/website/sqlc-generated"
	"github.com/PedroBinotto/website/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq" // Change: PostgreSQL driver instead of sqlite3
)

func main() {
	logFile, logFileErr := os.OpenFile("server.log.jsonl", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if logFileErr != nil {
		panic(logFileErr)
	}
	db, dbErr := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnv("DB_USER"), getEnv("DB_PASSWORD"), getEnv("DB_HOST"), getEnv("DB_PORT"), getEnv("DB_NAME"),
	))
	if dbErr != nil {
		panic("Unable to connect to application database.")
	}

	e := echo.New()
	e.Logger.SetOutput(logFile)
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))

	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		index := templates.Layout(templates.Index())
		return index.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/hello", func(c echo.Context) error {
		component := templates.Layout(templates.Hello())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/resume", func(c echo.Context) error {
		component := templates.Layout(templates.Resume())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/playground", func(c echo.Context) error {
		component := templates.Layout(templates.Playground())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/blogs", func(c echo.Context) error {
		/* TODO: make this route a paged list
		 * Also: find out a way to include images in posts; develop post view */
		queries := generated.New(db)
		rows, err := queries.GetBlogs(c.Request().Context()) // Use request context instead
		if err != nil {
			return err
		}

		result := make([]shared.BlogPayload, 0, len(rows))
		for _, r := range rows {
			transformed := shared.BlogPayload{
				ID:        r.ID,
				Title:     r.Title,
				Synopsis:  r.Synopsis,
				Url:       r.Url,
				Body:      r.Body,
				CreatedAt: r.CreatedAt,
			}

			if tagBytes, ok := r.Tags.([]byte); ok && len(tagBytes) > 0 {
				if err := json.Unmarshal(tagBytes, &transformed.Tags); err != nil {
					log.Printf("failed to unmarshal tags for blog %d: %v", r.ID, err)
				}
			}

			result = append(result, transformed)
		}

		component := templates.Layout(templates.Blogs(result))
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Info("Server starting...")
	e.Logger.Fatal((e.Start(":8080")))
}

// Helper function to get environment variable with default value
func getEnv(key string) string {
	return os.Getenv(key)
}
