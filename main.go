package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not set
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		path := c.Path()
		fmt.Println(path)
		fmt.Println(c.Request().URL)

		return c.String(http.StatusOK, path)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
