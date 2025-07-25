---
page_title: "awscc_opensearchserverless_collection Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Amazon OpenSearchServerless collection resource
---

# awscc_opensearchserverless_collection (Resource)

Amazon OpenSearchServerless collection resource

## Example Usage

### Simple Collection
```terraform
# Create a Collection
resource "awscc_opensearchserverless_collection" "simple_collection" {
  name = "awscc-collection"
  depends_on = [
    awscc_opensearchserverless_security_policy.security_policy
  ]
}

# Encryption SecurityPolicy
resource "awscc_opensearchserverless_security_policy" "security_policy" {
  name        = "awscc-security-policy"
  description = "created via awscc"
  type        = "encryption"
  policy = jsonencode({
    "Rules" = [
      {
        "ResourceType" = "collection",
        "Resource" = [
          "collection/awscc-collection"
        ]
      }
    ],
    "AWSOwnedKey" = true
  })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the collection.

The name must meet the following criteria:
Unique to your account and AWS Region
Starts with a lowercase letter
Contains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)
Contains between 3 and 32 characters

### Optional

- `description` (String) The description of the collection
- `standby_replicas` (String) The possible standby replicas for the collection
- `tags` (Attributes List) List of tags to be added to the resource (see [below for nested schema](#nestedatt--tags))
- `type` (String) The possible types for the collection

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the collection.
- `collection_endpoint` (String) The endpoint for the collection.
- `collection_id` (String) The identifier of the collection
- `dashboard_endpoint` (String) The OpenSearch Dashboards endpoint for the collection.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key in the key-value pair
- `value` (String) The value in the key-value pair

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_opensearchserverless_collection.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_opensearchserverless_collection.example "id"
```