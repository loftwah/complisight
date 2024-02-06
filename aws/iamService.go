package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

// IAMService wraps the AWS SDK's IAM client
type IAMService struct {
	Client *iam.Client
}

// NewIAMService creates a new instance of IAMService
func NewIAMService(ctx context.Context) (*IAMService, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &IAMService{
		Client: iam.NewFromConfig(cfg),
	}, nil
}

// ListRoles lists all IAM roles in the account
func (s *IAMService) ListRoles(ctx context.Context) ([]string, error) {
	paginator := iam.NewListRolesPaginator(s.Client, &iam.ListRolesInput{})
	var roles []string
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to list IAM roles, %v", err)
		}
		for _, role := range output.Roles {
			roles = append(roles, aws.ToString(role.RoleName))
		}
	}

	return roles, nil
}
