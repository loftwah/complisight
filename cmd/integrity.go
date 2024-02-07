package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/spf13/cobra"
)

var integrityCmd = &cobra.Command{
	Use:   "integrity",
	Short: "Check integrity compliance for AWS resources",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Unable to load AWS SDK config, %v", err)
		}

		checkCloudTrail(ctx, cfg)
	},
}

func checkCloudTrail(ctx context.Context, cfg aws.Config) {
	ctClient := cloudtrail.NewFromConfig(cfg)

	// List all trails
	trails, err := ctClient.DescribeTrails(ctx, &cloudtrail.DescribeTrailsInput{})
	if err != nil {
		log.Fatalf("Unable to describe CloudTrail trails: %v", err)
	}

	if len(trails.TrailList) == 0 {
		fmt.Println("No CloudTrail trails found. Compliance check failed.")
		return
	}

	for _, trail := range trails.TrailList {
		// Check if the trail is logging
		status, err := ctClient.GetTrailStatus(ctx, &cloudtrail.GetTrailStatusInput{
			Name: aws.String(*trail.TrailARN),
		})
		if err != nil {
			fmt.Printf("Unable to get status for trail %s: %v\n", *trail.Name, err)
			continue
		}

		if aws.ToBool(status.IsLogging) {
			fmt.Printf("Trail %s is enabled and logging.\n", *trail.Name)
		} else {
			fmt.Printf("Trail %s is not logging. It may be managed outside of this account, or not configured correctly. Verify its configuration, especially if managed at the organizational level.\n", *trail.Name)
		}
	}
}

func init() {
	rootCmd.AddCommand(integrityCmd)
}
