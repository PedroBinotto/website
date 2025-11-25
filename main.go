package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

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

	e := echo.New()
	e.Logger.SetOutput(logFile)
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_HOST", ""),
		getEnv("DB_PORT", ""),
		getEnv("DB_NAME", ""),
	)

	// Change: Use "postgres" driver instead of "sqlite3"
	db, dbErr := sql.Open("postgres", dbURL)
	e.Logger.Info("PostgreSQL connection string")
	e.Logger.Info(dbURL)

	if dbErr != nil {
		e.Logger.Fatal(dbErr)
		panic("Unable to connect to application database.")
	}

	// Add: Test the connection
	if pingErr := db.Ping(); pingErr != nil {
		e.Logger.Fatal(pingErr)
		panic("Unable to ping database.")
	}

	// Add: Configure connection pool (optional but recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

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
		blogs, _ := queries.GetBlogs(context.Background())
		c.Echo().Logger.Info("DB entries: ")
		c.Echo().Logger.Info(strconv.Itoa(len(blogs)))
		component := templates.Layout(templates.Blogs(blogs))
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

// Helper function to get environment variable with default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
