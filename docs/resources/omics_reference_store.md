---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_omics_reference_store Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Omics::ReferenceStore Resource Type
---

# awscc_omics_reference_store (Resource)

Definition of AWS::Omics::ReferenceStore Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) A name for the store.

### Optional

- `description` (String) A description for the store.
- `sse_config` (Attributes) Server-side encryption (SSE) settings for a store. (see [below for nested schema](#nestedatt--sse_config))
- `tags` (Map of String)

### Read-Only

- `arn` (String) The store's ARN.
- `creation_time` (String) When the store was created.
- `id` (String) Uniquely identifies the resource.
- `reference_store_id` (String)

<a id="nestedatt--sse_config"></a>
### Nested Schema for `sse_config`

Optional:

- `key_arn` (String) An encryption key ARN.
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_omics_reference_store.example "reference_store_id"
```
