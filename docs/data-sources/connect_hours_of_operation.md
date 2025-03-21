---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_hours_of_operation Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::HoursOfOperation
---

# awscc_connect_hours_of_operation (Data Source)

Data Source schema for AWS::Connect::HoursOfOperation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `config` (Attributes Set) Configuration information for the hours of operation: day, start time, and end time. (see [below for nested schema](#nestedatt--config))
- `description` (String) The description of the hours of operation.
- `hours_of_operation_arn` (String) The Amazon Resource Name (ARN) for the hours of operation.
- `hours_of_operation_overrides` (Attributes List) One or more hours of operation overrides assigned to an hour of operation. (see [below for nested schema](#nestedatt--hours_of_operation_overrides))
- `instance_arn` (String) The identifier of the Amazon Connect instance.
- `name` (String) The name of the hours of operation.
- `tags` (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))
- `time_zone` (String) The time zone of the hours of operation.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Read-Only:

- `day` (String) The day that the hours of operation applies to.
- `end_time` (Attributes) The end time that your contact center closes. (see [below for nested schema](#nestedatt--config--end_time))
- `start_time` (Attributes) The start time that your contact center opens. (see [below for nested schema](#nestedatt--config--start_time))

<a id="nestedatt--config--end_time"></a>
### Nested Schema for `config.end_time`

Read-Only:

- `hours` (Number) The hours.
- `minutes` (Number) The minutes.


<a id="nestedatt--config--start_time"></a>
### Nested Schema for `config.start_time`

Read-Only:

- `hours` (Number) The hours.
- `minutes` (Number) The minutes.



<a id="nestedatt--hours_of_operation_overrides"></a>
### Nested Schema for `hours_of_operation_overrides`

Read-Only:

- `effective_from` (String) The date from which the hours of operation override would be effective.
- `effective_till` (String) The date till which the hours of operation override would be effective.
- `hours_of_operation_override_id` (String) The Resource Identifier for the hours of operation override.
- `override_config` (Attributes Set) Configuration information for the hours of operation override: day, start time, and end time. (see [below for nested schema](#nestedatt--hours_of_operation_overrides--override_config))
- `override_description` (String) The description of the hours of operation override.
- `override_name` (String) The name of the hours of operation override.

<a id="nestedatt--hours_of_operation_overrides--override_config"></a>
### Nested Schema for `hours_of_operation_overrides.override_config`

Read-Only:

- `day` (String) The day that the hours of operation override applies to.
- `end_time` (Attributes) The new end time that your contact center closes for the overriden days. (see [below for nested schema](#nestedatt--hours_of_operation_overrides--override_config--end_time))
- `start_time` (Attributes) The new start time that your contact center opens for the overriden days. (see [below for nested schema](#nestedatt--hours_of_operation_overrides--override_config--start_time))

<a id="nestedatt--hours_of_operation_overrides--override_config--end_time"></a>
### Nested Schema for `hours_of_operation_overrides.override_config.end_time`

Read-Only:

- `hours` (Number) The hours.
- `minutes` (Number) The minutes.


<a id="nestedatt--hours_of_operation_overrides--override_config--start_time"></a>
### Nested Schema for `hours_of_operation_overrides.override_config.start_time`

Read-Only:

- `hours` (Number) The hours.
- `minutes` (Number) The minutes.




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
