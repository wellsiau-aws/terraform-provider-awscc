---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_vpc_block_public_access_exclusion Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::VPCBlockPublicAccessExclusion
---

# awscc_ec2_vpc_block_public_access_exclusion (Data Source)

Data Source schema for AWS::EC2::VPCBlockPublicAccessExclusion



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `exclusion_id` (String) The ID of the exclusion
- `internet_gateway_exclusion_mode` (String) The desired Block Public Access Exclusion Mode for a specific VPC/Subnet.
- `subnet_id` (String) The ID of the subnet. Required only if you don't specify VpcId
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `vpc_id` (String) The ID of the vpc. Required only if you don't specify SubnetId.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.