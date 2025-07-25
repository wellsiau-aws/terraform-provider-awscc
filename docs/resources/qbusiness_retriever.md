---
page_title: "awscc_qbusiness_retriever Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::QBusiness::Retriever Resource Type
---

# awscc_qbusiness_retriever (Resource)

Definition of AWS::QBusiness::Retriever Resource Type

## Example Usage

### Create a native retriever for Q Business application using the Q Business Index.

```terraform
resource "awscc_qbusiness_retriever" "example" {
  application_id = awscc_qbusiness_application.example.application_id
  display_name   = "example_q_retriever"
  type           = "NATIVE_INDEX"

  configuration = {
    native_index_configuration = {
      index_id = awscc_qbusiness_index.example.index_id
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

### Required

- `application_id` (String)
- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `display_name` (String)
- `type` (String)

### Optional

- `role_arn` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `retriever_arn` (String)
- `retriever_id` (String)
- `status` (String)
- `updated_at` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Optional:

- `kendra_index_configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration--kendra_index_configuration))
- `native_index_configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration--native_index_configuration))

<a id="nestedatt--configuration--kendra_index_configuration"></a>
### Nested Schema for `configuration.kendra_index_configuration`

Optional:

- `index_id` (String)


<a id="nestedatt--configuration--native_index_configuration"></a>
### Nested Schema for `configuration.native_index_configuration`

Optional:

- `index_id` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_qbusiness_retriever.example
  id = "application_id|retriever_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_qbusiness_retriever.example "application_id|retriever_id"
```