
---
page_title: "awscc_mediaconnect_flow_vpc_interface Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::MediaConnect::FlowVpcInterface
---

# awscc_mediaconnect_flow_vpc_interface (Resource)

Resource schema for AWS::MediaConnect::FlowVpcInterface

## Example Usage

### MediaConnect Flow VPC Interface Configuration

Creates a VPC interface for AWS MediaConnect Flow with necessary networking components (VPC, subnet, security group) and IAM role permissions, enabling secure media content delivery through private networks.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_region" "current" {}

# VPC and Networking Resources
resource "awscc_ec2_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = [{
    key   = "Name"
    value = "mediaconnect-example-vpc"
  }]
}

resource "awscc_ec2_subnet" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "mediaconnect-example-subnet"
  }]
}

resource "awscc_ec2_security_group" "example" {
  vpc_id            = awscc_ec2_vpc.example.id
  group_description = "Security group for MediaConnect VPC Interface"
  tags = [{
    key   = "Name"
    value = "mediaconnect-example-sg"
  }]
}

# IAM Role for MediaConnect
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["mediaconnect.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "vpc_interface" {
  statement {
    effect = "Allow"
    actions = [
      "ec2:CreateNetworkInterface",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DescribeSubnets",
      "ec2:DeleteNetworkInterface"
    ]
    resources = ["*"]
  }
}

resource "awscc_iam_role" "mediaconnect" {
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.assume_role.json))
  description                 = "Role for MediaConnect VPC Interface"
  role_name                   = "MediaConnectVPCInterface"
  managed_policy_arns         = []
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role_policy" "mediaconnect" {
  policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.vpc_interface.json))
  policy_name     = "MediaConnectVPCInterfacePolicy"
  role_name       = awscc_iam_role.mediaconnect.role_name
}

# MediaConnect Flow
resource "awscc_mediaconnect_flow" "example" {
  name              = "example-flow"
  availability_zone = "${data.aws_region.current.name}a"
  source = {
    name               = "example-source"
    description        = "Example source"
    protocol           = "zixi-push"
    ingest_port        = 2088
    max_bit_rate       = 10000000
    min_latency        = 2000
    vpc_interface_name = "example-vpc-interface"
    whitelist_cidr     = "0.0.0.0/0"
  }
}

# MediaConnect Flow VPC Interface
resource "awscc_mediaconnect_flow_vpc_interface" "example" {
  flow_arn           = awscc_mediaconnect_flow.example.flow_arn
  name               = "example-vpc-interface"
  role_arn           = awscc_iam_role.mediaconnect.arn
  security_group_ids = [awscc_ec2_security_group.example.id]
  subnet_id          = awscc_ec2_subnet.example.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `flow_arn` (String) The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
- `name` (String) Immutable and has to be a unique against other VpcInterfaces in this Flow.
- `role_arn` (String) Role Arn MediaConnect can assume to create ENIs in customer's account.
- `security_group_ids` (List of String) Security Group IDs to be used on ENI.
- `subnet_id` (String) Subnet must be in the AZ of the Flow

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `network_interface_ids` (List of String) IDs of the network interfaces created in customer's account by MediaConnect.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_mediaconnect_flow_vpc_interface.example
  id = "flow_arn|name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_mediaconnect_flow_vpc_interface.example "flow_arn|name"
```
