package main

import (
	"net/http"
	"pedrobinotto/template"
	"pedrobinotto/utils/md"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()

	e.Static("/dist", "dist")

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	template.NewTemplateRenderer(e, "public/*.html")
	e.GET("/hello", func(e echo.Context) error {
		res := map[string]interface{}{
			"Name": "Pedro",
		}
		return e.Render(http.StatusOK, "index", res)
	})

	e.GET("/get-info", func(c echo.Context) error {
		res := map[string]interface{}{
			"Name":  "Pedro Binotto",
			"Phone": "00000000",
			"Email": "pedro@email.com",
		}
		return c.Render(http.StatusOK, "name_card", res)
	})

	e.GET("/markdown", func(c echo.Context) error {
		markdowns := "# Header1\n## Header 2"
		parsed, err := md.ParseMD(markdowns)
		if err != nil {
			return err
		}
		res := map[string]interface{}{
			"ParsedMarkdown": parsed,
		}
		return c.Render(http.StatusOK, "markdown", res)
	})

	e.GET("/editor", func(c echo.Context) error {
		return c.Render(http.StatusOK, "editor", nil)
	})

	e.POST("/editor/render", func(c echo.Context) error {
		var payload md.RenderReq
		err := c.Bind(&payload)
		if err != nil {
			return err
		}
		parsed, err := md.ParseMD(payload.Content)
		if err != nil {
			return err
		}
		res := map[string]interface{}{
			"ParsedMarkdown": parsed,
		}
		return c.Render(http.StatusOK, "markdown", res)
	})

	e.Logger.Fatal(e.Start(":4040"))
}
