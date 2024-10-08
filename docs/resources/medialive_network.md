---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_medialive_network Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::MediaLive::Network.
---

# awscc_medialive_network (Resource)

Resource schema for AWS::MediaLive::Network.



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

```shell
$ terraform import awscc_medialive_network.example "id"
```
