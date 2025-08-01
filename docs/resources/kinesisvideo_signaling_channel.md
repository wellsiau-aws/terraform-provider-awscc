
---
page_title: "awscc_kinesisvideo_signaling_channel Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type Definition for AWS::KinesisVideo::SignalingChannel
---

# awscc_kinesisvideo_signaling_channel (Resource)

Resource Type Definition for AWS::KinesisVideo::SignalingChannel

## Example Usage

### Create Kinesis Video Signaling Channel

Creates a single-master Kinesis Video Signaling Channel with a 60-second message TTL for real-time communication in video streaming applications.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Kinesis Video Signaling Channel
resource "awscc_kinesisvideo_signaling_channel" "example" {
  name                = "example-signaling-channel"
  type                = "SINGLE_MASTER"
  message_ttl_seconds = 60

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Output the ARN of the signaling channel
output "signaling_channel_arn" {
  value = awscc_kinesisvideo_signaling_channel.example.arn
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `message_ttl_seconds` (Number) The period of time a signaling channel retains undelivered messages before they are discarded.
- `name` (String) The name of the Kinesis Video Signaling Channel.
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `type` (String) The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the Kinesis Video Signaling Channel.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. Specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. Specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.  The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_kinesisvideo_signaling_channel.example
  id = "name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_kinesisvideo_signaling_channel.example "name"
```
