---
page_title: "awscc_inspectorv2_cis_scan_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  CIS Scan Configuration resource schema
---

# awscc_inspectorv2_cis_scan_configuration (Resource)

CIS Scan Configuration resource schema

## Example Usage

### Example with daily schedule set with target resource tags.

```terraform
resource "awscc_inspectorv2_cis_scan_configuration" "example" {
  scan_name = "example"
  schedule = {
    daily = {
      start_time = {
        time_of_day = "00:00"
        time_zone   = "UTC"
      }
    }
  }

  security_level = "LEVEL_1"
  targets = {
    # use SELF for current account to remove any drift as property transformation is not currently supported.
    account_ids = ["123456789012"]
    target_resource_tags = {
      "Modified By" = ["AWSCC"]
    }

  }

  tags = {
    "Modified By" = "AWSCC"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `scan_name` (String) Name of the scan
- `schedule` (Attributes) Choose a Schedule cadence (see [below for nested schema](#nestedatt--schedule))
- `security_level` (String)
- `targets` (Attributes) (see [below for nested schema](#nestedatt--targets))

### Optional

- `tags` (Map of String)

### Read-Only

- `arn` (String) CIS Scan configuration unique identifier
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--schedule"></a>
### Nested Schema for `schedule`

Optional:

- `daily` (Attributes) (see [below for nested schema](#nestedatt--schedule--daily))
- `monthly` (Attributes) (see [below for nested schema](#nestedatt--schedule--monthly))
- `one_time` (String)
- `weekly` (Attributes) (see [below for nested schema](#nestedatt--schedule--weekly))

<a id="nestedatt--schedule--daily"></a>
### Nested Schema for `schedule.daily`

Optional:

- `start_time` (Attributes) (see [below for nested schema](#nestedatt--schedule--daily--start_time))

<a id="nestedatt--schedule--daily--start_time"></a>
### Nested Schema for `schedule.daily.start_time`

Optional:

- `time_of_day` (String)
- `time_zone` (String)



<a id="nestedatt--schedule--monthly"></a>
### Nested Schema for `schedule.monthly`

Optional:

- `day` (String)
- `start_time` (Attributes) (see [below for nested schema](#nestedatt--schedule--monthly--start_time))

<a id="nestedatt--schedule--monthly--start_time"></a>
### Nested Schema for `schedule.monthly.start_time`

Optional:

- `time_of_day` (String)
- `time_zone` (String)



<a id="nestedatt--schedule--weekly"></a>
### Nested Schema for `schedule.weekly`

Optional:

- `days` (List of String)
- `start_time` (Attributes) (see [below for nested schema](#nestedatt--schedule--weekly--start_time))

<a id="nestedatt--schedule--weekly--start_time"></a>
### Nested Schema for `schedule.weekly.start_time`

Optional:

- `time_of_day` (String)
- `time_zone` (String)




<a id="nestedatt--targets"></a>
### Nested Schema for `targets`

Required:

- `account_ids` (List of String)
- `target_resource_tags` (Map of List of String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_inspectorv2_cis_scan_configuration.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_inspectorv2_cis_scan_configuration.example "arn"
```