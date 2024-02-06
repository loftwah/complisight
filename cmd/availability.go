package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/spf13/cobra"
)

var availabilityCmd = &cobra.Command{
	Use:   "availability",
	Short: "Check availability compliance for RDS instances",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatalf("Unable to load AWS SDK config, %v", err)
		}

		checkRDSMultiAZ(ctx, cfg)
	},
}

func checkRDSMultiAZ(ctx context.Context, cfg aws.Config) {
	rdsClient := rds.NewFromConfig(cfg)

	// List all RDS instances
	paginator := rds.NewDescribeDBInstancesPaginator(rdsClient, &rds.DescribeDBInstancesInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			log.Fatalf("Unable to describe RDS instances: %v", err)
		}
		for _, instance := range page.DBInstances {
			// Check if the RDS instance is Multi-AZ
			if instance.MultiAZ {
				fmt.Printf("RDS instance %s is deployed in Multi-AZ configuration.\n", *instance.DBInstanceIdentifier)
			} else {
				fmt.Printf("RDS instance %s is not deployed in Multi-AZ configuration.\n", *instance.DBInstanceIdentifier)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(availabilityCmd)
}
