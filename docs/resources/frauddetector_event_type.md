
---
page_title: "awscc_frauddetector_event_type Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A resource schema for an EventType in Amazon Fraud Detector.
---

# awscc_frauddetector_event_type (Resource)

A resource schema for an EventType in Amazon Fraud Detector.

## Example Usage

### Configure Fraud Detector Event Type

Creates an Amazon Fraud Detector event type with customer entity type, IP address and email variables, and defines fraud/legitimate outcome labels for transaction monitoring.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a Fraud Detector Event Type
resource "awscc_frauddetector_event_type" "example" {
  name        = "example-event-type"
  description = "Example Fraud Detector Event Type"

  # Entity types define the actors involved in the event
  entity_types = [
    {
      name        = "customer"
      description = "Customer entity type"
      inline      = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    }
  ]

  # Event variables define the data points collected for the event
  event_variables = [
    {
      name          = "ip_address"
      data_type     = "STRING"
      data_source   = "EVENT"
      description   = "IP address of the customer"
      variable_type = "IP_ADDRESS"
      default_value = "0.0.0.0"
      inline        = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    },
    {
      name          = "email_address"
      data_type     = "STRING"
      data_source   = "EVENT"
      description   = "Email address of the customer"
      variable_type = "EMAIL_ADDRESS"
      default_value = "example@example.com"
      inline        = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    }
  ]

  # Labels define the possible outcomes of fraud detection
  labels = [
    {
      name        = "fraud"
      description = "Fraudulent transaction"
      inline      = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    },
    {
      name        = "legitimate"
      description = "Legitimate transaction"
      inline      = true
      tags = [{
        key   = "Modified By"
        value = "AWSCC"
      }]
    }
  ]

  # Add tags to the event type
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entity_types` (Attributes List) (see [below for nested schema](#nestedatt--entity_types))
- `event_variables` (Attributes List) (see [below for nested schema](#nestedatt--event_variables))
- `labels` (Attributes List) (see [below for nested schema](#nestedatt--labels))
- `name` (String) The name for the event type

### Optional

- `description` (String) The description of the event type.
- `tags` (Attributes List) Tags associated with this event type. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The ARN of the event type.
- `created_time` (String) The time when the event type was created.
- `id` (String) Uniquely identifies the resource.
- `last_updated_time` (String) The time when the event type was last updated.

<a id="nestedatt--entity_types"></a>
### Nested Schema for `entity_types`

Optional:

- `arn` (String)
- `created_time` (String) The time when the event type was created.
- `description` (String) The description.
- `inline` (Boolean)
- `last_updated_time` (String) The time when the event type was last updated.
- `name` (String)
- `tags` (Attributes List) Tags associated with this event type. (see [below for nested schema](#nestedatt--entity_types--tags))

<a id="nestedatt--entity_types--tags"></a>
### Nested Schema for `entity_types.tags`

Optional:

- `key` (String)
- `value` (String)



<a id="nestedatt--event_variables"></a>
### Nested Schema for `event_variables`

Optional:

- `arn` (String)
- `created_time` (String) The time when the event type was created.
- `data_source` (String)
- `data_type` (String)
- `default_value` (String)
- `description` (String) The description.
- `inline` (Boolean)
- `last_updated_time` (String) The time when the event type was last updated.
- `name` (String)
- `tags` (Attributes List) Tags associated with this event type. (see [below for nested schema](#nestedatt--event_variables--tags))
- `variable_type` (String)

<a id="nestedatt--event_variables--tags"></a>
### Nested Schema for `event_variables.tags`

Optional:

- `key` (String)
- `value` (String)



<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Optional:

- `arn` (String)
- `created_time` (String) The time when the event type was created.
- `description` (String) The description.
- `inline` (Boolean)
- `last_updated_time` (String) The time when the event type was last updated.
- `name` (String)
- `tags` (Attributes List) Tags associated with this event type. (see [below for nested schema](#nestedatt--labels--tags))

<a id="nestedatt--labels--tags"></a>
### Nested Schema for `labels.tags`

Optional:

- `key` (String)
- `value` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_frauddetector_event_type.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_frauddetector_event_type.example "arn"
```
