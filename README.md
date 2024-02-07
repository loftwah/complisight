# CompliSight SOC2 Compliance Checker CLI Tool

**CompliSight**: Revolutionizing Compliance Management with Automation and Insight

![Complisight](https://github.com/loftwah/complisight/assets/19922556/ceba8d0f-c638-4e2d-bee4-03528a4c3ade)

In an era where digital transformation accelerates, regulatory landscapes evolve, and the stakes for information security compliance have never been higher, CompliSight emerges as the game-changer. Our cutting-edge platform is designed to demystify and automate the complex world of compliance, starting with SOC2 and expanding to a comprehensive suite of information security standards.

CompliSight leverages the latest in automation technology, cloud computing, and AI to offer a seamless, intuitive experience that simplifies compliance checks, reporting, and management across multiple frameworks. With CompliSight, businesses can effortlessly navigate the intricacies of SOC2, GDPR, HIPAA, and beyond, ensuring that compliance is no longer a bottleneck but a catalyst for growth and trust.

Our mission is clear: to empower organizations of all sizes to achieve and maintain compliance with confidence, efficiency, and unparalleled insight. Whether you're a startup navigating SOC2 for the first time or a multinational seeking to streamline your compliance processes, CompliSight is your partner in compliance excellence.

**Key Features:**

* **Automated Compliance Checks**: Streamline your SOC2 compliance process with automated checks for AWS-hosted applications, ensuring your security, privacy, and integrity standards are met with precision.
* **Expansive Framework Support**: Beyond SOC2, CompliSight grows with your compliance needs, offering support for GDPR, HIPAA, ISO standards, and more.
* **Actionable Insights**: Our platform doesn't just identify compliance gaps; it provides clear, actionable guidance to address them, turning compliance challenges into opportunities for improvement.
* **Continuous Monitoring**: Stay ahead of compliance with continuous monitoring and real-time alerts, ensuring that your organization remains compliant amidst evolving regulations and infrastructure changes.

**Vision**: To capture the untapped potential of automated compliance tools, CompliSight is not just a product but a movement towards a future where compliance empowers businesses rather than encumbers them. We are on a mission to secure millions in value by making compliance accessible, manageable, and insightful for everyone, everywhere.

Join us as we pave the way to this future, one compliance check at a time.

This documentation outlines the SOC2 Compliance Checker CLI tool designed for assessing SOC2 compliance of AWS-hosted Ruby on Rails applications. Utilizing the AWS SDK for Go, this tool automates the process of checking against the SOC2 Trust Services Criteria: Security, Availability, Processing Integrity, Confidentiality, and Privacy.

## Features and Commands

### **1. Root Command**

* **Features:** Entry point for the application, offering help and general information.
* **Implementation:** Implemented in `root.go`, leveraging Cobra's help system.

### **2. Security Command**

* **Features:** Checks IAM policies, S3 bucket settings, and encryption status of RDS and ElastiCache instances.
* **Implementation:** Resides in `security.go`, utilizing the AWS SDK for interactions.

### **3. Privacy Command**

* **Features:** Evaluates data handling and encryption practices for personal information.
* **Implementation:** Found in `privacy.go`, employing AWS SDK to query S3 and IAM configurations.

### **4. Integrity Command**

* **Features:** Monitors unauthorized changes and validates data processing accuracy.
* **Implementation:** Implemented in `integrity.go`, using AWS CloudTrail and RDS logs.

### **5. Confidentiality Command**

* **Features:** Ensures encryption and access control for confidential data.
* **Implementation:** Resides in `confidentiality.go`, checking S3, RDS, and IAM settings.

### **6. Availability Command**

* **Features:** Assesses system availability and sets up health monitoring.
* **Implementation:** Found in `availability.go`, inspecting AWS service configurations.

### **7. Assess Command**

* **Features:** Runs all checks sequentially for a comprehensive compliance assessment.
* **Implementation:** Potentially in `assess.go`, orchestrating calls to individual checks.

## Usage

### **Configure AWS Credentials**

The AWS SDK for Go requires credentials (AWS access key ID and secret access key) to interact with AWS services. You can configure credentials in several ways:

* **Environment Variables:** Set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`, and optionally, `AWS_SESSION_TOKEN`.

* **Shared Credentials File:** Typically located at `~/.aws/credentials` on Linux and macOS, or `%USERPROFILE%\.aws\credentials` on Windows.

* **IAM Role:** When running on AWS services such as EC2, ECS, or Lambda, you can assign an IAM role with appropriate permissions to the service.

## Reporting and SOC2 Report Types

### **Output and Reporting:**

* After performing checks, the tool outputs reports in a structured format. Consider JSON for machine parsing and detailed human-readable summaries for direct consumption.
* Reports should highlight compliance status, identify gaps, and suggest remediation steps.

### **Handling SOC2 Type 1 and Type 2 Reports:**

* **Type 1 Report:** Focuses on the suitability of the design of controls at a specific point in time. The tool should provide a snapshot report detailing the current compliance status across all checked criteria.

* **Type 2 Report:** Examines the operational effectiveness of controls over a defined period, typically at least six months. While the tool primarily provides instant checks (akin to Type 1), it can be adapted to support Type 2 by:

  * Integrating with AWS CloudTrail and Config to retrieve historical data.
  * Offering guidance on tracking changes and maintaining compliance over time.

### **General Structure and Approach:**

* **AWS SDK Integration:** Commands use the AWS SDK for Go to interact with AWS services, ensuring that checks are accurate and based on the latest AWS configurations.
* **Environment and Configuration:** Users can specify AWS regions and profiles via global flags or a configuration file, tailoring the checks to their specific AWS environment.
* **Continuous Compliance:** Encourage users to run the tool regularly or integrate it into their CI/CD pipeline for ongoing SOC2 compliance monitoring.

## Getting Started

* **Installation:** Instructions on installing the CLI tool, setting up AWS credentials, and configuring necessary permissions.
* **Usage:** Examples of common commands, including how to run individual checks and perform a comprehensive assessment with the `assess` command.
* **Contributing:** Guidelines for contributing to the tool, reporting issues, and suggesting enhancements.

## Conclusion

The SOC2 Compliance Checker CLI tool is designed to simplify the process of SOC2 compliance assessment for AWS-hosted applications, providing valuable insights into the security and compliance posture of Ruby on Rails applications. By automating the evaluation process and offering actionable reports, it aids organizations in maintaining SOC2 compliance efficiently.
