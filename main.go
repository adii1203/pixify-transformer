package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// _, err := s3.NewS3Client()
	// if err != nil {
	// 	fmt.Println("Error creating S3 client: ", err.Error())
	// 	os.Exit(1)
	// }
	port := os.Getenv("PORT")

	e := echo.New()

	e.GET("/:id", func(c echo.Context) error {
		path := c.Param("id")
		fmt.Println("Path: " + path + "\n")
		fmt.Println(c.Request().URL)

		return c.String(http.StatusOK, path)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
