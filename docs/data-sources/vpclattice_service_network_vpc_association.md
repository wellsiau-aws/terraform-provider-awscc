---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_vpclattice_service_network_vpc_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::VpcLattice::ServiceNetworkVpcAssociation
---

# awscc_vpclattice_service_network_vpc_association (Data Source)

Data Source schema for AWS::VpcLattice::ServiceNetworkVpcAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `created_at` (String)
- `security_group_ids` (Set of String)
- `service_network_arn` (String)
- `service_network_id` (String)
- `service_network_identifier` (String)
- `service_network_name` (String)
- `service_network_vpc_association_id` (String)
- `status` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `vpc_id` (String)
- `vpc_identifier` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)