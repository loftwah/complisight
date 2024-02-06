package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Service wraps the AWS SDK's S3 client
type S3Service struct {
	Client *s3.Client
}

// NewS3Service creates a new instance of S3Service
func NewS3Service(ctx context.Context) (*S3Service, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &S3Service{
		Client: s3.NewFromConfig(cfg),
	}, nil
}

// ListBuckets lists all S3 buckets in the account
func (s *S3Service) ListBuckets(ctx context.Context) ([]string, error) {
	result, err := s.Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("unable to list buckets, %v", err)
	}

	buckets := make([]string, len(result.Buckets))
	for i, bucket := range result.Buckets {
		buckets[i] = aws.ToString(bucket.Name)
	}

	return buckets, nil
}
