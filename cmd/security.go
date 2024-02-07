package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
)

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

	buckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to list buckets: %v", err)
	}

	for _, bucket := range buckets.Buckets {
		fmt.Printf("Checking bucket: %s\n", aws.ToString(bucket.Name))
		checkBucketEncryption(ctx, s3Client, aws.ToString(bucket.Name))
		checkBucketPublicAccessBlock(ctx, awsConfig, aws.ToString(bucket.Name))
	}
}

func checkBucketEncryption(ctx context.Context, s3Client *s3.Client, bucketName string) {
	result, err := s3Client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("Bucket is not encrypted or unable to check encryption:", bucketName)
	} else {
		fmt.Println("Bucket is encrypted:", bucketName, "Algorithm:", result.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm)
	}
}

func checkBucketPublicAccessBlock(ctx context.Context, cfg aws.Config, bucketName string) {
	stsClient := sts.NewFromConfig(cfg)
	identity, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Unable to get AWS account ID: %v", err)
	}

	s3ControlClient := s3control.NewFromConfig(cfg)
	input := &s3control.GetPublicAccessBlockInput{
		AccountId: identity.Account,
	}

	result, err := s3ControlClient.GetPublicAccessBlock(ctx, input)
	if err != nil {
		fmt.Printf("Unable to check public access block for account %s: %v\n", *identity.Account, err)
	} else {
		fmt.Printf("Public access block configuration for account %s: %+v\n", *identity.Account, result.PublicAccessBlockConfiguration)
	}
}

func init() {
	rootCmd.AddCommand(securityCmd)
}
