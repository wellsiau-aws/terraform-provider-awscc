
---
page_title: "awscc_imagebuilder_workflow Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::ImageBuilder::Workflow
---

# awscc_imagebuilder_workflow (Resource)

Resource schema for AWS::ImageBuilder::Workflow

## Example Usage

### Create Image Builder Test Workflow

Creates an AWS Image Builder workflow with TEST type configuration, including a simple wait-for-action step and version control with basic tagging structure.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Example Image Builder workflow with minimal required attributes
resource "awscc_imagebuilder_workflow" "example" {
  name    = "example-workflow"
  type    = "TEST" # Can be BUILD, TEST, or DISTRIBUTE
  version = "1.0.0"

  data = jsonencode({
    name          = "example-workflow"
    schemaVersion = 1.0
    steps = [
      {
        name   = "test-step"
        action = "WaitForAction"
      }
    ]
  })

  tags = {
    Environment = "test"
    Modified_By = "AWSCC"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the workflow.
- `type` (String) The type of the workflow denotes whether the workflow is used to build, test, or distribute.
- `version` (String) The version of the workflow.

### Optional

- `change_description` (String) The change description of the workflow.
- `data` (String) The data of the workflow.
- `description` (String) The description of the workflow.
- `kms_key_id` (String) The KMS key identifier used to encrypt the workflow.
- `tags` (Map of String) The tags associated with the workflow.
- `uri` (String) The uri of the workflow.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the workflow.
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_imagebuilder_workflow.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_imagebuilder_workflow.example "arn"
```
