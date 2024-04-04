---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datazone_environment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::DataZone::Environment Resource Type
---

# awscc_datazone_environment (Resource)

Definition of AWS::DataZone::Environment Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_identifier` (String) The identifier of the Amazon DataZone domain in which the environment would be created.
- `environment_profile_identifier` (String) The ID of the environment profile with which the Amazon DataZone environment would be created.
- `name` (String) The name of the environment.
- `project_identifier` (String) The ID of the Amazon DataZone project in which the environment would be created.

### Optional

- `description` (String) The description of the Amazon DataZone environment.
- `glossary_terms` (List of String) The glossary terms that can be used in the Amazon DataZone environment.
- `user_parameters` (Attributes List) The user parameters of the Amazon DataZone environment. (see [below for nested schema](#nestedatt--user_parameters))

### Read-Only

- `aws_account_id` (String) The AWS account in which the Amazon DataZone environment is created.
- `aws_account_region` (String) The AWS region in which the Amazon DataZone environment is created.
- `created_at` (String) The timestamp of when the environment was created.
- `created_by` (String) The Amazon DataZone user who created the environment.
- `domain_id` (String) The identifier of the Amazon DataZone domain in which the environment is created.
- `environment_blueprint_id` (String) The ID of the blueprint with which the Amazon DataZone environment was created.
- `environment_id` (String) The ID of the Amazon DataZone environment.
- `environment_profile_id` (String) The ID of the environment profile with which the Amazon DataZone environment was created.
- `id` (String) Uniquely identifies the resource.
- `project_id` (String) The ID of the Amazon DataZone project in which the environment is created.
- `provider_name` (String) The provider of the Amazon DataZone environment.
- `status` (String) The status of the Amazon DataZone environment.
- `updated_at` (String) The timestamp of when the environment was updated.

<a id="nestedatt--user_parameters"></a>
### Nested Schema for `user_parameters`

Optional:

- `name` (String) The name of an environment parameter.
- `value` (String) The value of an environment parameter.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_datazone_environment.example <resource ID>
```