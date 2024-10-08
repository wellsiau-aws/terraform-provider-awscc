---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appintegrations_event_integration Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::AppIntegrations::EventIntegration
---

# awscc_appintegrations_event_integration (Resource)

Resource Type definition for AWS::AppIntegrations::EventIntegration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `event_bridge_bus` (String) The Amazon Eventbridge bus for the event integration.
- `event_filter` (Attributes) The EventFilter (source) associated with the event integration. (see [below for nested schema](#nestedatt--event_filter))
- `name` (String) The name of the event integration.

### Optional

- `description` (String) The event integration description.
- `tags` (Attributes List) The tags (keys and values) associated with the event integration. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `event_integration_arn` (String) The Amazon Resource Name (ARN) of the event integration.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--event_filter"></a>
### Nested Schema for `event_filter`

Required:

- `source` (String) The source of the events.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) A key to identify the tag.
- `value` (String) Corresponding tag value for the key.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_appintegrations_event_integration.example "name"
```
