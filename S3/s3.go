package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	S3 *s3.Client
}

func NewS3Client() (*S3Client, error) {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(config)
	return &S3Client{
		S3: client,
	}, nil
}

func (s3Client *S3Client) GetObjectFromRawBucket(key string) (*s3.GetObjectOutput, error) {
	output, err := s3Client.S3.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String("pixify-raw-images-bucket"),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s3client *S3Client) PutObjectInProcessedBucket(key string, body []byte, contentType *string) error {
	_, err := s3client.S3.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String("pixify-transformed-images-bucket"),
		Key:         aws.String(key),
		Body:        bytes.NewReader(body),
		ContentType: contentType,
	})
	if err != nil {
		return err
	}
	return nil
}
