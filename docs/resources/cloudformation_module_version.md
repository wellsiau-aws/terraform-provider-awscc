---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudformation_module_version Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A module that has been registered in the CloudFormation registry.
---

# awscc_cloudformation_module_version (Resource)

A module that has been registered in the CloudFormation registry.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `module_name` (String) The name of the module being registered.

Recommended module naming pattern: company_or_organization::service::type::MODULE.
- `module_package` (String) The url to the S3 bucket containing the schema and template fragment for the module you want to register.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the module.
- `description` (String) The description of the registered module.
- `documentation_url` (String) The URL of a page providing detailed documentation for this module.
- `id` (String) Uniquely identifies the resource.
- `is_default_version` (Boolean) Indicator of whether this module version is the current default version
- `schema` (String) The schema defining input parameters to and resources generated by the module.
- `time_created` (String) The time that the specified module version was registered.
- `version_id` (String) The version ID of the module represented by this module instance.
- `visibility` (String) The scope at which the type is visible and usable in CloudFormation operations.

The only allowed value at present is:

PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cloudformation_module_version.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cloudformation_module_version.example "arn"
```
