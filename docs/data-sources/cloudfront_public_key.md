---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_public_key Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudFront::PublicKey
---

# awscc_cloudfront_public_key (Data Source)

Data Source schema for AWS::CloudFront::PublicKey



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `created_time` (String)
- `public_key_config` (Attributes) Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html). (see [below for nested schema](#nestedatt--public_key_config))
- `public_key_id` (String)

<a id="nestedatt--public_key_config"></a>
### Nested Schema for `public_key_config`

Read-Only:

- `caller_reference` (String) A string included in the request to help make sure that the request can't be replayed.
- `comment` (String) A comment to describe the public key. The comment cannot be longer than 128 characters.
- `encoded_key` (String) The public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).
- `name` (String) A name to help identify the public key.
