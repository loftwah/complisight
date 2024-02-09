#!/bin/bash

# Directory to store output
OUTPUT_DIR="./data"
mkdir -p "$OUTPUT_DIR"

# Setup logging
LOG_FILE="${OUTPUT_DIR}/script_output_$(date +%Y-%m-%d_%H-%M-%S).log"

# Redirect stdout and stderr to log file
exec > >(tee "$LOG_FILE") 2>&1

echo "Starting SOC2 Compliance Data Collection..."

# Function to handle service data collection with error checking
collect_service_data() {
    local region=$1
    local service=$2
    local command=$3
    local file_suffix=$4
    local AWS_ACCOUNT_ID=$(aws sts get-caller-identity --query "Account" --output text --region "$region")
    local output_file="${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${region}-${file_suffix}.json"

    # Append timestamp if file exists
    if [ -f "$output_file" ]; then
        local timestamp=$(date +%Y-%m-%d_%H-%M-%S)
        output_file="${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${region}-${file_suffix}-${timestamp}.json"
    fi

    echo "Collecting ${service} data in ${region}..."
    if ! eval "$command --region ${region}" > "$output_file"; then
        echo "Error collecting ${service} data in ${region}. See log for details."
    else
        echo "${service} data saved to ${output_file}"
    fi
}

# Enhanced S3 data collection with parallel processing
collect_s3_data() {
    local region=$1
    local buckets=$(aws s3api list-buckets --query 'Buckets[].Name' --output text --region "$region")

    echo "Collecting S3 bucket data in ${region}..."
    for bucket in $buckets; do
        echo "Processing $bucket..."

        # Parallelize internal commands for efficiency
        {
            # Bucket Policy
            aws s3api get-bucket-policy --bucket "$bucket" --region "$region" > "${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${region}-s3-bucket-policy-${bucket}.json" 2>/dev/null || echo "No policy for $bucket"
            
            # Bucket Versioning
            aws s3api get-bucket-versioning --bucket "$bucket" --region "$region" > "${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${region}-s3-bucket-versioning-${bucket}.json"
            
            # Bucket Logging
            aws s3api get-bucket-logging --bucket "$bucket" --region "$region" > "${OUTPUT_DIR}/output-${AWS_ACCOUNT_ID}-${region}-s3-bucket-logging-${bucket}.json"
        } &
    done
    wait
}

# Replace the regions collection logic if necessary, or keep as is for simplicity
REGIONS="${@:-$(aws configure get region)}"
echo "Collecting data in regions: $REGIONS"

# Main collection loop - Enhanced with better structure and additional services
for REGION in $REGIONS; do
    echo "Processing region: $REGION"
    
    # Your existing and new service data collection calls
    collect_service_data "$REGION" "IAM" "aws iam list-users" "iam-users"
    collect_service_data "$REGION" "IAM" "aws iam list-roles" "iam-roles"
    collect_service_data "$REGION" "EC2" "aws ec2 describe-instances" "ec2-instances"
    collect_s3_data "$REGION"
    collect_service_data "$REGION" "RDS" "aws rds describe-db-instances" "rds-instances"
    collect_service_data "$REGION" "ElastiCache" "aws elasticache describe-cache-clusters" "elasticache-clusters"
    collect_service_data "$REGION" "KMS" "aws kms list-keys" "kms-keys"
    collect_service_data "$REGION" "VPC" "aws ec2 describe-vpcs" "vpc"
    collect_service_data "$REGION" "CloudTrail" "aws cloudtrail describe-trails" "cloudtrail"
    collect_service_data "$REGION" "CloudFormation" "aws cloudformation describe-stacks" "cloudformation-stacks"
    collect_service_data "$REGION" "Config" "aws configservice describe-configuration-recorders" "config-recorders"
    collect_service_data "$REGION" "Config" "aws configservice describe-conformance-pack-status" "config-conformance-packs"
    collect_service_data "$REGION" "CloudWatch" "aws cloudwatch describe-alarms" "cloudwatch-alarms"
    collect_service_data "$REGION" "Lambda" "aws lambda list-functions" "lambda-functions"
    collect_service_data "$REGION" "SecretsManager" "aws secretsmanager list-secrets" "secretsmanager-secrets"
    collect_service_data "$REGION" "ELB" "aws elb describe-load-balancers" "elb-load-balancers"
    # ECS cluster check improved for multiple clusters
    local ecs_clusters=$(aws ecs list-clusters --region "$REGION" --query "clusterArns[]" --output text)
    if [ -n "$ecs_clusters" ]; then
        for cluster_arn in $ecs_clusters; do
            collect_service_data "$REGION" "ECS" "aws ecs list-services --cluster $cluster_arn" "ecs-services-${cluster_arn##*/}"
        done
    else
        echo "No ECS clusters found in $REGION."
    fi
    collect_service_data "$REGION" "ECR" "aws ecr describe-repositories" "ecr-repositories"
    collect_service_data "$REGION" "SNS" "aws sns list-topics" "sns-topics"
    collect_service_data "$REGION" "SQS" "aws sqs list-queues" "sqs-queues"
    collect_service_data "$REGION" "Kinesis" "aws kinesis list-streams" "kinesis-streams"
    collect_service_data "$REGION" "DynamoDB" "aws dynamodb list-tables" "dynamodb-tables"
    collect_service_data "$REGION" "Redshift" "aws redshift describe-clusters" "redshift-clusters"
    collect_service_data "$REGION" "CloudWatch" "aws logs describe-log-groups" "cloudwatch-log-groups"

    # GuardDuty handling improved to check for detector existence
detector_id=$(aws guardduty list-detectors --region "$REGION" --query "DetectorIds[0]" --output text 2>/dev/null)
if [ -n "$detector_id" ]; then
    collect_service_data "$REGION" "GuardDuty" "aws guardduty list-findings --detector-id $detector_id" "guardduty-findings"
else
    echo "No GuardDuty detector found in $REGION."
fi
done

echo "Data collection completed."
