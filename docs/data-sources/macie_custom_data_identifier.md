---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_macie_custom_data_identifier Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Macie::CustomDataIdentifier
---

# awscc_macie_custom_data_identifier (Data Source)

Data Source schema for AWS::Macie::CustomDataIdentifier



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) Custom data identifier ARN.
- `custom_data_identifier_id` (String) Custom data identifier ID.
- `description` (String) Description of custom data identifier.
- `ignore_words` (List of String) Words to be ignored.
- `keywords` (List of String) Keywords to be matched against.
- `maximum_match_distance` (Number) Maximum match distance.
- `name` (String) Name of custom data identifier.
- `regex` (String) Regular expression for custom data identifier.
- `tags` (Attributes List) A collection of tags associated with a resource (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The tag's key.
- `value` (String) The tag's value.
