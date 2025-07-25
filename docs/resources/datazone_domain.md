---
page_title: "awscc_datazone_domain Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A domain is an organizing entity for connecting together assets, users, and their projects
---

# awscc_datazone_domain (Resource)

A domain is an organizing entity for connecting together assets, users, and their projects

## Example Usage

Creates an Amazon DataZone domain.

```terraform
resource "awscc_datazone_domain" "example" {
  name                  = "example"
  domain_execution_role = awscc_iam_role.example.arn
  description           = "Datazone domain example"


  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "example" {
  path = "/service-role/"
  assume_role_policy_document = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "datazone.amazonaws.com"
        },
        "Action" : [
          "sts:AssumeRole",
          "sts:TagSession"
        ],
        "Condition" : {
          "StringEquals" : {
            "aws:SourceAccount" : var.source_account_id
          },
          "ForAllValues:StringLike" : {
            "aws:TagKeys" : "datazone*"
          }
        }
      }
    ]
  })
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonDataZoneDomainExecutionRolePolicy"]
}

variable "source_account_id" {
  type        = string
  description = "Source AWS account id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_execution_role` (String) The domain execution role that is created when an Amazon DataZone domain is created. The domain execution role is created in the AWS account that houses the Amazon DataZone domain.
- `name` (String) The name of the Amazon DataZone domain.

### Optional

- `description` (String) The description of the Amazon DataZone domain.
- `domain_version` (String) The version of the domain.
- `kms_key_identifier` (String) The identifier of the AWS Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data.
- `service_role` (String) The service role of the domain that is created.
- `single_sign_on` (Attributes) The single-sign on configuration of the Amazon DataZone domain. (see [below for nested schema](#nestedatt--single_sign_on))
- `tags` (Attributes Set) The tags specified for the Amazon DataZone domain. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The ARN of the Amazon DataZone domain.
- `created_at` (String) The timestamp of when the Amazon DataZone domain was last updated.
- `domain_id` (String) The id of the Amazon DataZone domain.
- `id` (String) Uniquely identifies the resource.
- `last_updated_at` (String) The timestamp of when the Amazon DataZone domain was last updated.
- `managed_account_id` (String) The identifier of the AWS account that manages the domain.
- `portal_url` (String) The URL of the data portal for this Amazon DataZone domain.
- `root_domain_unit_id` (String) The ID of the root domain in Amazon Datazone.
- `status` (String) The status of the Amazon DataZone domain.

<a id="nestedatt--single_sign_on"></a>
### Nested Schema for `single_sign_on`

Optional:

- `idc_instance_arn` (String) The ARN of the AWS Identity Center instance.
- `type` (String) The type of single sign-on in Amazon DataZone.
- `user_assignment` (String) The single sign-on user assignment in Amazon DataZone.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag.
- `value` (String) The value for the tag.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_datazone_domain.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_datazone_domain.example "id"
```