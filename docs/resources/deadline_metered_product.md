---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_deadline_metered_product Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Deadline::MeteredProduct Resource Type
---

# awscc_deadline_metered_product (Resource)

Definition of AWS::Deadline::MeteredProduct Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `family` (String)
- `license_endpoint_id` (String)
- `port` (Number)
- `product_id` (String)
- `vendor` (String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_deadline_metered_product.example <resource ID>
```
