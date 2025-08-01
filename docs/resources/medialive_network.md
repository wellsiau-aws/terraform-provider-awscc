
---
page_title: "awscc_medialive_network Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::MediaLive::Network.
---

# awscc_medialive_network (Resource)

Resource schema for AWS::MediaLive::Network.

## Example Usage

### MediaLive Network with IP Pools

Creates an AWS MediaLive network with two IP pools in different CIDR ranges and a default route configuration for media streaming infrastructure.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Example MediaLive Network
resource "awscc_medialive_network" "example" {
  name = "example-network"

  ip_pools = [
    {
      cidr = "10.0.0.0/24"
    },
    {
      cidr = "10.0.1.0/24"
    }
  ]

  routes = [
    {
      cidr    = "0.0.0.0/0"
      gateway = "10.0.0.1"
    }
  ]

  tags = [
    {
      key   = "Environment"
      value = "Production"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ip_pools` (Attributes List) The list of IP address cidr pools for the network (see [below for nested schema](#nestedatt--ip_pools))
- `name` (String) The user-specified name of the Network to be created.

### Optional

- `routes` (Attributes List) The routes for the network (see [below for nested schema](#nestedatt--routes))
- `tags` (Attributes List) A collection of key-value pairs. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The ARN of the Network.
- `associated_cluster_ids` (List of String)
- `id` (String) Uniquely identifies the resource.
- `network_id` (String) The unique ID of the Network.
- `state` (String) The current state of the Network.

<a id="nestedatt--ip_pools"></a>
### Nested Schema for `ip_pools`

Optional:

- `cidr` (String) IP address cidr pool


<a id="nestedatt--routes"></a>
### Nested Schema for `routes`

Optional:

- `cidr` (String) Ip address cidr
- `gateway` (String) IP address for the route packet paths


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
  to = awscc_medialive_network.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_medialive_network.example "id"
```
