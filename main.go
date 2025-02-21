package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	client "github.com/adii1203/pixify-transformer/S3"
	"github.com/adii1203/pixify-transformer/transformer"
	"github.com/adii1203/pixify-transformer/utils"
	"github.com/h2non/bimg"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	v := bimg.Version
	fmt.Println("bimg Version:", v)
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

	e.GET("/:id/:tr", func(c echo.Context) error {
		object, err := s3Client.GetObjectFromRawBucket(c.Param("id"))
		if err != nil {
			return c.JSON(500, "Error fetching image")
		}

		defer object.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(object.Body)
		m := utils.ExtractTransformationsOptions(c.Param("tr"))
		img, err := transformer.ApplyTransformations(buf.Bytes(), m)
		if err != nil {
			return c.JSON(500, "Error transforming image")
		}

		cacheKey := fmt.Sprintf("%s/%s", c.Param("id"), c.Param("tr"))
		err = s3Client.PutObjectInProcessedBucket(cacheKey, img, object.ContentType)
		if err != nil {
			return c.JSON(500, "Error saving image")
		}

		return c.Stream(200, *object.ContentType, bytes.NewReader(img))
	})

	e.GET("/:key", func(c echo.Context) error {
		object, err := s3Client.GetObjectFromRawBucket(c.Param("key"))
		if err != nil {
			return c.JSON(500, "Error fetching image")
		}

		defer object.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(object.Body)

		img := bimg.NewImage(buf.Bytes())

		err = s3Client.PutObjectInProcessedBucket(c.Param("key"), img.Image(), object.ContentType)
		if err != nil {
			return c.JSON(500, "Error saving image")
		}

		return c.Stream(200, *object.ContentType, bytes.NewReader(img.Image()))
	})

	e.Logger.Fatal(e.Start(":" + port))
}
