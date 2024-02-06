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

// Assume this global variable is set up elsewhere, like in your main.go or during command initialization
var awsConfig aws.Config

var securityCmd = &cobra.Command{
	Use:   "security",
	Short: "Perform security compliance checks for S3 buckets",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Unable to load AWS SDK config, %v", err)
		}
		awsConfig = cfg

		checkS3Buckets(ctx)
	},
}

func checkS3Buckets(ctx context.Context) {
	s3Client := s3.NewFromConfig(awsConfig)

	// List all S3 buckets
	buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to list buckets: %v", err)
	}

	for _, bucket := range buckets.Buckets {
		fmt.Printf("Checking bucket: %s\n", aws.ToString(bucket.Name))
		// Placeholder for bucket encryption and public access checks
		// In practice, you'd call your S3Service methods here
		checkBucketEncryption(ctx, s3Client, aws.ToString(bucket.Name))
		checkBucketPublicAccessBlock(ctx, s3Client, aws.ToString(bucket.Name))
	}
}

func checkBucketEncryption(ctx context.Context, s3Client *s3.Client, bucketName string) {
	// Example function to check bucket encryption
	// This is a simplified version; actual implementation may vary
	result, err := s3Client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("Bucket is not encrypted or unable to check encryption:", bucketName)
	} else {
		fmt.Println("Bucket is encrypted:", bucketName, "Algorithm:", result.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm)
	}
}

func checkBucketPublicAccessBlock(ctx context.Context, s3Client *s3.Client, bucketName string) {
	// Example function to check bucket public access block
	result, err := s3Client.GetBucketPublicAccessBlock(ctx, &s3.GetBucketPublicAccessBlockInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("Bucket public access block is not enabled or unable to check:", bucketName)
	} else {
		fmt.Printf("Bucket public access block is enabled for %s\n", bucketName)
	}
}

func init() {
	rootCmd.AddCommand(securityCmd)
}
