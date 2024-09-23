---
page_title: "awscc_lambda_alias Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Lambda::Alias
---

# awscc_lambda_alias (Resource)

Resource Type definition for AWS::Lambda::Alias

## Example Usage

### Basic usage with $LATEST version.

```terraform
resource "awscc_lambda_alias" "example" {
  function_name    = awscc_lambda_function.example.arn
  function_version = "$LATEST"
  name             = "example-alias"
  description      = "example alias"
}
```

### Alias with weighted config 

```terraform
resource "awscc_lambda_alias" "example" {
  function_name    = awscc_lambda_function.example.arn
  function_version = "v1"
  name             = "example-alias"
  description      = "example alias"
  routing_config = {
    additional_version_weights = [
      {
        function_version = "v2"
        function_weight  = 0.5
      }
    ]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `function_name` (String) The name of the Lambda function.
- `function_version` (String) The function version that the alias invokes.
- `name` (String) The name of the alias.

### Optional

- `description` (String) A description of the alias.
- `provisioned_concurrency_config` (Attributes) Specifies a provisioned concurrency configuration for a function's alias. (see [below for nested schema](#nestedatt--provisioned_concurrency_config))
- `routing_config` (Attributes) The routing configuration of the alias. (see [below for nested schema](#nestedatt--routing_config))

### Read-Only

- `alias_arn` (String) Lambda Alias ARN generated by the service.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--provisioned_concurrency_config"></a>
### Nested Schema for `provisioned_concurrency_config`

Optional:

- `provisioned_concurrent_executions` (Number) The amount of provisioned concurrency to allocate for the alias.


<a id="nestedatt--routing_config"></a>
### Nested Schema for `routing_config`

Optional:

- `additional_version_weights` (Attributes Set) The second version, and the percentage of traffic that's routed to it. (see [below for nested schema](#nestedatt--routing_config--additional_version_weights))

<a id="nestedatt--routing_config--additional_version_weights"></a>
### Nested Schema for `routing_config.additional_version_weights`

Optional:

- `function_version` (String) The qualifier of the second version.
- `function_weight` (Number) The percentage of traffic that the alias routes to the second version.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_lambda_alias.example "alias_arn"
```