
---
page_title: "awscc_iot_thing_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IoT::ThingGroup
---

# awscc_iot_thing_group (Resource)

Resource Type definition for AWS::IoT::ThingGroup

## Example Usage

### Create IoT Thing Group with Parent-Child Hierarchy

To create an IoT Thing Group with customizable attributes and an optional nested child group structure, allowing for hierarchical organization of IoT devices.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Data sources for AWS account and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Main IoT Thing Group
resource "awscc_iot_thing_group" "example" {
  thing_group_name = "example-thing-group"

  thing_group_properties = {
    thing_group_description = "Example IoT Thing Group created by AWSCC provider"
    attribute_payload = {
      attributes = {
        "Environment" = "Test"
        "Project"     = "Demo"
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Optional nested group example
resource "awscc_iot_thing_group" "child_group" {
  thing_group_name  = "example-child-thing-group"
  parent_group_name = awscc_iot_thing_group.example.thing_group_name

  thing_group_properties = {
    thing_group_description = "Child IoT Thing Group example"
    attribute_payload = {
      attributes = {
        "Environment" = "Test"
        "Parent"      = awscc_iot_thing_group.example.thing_group_name
      }
    }
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parent_group_name` (String)
- `query_string` (String)
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `thing_group_name` (String)
- `thing_group_properties` (Attributes) (see [below for nested schema](#nestedatt--thing_group_properties))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `thing_group_id` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--thing_group_properties"></a>
### Nested Schema for `thing_group_properties`

Optional:

- `attribute_payload` (Attributes) (see [below for nested schema](#nestedatt--thing_group_properties--attribute_payload))
- `thing_group_description` (String)

<a id="nestedatt--thing_group_properties--attribute_payload"></a>
### Nested Schema for `thing_group_properties.attribute_payload`

Optional:

- `attributes` (Map of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iot_thing_group.example "thing_group_name"
```
