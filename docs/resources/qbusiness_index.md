---
page_title: "awscc_qbusiness_index Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::QBusiness::Index Resource Type
---

# awscc_qbusiness_index (Resource)

Definition of AWS::QBusiness::Index Resource Type

## Example Usage

### Create a Q Business Index and associate it with a Q Business application.

```terraform
resource "awscc_qbusiness_index" "example" {
  application_id = awscc_qbusiness_application.example.application_id
  display_name   = "example_q_index"
  description    = "Example QBusiness Index"
  type           = "ENTERPRISE"
  capacity_configuration = {
    units = 1
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String)
- `display_name` (String)

### Optional

- `capacity_configuration` (Attributes) (see [below for nested schema](#nestedatt--capacity_configuration))
- `description` (String)
- `document_attribute_configurations` (Attributes List) (see [below for nested schema](#nestedatt--document_attribute_configurations))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `type` (String)

### Read-Only

- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `index_arn` (String)
- `index_id` (String)
- `index_statistics` (Attributes) (see [below for nested schema](#nestedatt--index_statistics))
- `status` (String)
- `updated_at` (String)

<a id="nestedatt--capacity_configuration"></a>
### Nested Schema for `capacity_configuration`

Optional:

- `units` (Number)


<a id="nestedatt--document_attribute_configurations"></a>
### Nested Schema for `document_attribute_configurations`

Optional:

- `name` (String)
- `search` (String)
- `type` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--index_statistics"></a>
### Nested Schema for `index_statistics`

Read-Only:

- `text_document_statistics` (Attributes) (see [below for nested schema](#nestedatt--index_statistics--text_document_statistics))

<a id="nestedatt--index_statistics--text_document_statistics"></a>
### Nested Schema for `index_statistics.text_document_statistics`

Read-Only:

- `indexed_text_bytes` (Number)
- `indexed_text_document_count` (Number)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_qbusiness_index.example
  id = "application_id|index_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_qbusiness_index.example "application_id|index_id"
```