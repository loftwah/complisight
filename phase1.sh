#!/bin/bash

# Configuration: Accept AWS region and profile from environment variables or script arguments
REGION=${1:-$(aws configure get region)}
PROFILE=${2:-$AWS_PROFILE}

# Check if REGION is provided or successfully retrieved from AWS configuration
if [ -z "$REGION" ]; then
    echo "No region specified and no default region found in AWS configuration."
    exit 1
fi

# Directory to store output
OUTPUT_DIR="./data"
mkdir -p "$OUTPUT_DIR"

# Function to retrieve and save AWS service data
collect_service_data() {
    local service=$1
    local command=$2
    local file_suffix=$3
    local output_file="${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${REGION}-${file_suffix}.json"

    # Check if file exists, append timestamp if it does
    if [ -f "$output_file" ]; then
        timestamp=$(date +%s)
        output_file="${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${REGION}-${file_suffix}-${timestamp}.json"
    fi

    echo "Collecting ${service} data..."
    if eval "$command --region ${REGION}" > "$output_file"; then
        echo "${service} data saved to ${output_file}"
    else
        echo "Failed to collect ${service} data"
    fi
}

# Set AWS_PROFILE if provided
if [ -n "$PROFILE" ]; then
    export AWS_PROFILE="$PROFILE"
fi

# Fetch AWS account ID
AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text --region $REGION)

# Insert service collection calls
# IAM
collect_service_data "IAM" "aws iam list-users" "iam-users"
collect_service_data "IAM" "aws iam list-roles" "iam-roles"

# EC2
collect_service_data "EC2" "aws ec2 describe-instances" "ec2-instances"

# S3
# Adjusted S3 data collection to handle buckets individually and check for existing files
aws s3api list-buckets --query 'Buckets[].Name' --output text --region $REGION | while read -r bucket; do
    # For bucket policy
    output_file="${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${REGION}-s3-bucket-policy-${bucket}.json"
    if [ -f "$output_file" ]; then
        timestamp=$(date +%s)
        output_file="${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${REGION}-s3-bucket-policy-${bucket}-${timestamp}.json"
    fi
    echo "Collecting S3 bucket policy for ${bucket}..."
    aws s3api get-bucket-policy --bucket "$bucket" --output json --region $REGION > "$output_file" 2>/dev/null || echo "Failed to collect S3 bucket policy for ${bucket}"

    # Repeat for bucket encryption and any other S3-related collections
done

# RDS
collect_service_data "RDS" "aws rds describe-db-instances" "rds-instances"

# ElastiCache
collect_service_data "ElastiCache" "aws elasticache describe-cache-clusters" "elasticache-clusters"

# KMS
collect_service_data "KMS" "aws kms list-keys" "kms-keys"

# VPC
collect_service_data "VPC" "aws ec2 describe-vpcs" "vpc"

# CloudTrail
collect_service_data "CloudTrail" "aws cloudtrail describe-trails" "cloudtrail"

# AWS CloudFormation
collect_service_data "CloudFormation" "aws cloudformation describe-stacks" "cloudformation-stacks"

# AWS Config
collect_service_data "Config" "aws configservice describe-configuration-recorders" "config-recorders"
collect_service_data "Config" "aws configservice describe-conformance-pack-status" "config-conformance-packs"

# Amazon CloudWatch
collect_service_data "CloudWatch" "aws cloudwatch describe-alarms" "cloudwatch-alarms"

# AWS Lambda
collect_service_data "Lambda" "aws lambda list-functions" "lambda-functions"

# AWS Secrets Manager
collect_service_data "SecretsManager" "aws secretsmanager list-secrets" "secretsmanager-secrets"

# Amazon GuardDuty
DETECTOR_ID=$(aws guardduty list-detectors --output text --query 'DetectorIds[0]' --region $REGION)
if [ ! -z "$DETECTOR_ID" ]; then
    collect_service_data "GuardDuty" "aws guardduty list-findings --detector-id $DETECTOR_ID" "guardduty-findings"
else
    echo "No GuardDuty detector found in region $REGION"
fi

# Ensure Route 53 data collection happens only once (global service)
if [ -z "$ROUTE53_COLLECTED" ]; then
    ROUTE53_COLLECTED=true
    collect_service_data "Route53" "aws route53 list-hosted-zones" "route53-hosted-zones"
    # Additional logic for Route 53 record sets if needed goes here
fi

echo "Data collection completed."
