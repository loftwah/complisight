package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check AWS configuration and connectivity",
	Long:  `This command checks if the AWS configuration is correctly set up and if the AWS services are accessible.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Unable to load AWS SDK config, %v", err)
		}

		fmt.Println("AWS configuration loaded successfully.")

		// Example check: List S3 buckets
		s3Client := s3.NewFromConfig(cfg)
		_, err = s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
		if err != nil {
			log.Fatalf("Unable to list S3 buckets, %v", err)
		}
		fmt.Println("Successfully accessed S3 service. AWS configuration and connectivity check passed.")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
