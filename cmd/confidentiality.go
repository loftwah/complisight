package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

var confidentialityCmd = &cobra.Command{
	Use:   "confidentiality",
	Short: "Check confidentiality compliance for AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Unable to load AWS SDK config, %v", err)
		}

		checkS3BucketEncryption(ctx, cfg)
	},
}

func checkS3BucketEncryption(ctx context.Context, cfg aws.Config) {
	s3Client := s3.NewFromConfig(cfg)

	// List all S3 buckets
	buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to list S3 buckets: %v", err)
	}

	for _, bucket := range buckets.Buckets {
		// Check encryption on each bucket
		bucketName := aws.ToString(bucket.Name)
		result, err := s3Client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
			Bucket: aws.String(bucketName),
		})

		if err != nil {
			fmt.Printf("Bucket %s does not have encryption enabled or it cannot be verified: %v\n", bucketName, err)
		} else {
			fmt.Printf("Bucket %s has encryption enabled with %s algorithm.\n", bucketName, result.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm)
		}
	}
}

func init() {
	rootCmd.AddCommand(confidentialityCmd)
}
