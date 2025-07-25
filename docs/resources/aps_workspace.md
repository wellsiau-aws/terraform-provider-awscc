
---
page_title: "awscc_aps_workspace Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::APS::Workspace
---

# awscc_aps_workspace (Resource)

Resource Type definition for AWS::APS::Workspace

## Example Usage

### Amazon Prometheus Service Workspace with Logging and Encryption

Creates an Amazon Managed Service for Prometheus (APS) workspace with CloudWatch logging configuration and KMS encryption, utilizing custom KMS key for enhanced security and CloudWatch log group for workspace monitoring.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create CloudWatch log group for APS workspace logging
resource "awscc_logs_log_group" "aps_logs" {
  log_group_name = "/aws/aps/workspaces/example"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create KMS key for APS workspace encryption
resource "awscc_kms_key" "aps_key" {
  description = "KMS key for APS Workspace encryption"
  key_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow APS Service"
        Effect = "Allow"
        Principal = {
          Service = "aps.amazonaws.com"
        }
        Action = [
          "kms:Decrypt",
          "kms:GenerateDataKey"
        ]
        Resource = "*"
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the APS workspace
resource "awscc_aps_workspace" "example" {
  alias = "example-workspace"
  logging_configuration = {
    log_group_arn = awscc_logs_log_group.aps_logs.arn
  }
  kms_key_arn = awscc_kms_key.aps_key.arn
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `alert_manager_definition` (String) The AMP Workspace alert manager definition data
- `alias` (String) AMP Workspace alias.
- `kms_key_arn` (String) KMS Key ARN used to encrypt and decrypt AMP workspace data.
- `logging_configuration` (Attributes) Logging configuration (see [below for nested schema](#nestedatt--logging_configuration))
- `query_logging_configuration` (Attributes) Query logging configuration (see [below for nested schema](#nestedatt--query_logging_configuration))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `workspace_configuration` (Attributes) Workspace configuration (see [below for nested schema](#nestedatt--workspace_configuration))

### Read-Only

- `arn` (String) Workspace arn.
- `id` (String) Uniquely identifies the resource.
- `prometheus_endpoint` (String) AMP Workspace prometheus endpoint
- `workspace_id` (String) Required to identify a specific APS Workspace.

<a id="nestedatt--logging_configuration"></a>
### Nested Schema for `logging_configuration`

Optional:

- `log_group_arn` (String) CloudWatch log group ARN


<a id="nestedatt--query_logging_configuration"></a>
### Nested Schema for `query_logging_configuration`

Optional:

- `destinations` (Attributes List) The destinations configuration for query logging (see [below for nested schema](#nestedatt--query_logging_configuration--destinations))

<a id="nestedatt--query_logging_configuration--destinations"></a>
### Nested Schema for `query_logging_configuration.destinations`

Optional:

- `cloudwatch_logs` (Attributes) Represents a cloudwatch logs destination for query logging (see [below for nested schema](#nestedatt--query_logging_configuration--destinations--cloudwatch_logs))
- `filters` (Attributes) Filters for logging (see [below for nested schema](#nestedatt--query_logging_configuration--destinations--filters))

<a id="nestedatt--query_logging_configuration--destinations--cloudwatch_logs"></a>
### Nested Schema for `query_logging_configuration.destinations.cloudwatch_logs`

Optional:

- `log_group_arn` (String) The ARN of the CloudWatch Logs log group


<a id="nestedatt--query_logging_configuration--destinations--filters"></a>
### Nested Schema for `query_logging_configuration.destinations.filters`

Optional:

- `qsp_threshold` (Number) Query logs with QSP above this limit are vended




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--workspace_configuration"></a>
### Nested Schema for `workspace_configuration`

Optional:

- `limits_per_label_sets` (Attributes Set) An array of label set and associated limits (see [below for nested schema](#nestedatt--workspace_configuration--limits_per_label_sets))
- `retention_period_in_days` (Number) How many days that metrics are retained in the workspace

<a id="nestedatt--workspace_configuration--limits_per_label_sets"></a>
### Nested Schema for `workspace_configuration.limits_per_label_sets`

Optional:

- `label_set` (Attributes Set) An array of series labels (see [below for nested schema](#nestedatt--workspace_configuration--limits_per_label_sets--label_set))
- `limits` (Attributes) Limits that can be applied to a label set (see [below for nested schema](#nestedatt--workspace_configuration--limits_per_label_sets--limits))

<a id="nestedatt--workspace_configuration--limits_per_label_sets--label_set"></a>
### Nested Schema for `workspace_configuration.limits_per_label_sets.label_set`

Optional:

- `name` (String) Name of the label
- `value` (String) Value of the label


<a id="nestedatt--workspace_configuration--limits_per_label_sets--limits"></a>
### Nested Schema for `workspace_configuration.limits_per_label_sets.limits`

Optional:

- `max_series` (Number) The maximum number of active series that can be ingested for this label set

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_aps_workspace.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_aps_workspace.example "arn"
```
