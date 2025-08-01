---
page_title: "awscc_vpclattice_resource_configuration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  VpcLattice ResourceConfiguration CFN resource
---

# awscc_vpclattice_resource_configuration (Resource)

VpcLattice ResourceConfiguration CFN resource

## Example Usage

```terraform
resource "awscc_vpclattice_resource_configuration" "example" {
  name                        = "example-resource-configuration"
  port_ranges                 = ["80"]
  protocol_type               = "TCP"
  resource_gateway_id         = awscc_vpclattice_resource_gateway.example.id
  resource_configuration_type = "SINGLE"
  resource_configuration_definition = {
    dns_resource = {
      domain_name     = "example.com"
      ip_address_type = "IPV4"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `resource_configuration_type` (String)

### Optional

- `allow_association_to_sharable_service_network` (Boolean)
- `port_ranges` (List of String)
- `protocol_type` (String)
- `resource_configuration_auth_type` (String)
- `resource_configuration_definition` (Attributes) (see [below for nested schema](#nestedatt--resource_configuration_definition))
- `resource_configuration_group_id` (String)
- `resource_gateway_id` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `resource_configuration_id` (String)

<a id="nestedatt--resource_configuration_definition"></a>
### Nested Schema for `resource_configuration_definition`

Optional:

- `arn_resource` (String)
- `dns_resource` (Attributes) (see [below for nested schema](#nestedatt--resource_configuration_definition--dns_resource))
- `ip_resource` (String)

<a id="nestedatt--resource_configuration_definition--dns_resource"></a>
### Nested Schema for `resource_configuration_definition.dns_resource`

Optional:

- `domain_name` (String)
- `ip_address_type` (String)



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
  to = awscc_vpclattice_resource_configuration.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_vpclattice_resource_configuration.example "arn"
```
