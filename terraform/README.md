# Terraform Configuration for SOC2 Compliance

This Terraform configuration sets up AWS resources to demonstrate SOC2 compliant and non-compliant configurations.

## Requirements

- Terraform 0.14 or newer
- AWS Account and AWS CLI configured with access key and secret key

## Setup

1. **AWS Account ID**: Ensure your AWS account ID is set as a variable or hardcoded in the KMS policy document.
2. **VPC and Subnets**: Specify your VPC and subnet IDs for the RDS cluster in `compliant.tf`.
3. **Security**: Adjust the ingress CIDR blocks in the RDS security group as needed for your environment.

Get your AWS account ID using the AWS CLI:

```bash
aws sts get-caller-identity --query "Account" --output text
```

Get your AWS region using the AWS CLI:

```bash
aws configure get region
```

Get your VPC and subnet IDs using the AWS CLI:

```bash
aws ec2 describe-vpcs --query "Vpcs[].VpcId" --output text # or
aws ec2 describe-vpcs --region ap-southeast-2 --query "Vpcs[].VpcId" --output text
```

Get your subnet IDs using the AWS CLI:

```bash
aws ec2 describe-subnets --filters "Name=vpc-id,Values=<VPC-ID>" --query "Subnets[].SubnetId" --output text # or
aws ec2 describe-subnets --region ap-southeast-2 --filters "Name=vpc-id,Values=<VPC-ID>" --query "Subnets[].SubnetId" --output text
```


## Variables

- `aws_account_id` (required): Your AWS account ID for the KMS key policy.
- `master_username` (required): Master username for the RDS cluster.
- `master_password` (required): Master password for the RDS cluster. Use a secure method to set this variable.

## Running Terraform

Initialize Terraform:

````bash
terraform init

Plan the deployment:

```bash
terraform plan
````

Apply the configuration:

```bash
terraform apply
```

To destroy the resources:

```bash
terraform destroy
```

## Compliance Checks

- `compliant.tf` includes configurations for a private S3 bucket with server-side encryption and a PostgreSQL RDS instance with encryption at rest in a multi-AZ deployment.
- `nonCompliant.tf` creates a publicly readable S3 bucket.

Remember to destroy the resources after your checks to avoid unnecessary costs.
