package main

import (
	"log"
	"os"

	client "github.com/adii1203/pixify-transformer/S3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Client, err := client.NewS3Client()
	if err != nil {
		log.Fatal("Error creating S3 client")
	}

	port := os.Getenv("PORT")
	e := echo.New()

	e.GET("/:id", func(c echo.Context) error {
		object, err := s3Client.GetObjectFromRawBucket(c.Param("id"))
		object.Body.Close()

		if err != nil {
			return c.JSON(500, "Error fetching image")
		}

		return c.Stream(200, *object.ContentType, object.Body)
	})

	e.Logger.Fatal(e.Start(":" + port))
}
