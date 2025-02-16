package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"

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
		image := getFromS3(s3Client)

		return c.JSON(200, image.ETag)
	})

	e.Logger.Fatal(e.Start(":" + port))
}

func getFromS3(s3Client *client.S3Client) *s3.GetObjectOutput {
	output, err := s3Client.S3.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String("pixify-raw-images-bucket"),
		Key:    aws.String("114096753.jpg"),
	})
	if err != nil {
		fmt.Println("Error getting object from S3")
	}

	return output
}
