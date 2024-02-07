package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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
	// Read target regions from environment variable
	targetRegions := strings.Split(os.Getenv("TARGET_REGIONS"), ",")

	s3Client := s3.NewFromConfig(cfg)

	// List all S3 buckets
	buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to list S3 buckets: %v", err)
	}

	for _, bucket := range buckets.Buckets {
		// Get the region of each bucket
		regionOutput, err := s3Client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
			Bucket: bucket.Name,
		})
		if err != nil {
			fmt.Printf("Unable to get location for bucket %s: %v\n", *bucket.Name, err)
			continue
		}

		// Dynamically handle the bucket region
		bucketRegion := resolveBucketRegion(regionOutput.LocationConstraint)

		// Check if the bucket's region is in the list of target regions
		if contains(targetRegions, bucketRegion) {
			_, err := s3Client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
				Bucket: bucket.Name,
			})

			if err != nil {
				fmt.Printf("Bucket %s in region %s fails confidentiality compliance: Encryption is not enabled or cannot be verified.\n", *bucket.Name, bucketRegion)
			} else {
				fmt.Printf("Bucket %s in region %s passes confidentiality compliance: Encryption is enabled.\n", *bucket.Name, bucketRegion)
			}
		}
	}
}

// Helper function to resolve bucket region from LocationConstraint
func resolveBucketRegion(constraint types.BucketLocationConstraint) string {
	// Special handling for the "EU" location constraint
	if constraint == types.BucketLocationConstraintEu || constraint == "" {
		return "eu-west-1"
	}
	return string(constraint)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(confidentialityCmd)
}
