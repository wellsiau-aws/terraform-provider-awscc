---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datazone_environment_actions Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::DataZone::EnvironmentActions
---

# awscc_datazone_environment_actions (Data Source)

Data Source schema for AWS::DataZone::EnvironmentActions



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `description` (String) The description of the Amazon DataZone environment action.
- `domain_id` (String) The identifier of the Amazon DataZone domain in which the environment is created.
- `domain_identifier` (String) The identifier of the Amazon DataZone domain in which the environment would be created.
- `environment_actions_id` (String) The ID of the Amazon DataZone environment action.
- `environment_id` (String) The identifier of the Amazon DataZone environment in which the action is taking place
- `environment_identifier` (String) The identifier of the Amazon DataZone environment in which the action is taking place
- `identifier` (String) The ID of the Amazon DataZone environment action.
- `name` (String) The name of the environment action.
- `parameters` (Attributes) The parameters of the environment action. (see [below for nested schema](#nestedatt--parameters))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Read-Only:

- `uri` (String) The URI of the console link specified as part of the environment action.
