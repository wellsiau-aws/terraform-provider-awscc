---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_organizations_resource_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  You can use AWS::Organizations::ResourcePolicy to delegate policy management for AWS Organizations to specified member accounts to perform policy actions that are by default available only to the management account.
---

# awscc_organizations_resource_policy (Resource)

You can use AWS::Organizations::ResourcePolicy to delegate policy management for AWS Organizations to specified member accounts to perform policy actions that are by default available only to the management account.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `content` (String) The policy document. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.

### Optional

- `tags` (Attributes Set) A list of tags that you want to attach to the resource policy (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the resource policy.
- `id` (String) Uniquely identifies the resource.
- `resource_policy_id` (String) The unique identifier (ID) associated with this resource policy.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key identifier, or name, of the tag.
- `value` (String) The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_organizations_resource_policy.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_organizations_resource_policy.example "id"
```
