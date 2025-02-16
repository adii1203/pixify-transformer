package main

import (
	"bytes"
	"log"
	"os"

	client "github.com/adii1203/pixify-transformer/S3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	s3Client, err := client.NewS3Client()
	if err != nil {
		log.Fatal("Error creating S3 client", err.Error())
	}

	port := os.Getenv("PORT")
	e := echo.New()

	e.GET("/:id", func(c echo.Context) error {
		object, err := s3Client.GetObjectFromRawBucket(c.Param("id"))
		if err != nil {
			return c.JSON(500, "Error fetching image")
		}

		defer object.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(object.Body)

		return c.Stream(200, *object.ContentType, bytes.NewReader(buf.Bytes()))
	})

	e.Logger.Fatal(e.Start(":" + port))
}
