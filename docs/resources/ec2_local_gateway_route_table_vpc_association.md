---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_local_gateway_route_table_vpc_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.
---

# awscc_ec2_local_gateway_route_table_vpc_association (Resource)

Resource Type definition for Local Gateway Route Table VPC Association which describes an association between a local gateway route table and a VPC.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `local_gateway_route_table_id` (String) The ID of the local gateway route table.
- `vpc_id` (String) The ID of the VPC.

### Optional

- `tags` (Attributes Set) The tags for the association. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `local_gateway_id` (String) The ID of the local gateway.
- `local_gateway_route_table_vpc_association_id` (String) The ID of the association.
- `state` (String) The state of the association.

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
  to = awscc_ec2_local_gateway_route_table_vpc_association.example
  id = "local_gateway_route_table_vpc_association_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_local_gateway_route_table_vpc_association.example "local_gateway_route_table_vpc_association_id"
```
