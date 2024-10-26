package infra

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type Presigner struct {
	PresignClient *s3.PresignClient
}

func NewPresigner(presignClient *s3.PresignClient) *Presigner {
	return &Presigner{
		PresignClient: presignClient,
	}
}

func (presigner Presigner) PutObject(ctx context.Context, bucketName string, objectKey string, lifetimeSecs int64, file io.Reader) (*v4.PresignedHTTPRequest, error) {
	request, err := presigner.PresignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   strings.NewReader(""),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to put %v:%v. Here's why: %v\n", bucketName, objectKey, err)
	}
	return request, err
}

type BucketBasics struct {
	S3Client *s3.Client
}

func NewBucketBasics() *BucketBasics {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2")) // リージョンを指定
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
	return err
}
