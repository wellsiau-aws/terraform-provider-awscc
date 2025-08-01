
---
page_title: "awscc_identitystore_group_membership Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type Definition for AWS:IdentityStore::GroupMembership
---

# awscc_identitystore_group_membership (Resource)

Resource Type Definition for AWS:IdentityStore::GroupMembership

## Example Usage

### Identity Store Group Membership Assignment

This example demonstrates how to create a group membership in AWS IAM Identity Center (formerly AWS SSO) by assigning a user to a group, utilizing both the group and user resources as prerequisites.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get Identity Store ID
data "aws_ssoadmin_instances" "test" {}

# Create a Group
resource "awscc_identitystore_group" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = "TestGroup"
  description       = "Test Group for AWSCC example"
}

# Create a User in Identity Store
resource "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = "Test User"
  user_name         = "test.user"
  name {
    given_name  = "Test"
    family_name = "User"
  }
  emails {
    primary = true
    value   = "test.user@example.com"
  }
}

# Create Group Membership
resource "awscc_identitystore_group_membership" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  group_id          = awscc_identitystore_group.test.id
  member_id = {
    user_id = aws_identitystore_user.test.user_id
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The unique identifier for a group in the identity store.
- `identity_store_id` (String) The globally unique identifier for the identity store.
- `member_id` (Attributes) An object containing the identifier of a group member. (see [below for nested schema](#nestedatt--member_id))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `membership_id` (String) The identifier for a GroupMembership in the identity store.

<a id="nestedatt--member_id"></a>
### Nested Schema for `member_id`

Required:

- `user_id` (String) The identifier for a user in the identity store.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_identitystore_group_membership.example
  id = "membership_id|identity_store_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_identitystore_group_membership.example "membership_id|identity_store_id"
```
