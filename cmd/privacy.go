package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

var privacyCmd = &cobra.Command{
	Use:   "privacy",
	Short: "Check privacy compliance for AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Unable to load AWS SDK config, %v", err)
		}

		checkS3BucketPrivacy(ctx, cfg)
	},
}

func checkS3BucketPrivacy(ctx context.Context, cfg aws.Config) {
	s3Client := s3.NewFromConfig(cfg)

	buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to list S3 buckets: %v", err)
	}

	for _, bucket := range buckets.Buckets {
		bucketName := aws.ToString(bucket.Name)
		fmt.Printf("Checking privacy compliance for bucket: %s\n", bucketName)
		passes := true

		// Check bucket encryption
		_, err := s3Client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
			Bucket: aws.String(bucketName),
		})
		encryptionStatus := "Not Encrypted"
		if err == nil {
			encryptionStatus = "Encrypted"
		} else {
			passes = false
		}

		// Check public access block - adjust this part according to the correct API if needed
		publicAccessStatus := "Public Access Check Skipped" // Placeholder, adjust based on actual API availability and your requirements

		// Simplified policy check
		policyStatus := "No Policy"
		getPolicyOutput, err := s3Client.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{
			Bucket: aws.String(bucketName),
		})
		if err == nil {
			policyStatus = "Policy Found"
			if !containsHttps(*getPolicyOutput.Policy) {
				policyStatus += " (Insecure Transport)"
				passes = false
			}
		}

		status := "FAIL"
		if passes {
			status = "PASS"
		}
		fmt.Printf("Bucket: %s, Encryption: %s, Public Access: %s, Policy: %s, Compliance Status: %s\n", bucketName, encryptionStatus, publicAccessStatus, policyStatus, status)
	}
}

func containsHttps(policy string) bool {
	// Real policy check logic should be more comprehensive, considering actual security requirements
	return strings.Contains(policy, "https")
}

func init() {
	rootCmd.AddCommand(privacyCmd)
}
