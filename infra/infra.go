package infra

type Infrastructure struct {
	S3 *BucketBasics
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		// NewPresigner(presignClient *s3.PresignClient
		S3: NewBucketBasics(),
	}
}
