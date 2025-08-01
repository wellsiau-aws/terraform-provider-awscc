
---
page_title: "awscc_ec2_vpc_peering_connection Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::VPCPeeringConnection
---

# awscc_ec2_vpc_peering_connection (Resource)

Resource Type definition for AWS::EC2::VPCPeeringConnection

## Example Usage

### VPC Peering Connection Setup

Creates a VPC peering connection between two VPCs (10.0.0.0/16 and 172.16.0.0/16) within the same AWS account and region, demonstrating the basic configuration for VPC peering including the necessary VPC resources.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create VPC 1 (Requester)
resource "awscc_ec2_vpc" "vpc1" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "VPC1-Requester"
  }]
}

# Create VPC 2 (Accepter)
resource "awscc_ec2_vpc" "vpc2" {
  cidr_block = "172.16.0.0/16"
  tags = [{
    key   = "Name"
    value = "VPC2-Accepter"
  }]
}

# Create VPC Peering Connection
resource "awscc_ec2_vpc_peering_connection" "example" {
  vpc_id        = awscc_ec2_vpc.vpc1.id
  peer_vpc_id   = awscc_ec2_vpc.vpc2.id
  peer_owner_id = data.aws_caller_identity.current.account_id
  peer_region   = data.aws_region.current.name
  tags = [{
    key   = "Name"
    value = "VPC-Peering-Example"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `peer_vpc_id` (String) The ID of the VPC with which you are creating the VPC peering connection. You must specify this parameter in the request.
- `vpc_id` (String) The ID of the VPC.

### Optional

- `peer_owner_id` (String) The AWS account ID of the owner of the accepter VPC.
- `peer_region` (String) The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request.
- `peer_role_arn` (String) The Amazon Resource Name (ARN) of the VPC peer role for the peering connection in another AWS account.
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `vpc_peering_connection_id` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_vpc_peering_connection.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_vpc_peering_connection.example "id"
```
