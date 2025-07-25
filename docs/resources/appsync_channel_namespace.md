
---
page_title: "awscc_appsync_channel_namespace Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AppSync ChannelNamespace
---

# awscc_appsync_channel_namespace (Resource)

Resource schema for AppSync ChannelNamespace

## Example Usage

### AppSync Channel Namespace with API Key Authentication

Creates an AppSync Channel Namespace with API key authentication for both publish and subscribe operations, integrated with an AppSync GraphQL API.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create the AppSync API first
resource "aws_appsync_graphql_api" "example" {
  name                = "example-channel-namespace-api"
  authentication_type = "API_KEY"

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Add a delay to ensure the API is ready
resource "time_sleep" "wait_30_seconds" {
  depends_on      = [aws_appsync_graphql_api.example]
  create_duration = "30s"
}

# Create the channel namespace
resource "awscc_appsync_channel_namespace" "example" {
  depends_on = [time_sleep.wait_30_seconds]
  name       = "example-namespace"
  api_id     = aws_appsync_graphql_api.example.id

  # Example of publish and subscribe auth modes
  publish_auth_modes = [{
    auth_type = "API_KEY"
  }]

  subscribe_auth_modes = [{
    auth_type = "API_KEY"
  }]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_id` (String) AppSync Api Id that this Channel Namespace belongs to.
- `name` (String) Namespace indentifier.

### Optional

- `code_handlers` (String) String of APPSYNC_JS code to be used by the handlers.
- `code_s3_location` (String) The Amazon S3 endpoint where the code is located.
- `handler_configs` (Attributes) (see [below for nested schema](#nestedatt--handler_configs))
- `publish_auth_modes` (Attributes List) List of AuthModes supported for Publish operations. (see [below for nested schema](#nestedatt--publish_auth_modes))
- `subscribe_auth_modes` (Attributes List) List of AuthModes supported for Subscribe operations. (see [below for nested schema](#nestedatt--subscribe_auth_modes))
- `tags` (Attributes Set) An arbitrary set of tags (key-value pairs) for this AppSync API. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `channel_namespace_arn` (String) The Amazon Resource Name (ARN) for the Channel Namespace.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--handler_configs"></a>
### Nested Schema for `handler_configs`

Optional:

- `on_publish` (Attributes) (see [below for nested schema](#nestedatt--handler_configs--on_publish))
- `on_subscribe` (Attributes) (see [below for nested schema](#nestedatt--handler_configs--on_subscribe))

<a id="nestedatt--handler_configs--on_publish"></a>
### Nested Schema for `handler_configs.on_publish`

Optional:

- `behavior` (String) Integration behavior for a handler configuration.
- `integration` (Attributes) (see [below for nested schema](#nestedatt--handler_configs--on_publish--integration))

<a id="nestedatt--handler_configs--on_publish--integration"></a>
### Nested Schema for `handler_configs.on_publish.integration`

Optional:

- `data_source_name` (String) Data source to invoke for this integration.
- `lambda_config` (Attributes) (see [below for nested schema](#nestedatt--handler_configs--on_publish--integration--lambda_config))

<a id="nestedatt--handler_configs--on_publish--integration--lambda_config"></a>
### Nested Schema for `handler_configs.on_publish.integration.lambda_config`

Optional:

- `invoke_type` (String) Invocation type for direct lambda integrations.




<a id="nestedatt--handler_configs--on_subscribe"></a>
### Nested Schema for `handler_configs.on_subscribe`

Optional:

- `behavior` (String) Integration behavior for a handler configuration.
- `integration` (Attributes) (see [below for nested schema](#nestedatt--handler_configs--on_subscribe--integration))

<a id="nestedatt--handler_configs--on_subscribe--integration"></a>
### Nested Schema for `handler_configs.on_subscribe.integration`

Optional:

- `data_source_name` (String) Data source to invoke for this integration.
- `lambda_config` (Attributes) (see [below for nested schema](#nestedatt--handler_configs--on_subscribe--integration--lambda_config))

<a id="nestedatt--handler_configs--on_subscribe--integration--lambda_config"></a>
### Nested Schema for `handler_configs.on_subscribe.integration.lambda_config`

Optional:

- `invoke_type` (String) Invocation type for direct lambda integrations.





<a id="nestedatt--publish_auth_modes"></a>
### Nested Schema for `publish_auth_modes`

Optional:

- `auth_type` (String) Security configuration for your AppSync API.


<a id="nestedatt--subscribe_auth_modes"></a>
### Nested Schema for `subscribe_auth_modes`

Optional:

- `auth_type` (String) Security configuration for your AppSync API.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) A string used to identify this tag. You can specify a maximum of 128 characters for a tag key.
- `value` (String) A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_appsync_channel_namespace.example
  id = "channel_namespace_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_appsync_channel_namespace.example "channel_namespace_arn"
```
