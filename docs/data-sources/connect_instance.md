---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_instance Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::Instance
---

# awscc_connect_instance (Data Source)

Data Source schema for AWS::Connect::Instance



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) An instanceArn is automatically generated on creation based on instanceId.
- `attributes` (Attributes) The attributes for the instance. (see [below for nested schema](#nestedatt--attributes))
- `created_time` (String) Timestamp of instance creation logged as part of instance creation.
- `directory_id` (String) Existing directoryId user wants to map to the new Connect instance.
- `identity_management_type` (String) Specifies the type of directory integration for new instance.
- `instance_alias` (String) Alias of the new directory created as part of new instance creation.
- `instance_id` (String) An instanceId is automatically generated on creation and assigned as the unique identifier.
- `instance_status` (String) Specifies the creation status of new instance.
- `service_role` (String) Service linked role created as part of instance creation.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--attributes"></a>
### Nested Schema for `attributes`

Read-Only:

- `auto_resolve_best_voices` (Boolean) Boolean flag which enables AUTO_RESOLVE_BEST_VOICES on an instance.
- `contact_lens` (Boolean) Boolean flag which enables CONTACT_LENS on an instance.
- `contactflow_logs` (Boolean) Boolean flag which enables CONTACTFLOW_LOGS on an instance.
- `early_media` (Boolean) Boolean flag which enables EARLY_MEDIA on an instance.
- `enhanced_chat_monitoring` (Boolean) Boolean flag which enables ENHANCED_CHAT_MONITORING on an instance.
- `enhanced_contact_monitoring` (Boolean) Boolean flag which enables ENHANCED_CONTACT_MONITORING on an instance.
- `high_volume_out_bound` (Boolean) Boolean flag which enables HIGH_VOLUME_OUTBOUND on an instance.
- `inbound_calls` (Boolean) Mandatory element which enables inbound calls on new instance.
- `multi_party_chat_conference` (Boolean) Boolean flag which enables MULTI_PARTY_CHAT_CONFERENCE on an instance.
- `multi_party_conference` (Boolean) Boolean flag which enables MULTI_PARTY_CONFERENCE on an instance.
- `outbound_calls` (Boolean) Mandatory element which enables outbound calls on new instance.
- `use_custom_tts_voices` (Boolean) Boolean flag which enables USE_CUSTOM_TTS_VOICES on an instance.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
