package infra

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type BucketBasics struct {
	S3Client *s3.Client
}

func NewBucketBasics() *BucketBasics {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2")) // リージョンを指定
	if err != nil {
		fmt.Println("unable to load SDK config, " + err.Error())
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// S3クライアントを作成
	svc := s3.NewFromConfig(cfg)
	return &BucketBasics{
		S3Client: svc,
	}
}

func (i *Infrastructure) UploadFile(ctx context.Context, bucketName string, objectKey string, file io.Reader) error {
	_, err := i.S3.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return err
	}
	return nil
}
