package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

// RDSService wraps the AWS SDK's RDS client
type RDSService struct {
	Client *rds.Client
}

// NewRDSService creates a new instance of RDSService
func NewRDSService(ctx context.Context) (*RDSService, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &RDSService{
		Client: rds.NewFromConfig(cfg),
	}, nil
}

// DescribeDBInstances retrieves details about RDS instances
func (s *RDSService) DescribeDBInstances(ctx context.Context) ([]rds.DBInstance, error) {
	paginator := rds.NewDescribeDBInstancesPaginator(s.Client, &rds.DescribeDBInstancesInput{})
	var dbInstances []rds.DBInstance
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to describe RDS instances, %v", err)
		}
		dbInstances = append(dbInstances, output.DBInstances...)
	}

	return dbInstances, nil
}
