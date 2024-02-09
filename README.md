# CompliSight SOC2 Compliance Checker CLI Tool

**CompliSight**: Revolutionizing Compliance Management with Automation and Insight

![Complisight](https://github.com/loftwah/complisight/assets/19922556/ceba8d0f-c638-4e2d-bee4-03528a4c3ade)

In an era where digital transformation accelerates, regulatory landscapes evolve, and the stakes for information security compliance have never been higher, CompliSight emerges as the game-changer. Our cutting-edge platform is designed to demystify and automate the complex world of compliance, starting with SOC2 and expanding to a comprehensive suite of information security standards.

CompliSight leverages the latest in automation technology, cloud computing, and AI to offer a seamless, intuitive experience that simplifies compliance checks, reporting, and management across multiple frameworks. With CompliSight, businesses can effortlessly navigate the intricacies of SOC2, GDPR, HIPAA, and beyond, ensuring that compliance is no longer a bottleneck but a catalyst for growth and trust.

Our mission is clear: to empower organizations of all sizes to achieve and maintain compliance with confidence, efficiency, and unparalleled insight. Whether you're a startup navigating SOC2 for the first time or a multinational seeking to streamline your compliance processes, CompliSight is your partner in compliance excellence.

### Enhanced AWS Resource Configuration Solutions & SOC2 Considerations

#### S3 Bucket Configurations

- **Policies**: Implement strict access policies using the `aws_s3_bucket_policy` resource in Terraform or AWS CLI commands. Ensure policies enforce encryption with `aws:kms` or AES-256 and restrict public access completely.
- **Logging**: Enable access logging on all S3 buckets to track requests, using the AWS Management Console or the AWS CLI. This aids in auditing and monitoring activities.
- **Versioning**: Activate versioning on S3 buckets to protect against unintended deletions or overwrites. Use the AWS Console or CLI to enable this feature, providing a rollback mechanism.

#### AWS Service Configurations

- **CloudFormation**: Utilize CloudFormation templates to manage infrastructure, ensuring that all resources are defined in code for repeatable deployments. Validate templates against best practices using AWS CloudFormation Linter.
- **CloudTrail**: Guarantee that CloudTrail is enabled in all regions and configured to log to a central S3 bucket. Regularly review the trails for coverage and ensure logs are encrypted.
- **CloudWatch Alarms**: Set up alarms for high usage or unusual activity patterns to identify potential security incidents or system health issues. Use Amazon SNS to notify administrators.
- **Config Conformance Packs**: Deploy AWS Config rules and conformance packs to automatically audit configurations against SOC2-related AWS best practices.
- **EC2 Instances**: Regularly audit instance security groups for minimal necessary access ports. Use IAM roles for EC2 to provide necessary permissions without using static credentials.
- **ElastiCache**: Enable encryption in transit and at rest for ElastiCache to secure sensitive data, following AWS documentation for your cache engine.
- **GuardDuty**: Monitor and triage GuardDuty findings. Set up automated responses for common findings to enhance security posture.
- **IAM Roles/Users**: Use the principle of least privilege for IAM roles and users. Regularly review permissions and use AWS IAM Access Advisor to prune unnecessary permissions.
- **KMS Keys**: Implement automated rotation for KMS keys where possible and audit usage. Ensure that keys are used appropriately for encrypting data in other services.
- **Lambda**: Secure Lambda functions by restricting their execution role permissions, enabling CloudWatch Logs for monitoring, and using environment variables for sensitive data (encrypted at rest).
- **RDS Instances**: Ensure databases are encrypted at rest using AWS KMS. Enable automated backups and database logging for auditing purposes.
- **Secrets Manager**: Automate the rotation of secrets stored in Secrets Manager and ensure that access to secrets is logged and monitored.
- **VPC Configurations**: Secure VPCs by implementing private subnets for database and application servers, using security groups and NACLs effectively, and enabling VPC Flow Logs for network traffic analysis.

#### Expanded AWS Checks with CompliSight

- **Elastic Load Balancing (ELB)**: Verify HTTPS listeners and strong cipher usage for data in transit security.
- **Amazon ECS & ECR**: Confirm ECS task roles and ECR image scanning for vulnerabilities.
- **SNS & SQS**: Ensure server-side encryption and access policies are in place for messaging services.
- **AWS Kinesis**: Check for KMS-based encryption to secure data streams at rest.
- **DynamoDB**: Validate encryption at rest and access control for stored data.
- **AWS Redshift**: Assess encryption and access restrictions for data warehousing services.
- **CloudWatch Log Groups**: Confirm encryption with KMS keys and access control.

### GitHub Integration and SOC2-Related Checks

- **Repository Security**: Implement branch protection rules and enforce code reviews.
- **Secret Scanning**: Activate GitHub Advanced Security for detecting exposed secrets.
- **Dependency Scanning**: Use Dependabot or GitHub Actions for scanning and updating vulnerable dependencies.
- **Code Analysis**: Apply static code analysis via GitHub Actions to catch security issues pre-merge.
- **Access Controls & Audit Logs**: Regularly review access levels and monitor audit logs for unusual activities.

### Operational Best Practices for SOC2 Compliance

- **Regular Compliance Audits**: Utilize AWS Audit Manager and third-party tools for compliance checks.
- **Employee Training**: Implement ongoing training programs on security awareness and SOC2 compliance.
- **Documentation**: Maintain up-to-date documentation of policies, procedures, and AWS configurations for SOC2 compliance.
- **Vendor Management**: Conduct risk assessments on third-party services to ensure they meet SOC2 standards.

By leveraging CompliSight for comprehensive SOC2 compliance management, organizations can ensure a robust security posture, streamline compliance processes, and foster trust with clients and stakeholders. Continuous monitoring, regular audits, and adherence to best practices are pivotal in maintaining compliance and securing sensitive data across all operational environments.

## Phase 1: Initial SOC2 Compliance Data Collection

The `phase1.sh` script is a comprehensive tool designed for the initial collection of AWS resource configurations, aiding in SOC2 compliance assessments. This script automates the retrieval of configurations and settings from a wide range of AWS services, ensuring a thorough review of the environment against SOC2 requirements.

### How It Works

Upon execution, `phase1.sh` performs the following actions:

1. **Directory Setup**: Creates a `./data` directory to store the output of collected AWS service data.

2. **Logging**: Establishes a logging mechanism to capture both stdout and stderr, saving the output to a timestamped log file within the `./data` directory.

3. **Data Collection**: Utilizes AWS CLI commands to collect data across multiple AWS services, including IAM, EC2, S3, RDS, ElastiCache, KMS, VPC, CloudTrail, CloudFormation, Config, CloudWatch, Lambda, SecretsManager, ELB, ECS, ECR, SNS, SQS, Kinesis, DynamoDB, and Redshift.

   * **Enhanced S3 Data Collection**: Employs parallel processing to efficiently collect policy, versioning, and logging configurations for each S3 bucket.
   * **ECS Cluster Checks**: Improved to handle multiple clusters, collecting service information for each identified ECS cluster.
   * **GuardDuty Detector Checks**: Verifies the existence of GuardDuty detectors and collects findings if available.

### Usage

To run the script:

1. Ensure you have the AWS CLI installed and configured with appropriate access permissions.
2. Navigate to the directory containing `phase1.sh`.
3. Make the script executable: `chmod +x phase1.sh`.
4. Execute the script: `./phase1.sh`.
   * Optionally, specify AWS regions as arguments if you wish to collect data from specific regions only. `./phase.sh ap-southeast-2 us-west-1 eu-west1`.

### Output

* The script outputs data into JSON files within the `./data` directory, organizing them by service and AWS region.
* A detailed log file captures the execution process, including any errors encountered during data collection.

### Advantages

* **Automation**: Streamlines the initial data collection process, saving time and reducing manual effort.
* **Comprehensive Coverage**: Ensures a wide range of AWS services are audited, facilitating a thorough SOC2 compliance assessment.
* **Flexibility**: Allows for targeted data collection across specified AWS regions.

This script is an essential part of the SOC2 compliance toolkit, offering a solid foundation for subsequent analysis and remediation efforts.
