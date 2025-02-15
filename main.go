package main

import (
	"fmt"
	"net/http"
	"os"

	s3 "github.com/adii1203/pixify-transformer/S3"
	"github.com/labstack/echo/v4"
)

func main() {
	_, err := s3.NewS3Client()
	if err != nil {
		fmt.Println("Error creating S3 client: ", err.Error())
		os.Exit(1)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	e := echo.New()

	e.GET("/:id", func(c echo.Context) error {
		path := c.Param("id")
		fmt.Println("Path: " + path + "\n")
		fmt.Println(c.Request().URL)

		return c.String(http.StatusOK, path)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
