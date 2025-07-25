
---
page_title: "awscc_appsync_resolver Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::AppSync::Resolver resource defines the logical GraphQL resolver that you attach to fields in a schema. Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see Resolver Mapping Template Reference https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html.
  When you submit an update, CFNLong updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CFNshort template. Changing the S3 file content without changing a property value will not result in an update operation.
  See Update Behaviors of Stack Resources https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html in the User Guide.
---

# awscc_appsync_resolver (Resource)

The ``AWS::AppSync::Resolver`` resource defines the logical GraphQL resolver that you attach to fields in a schema. Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see [Resolver Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html).
  When you submit an update, CFNLong updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CFNshort template. Changing the S3 file content without changing a property value will not result in an update operation.
 See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *User Guide*.

## Example Usage

### AppSync Query Resolver with Caching

Creates an AppSync resolver for a Query type field with request/response mapping templates and caching configuration, integrated with a NONE type data source.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create the AppSync API
resource "awscc_appsync_graph_ql_api" "example" {
  name                = "example-api"
  authentication_type = "API_KEY"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create AppSync Datasource
resource "awscc_appsync_data_source" "example" {
  api_id = awscc_appsync_graph_ql_api.example.id
  name   = "example_datasource"
  type   = "NONE"
}

# Create an AppSync Resolver
resource "awscc_appsync_resolver" "example" {
  api_id           = awscc_appsync_graph_ql_api.example.id
  type_name        = "Query"
  field_name       = "hello"
  data_source_name = awscc_appsync_data_source.example.name
  kind             = "UNIT"

  request_mapping_template = <<EOF
{
    "version": "2018-05-29",
    "payload": "Hello from AppSync!"
}
EOF

  response_mapping_template = "$util.toJson($context.result)"

  caching_config = {
    ttl          = 60
    caching_keys = ["$context.identity.sub"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_id` (String) The APSYlong GraphQL API to which you want to attach this resolver.
- `field_name` (String) The GraphQL field on a type that invokes the resolver.
- `type_name` (String) The GraphQL type that invokes this resolver.

### Optional

- `caching_config` (Attributes) The caching configuration for the resolver. (see [below for nested schema](#nestedatt--caching_config))
- `code` (String) The ``resolver`` code that contains the request and response functions. When code is used, the ``runtime`` is required. The runtime value must be ``APPSYNC_JS``.
- `code_s3_location` (String) The Amazon S3 endpoint.
- `data_source_name` (String) The resolver data source name.
- `kind` (String) The resolver type.
  +  *UNIT*: A UNIT resolver type. A UNIT resolver is the default resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.
  +  *PIPELINE*: A PIPELINE resolver type. You can use a PIPELINE resolver to invoke a series of ``Function`` objects in a serial manner. You can use a pipeline resolver to run a GraphQL query against multiple data sources.
- `max_batch_size` (Number) The maximum number of resolver request inputs that will be sent to a single LAMlong function in a ``BatchInvoke`` operation.
- `metrics_config` (String) Enables or disables enhanced resolver metrics for specified resolvers. Note that ``MetricsConfig`` won't be used unless the ``resolverLevelMetricsBehavior`` value is set to ``PER_RESOLVER_METRICS``. If the ``resolverLevelMetricsBehavior`` is set to ``FULL_REQUEST_RESOLVER_METRICS`` instead, ``MetricsConfig`` will be ignored. However, you can still set its value.
- `pipeline_config` (Attributes) Functions linked with the pipeline resolver. (see [below for nested schema](#nestedatt--pipeline_config))
- `request_mapping_template` (String) The request mapping template.
 Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.
- `request_mapping_template_s3_location` (String) The location of a request mapping template in an S3 bucket. Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template.
- `response_mapping_template` (String) The response mapping template.
- `response_mapping_template_s3_location` (String) The location of a response mapping template in an S3 bucket. Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template.
- `runtime` (Attributes) Describes a runtime used by an APSYlong resolver or APSYlong function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. (see [below for nested schema](#nestedatt--runtime))
- `sync_config` (Attributes) The ``SyncConfig`` for a resolver attached to a versioned data source. (see [below for nested schema](#nestedatt--sync_config))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `resolver_arn` (String)

<a id="nestedatt--caching_config"></a>
### Nested Schema for `caching_config`

Optional:

- `caching_keys` (List of String) The caching keys for a resolver that has caching activated.
 Valid values are entries from the ``$context.arguments``, ``$context.source``, and ``$context.identity`` maps.
- `ttl` (Number) The TTL in seconds for a resolver that has caching activated.
 Valid values are 1?3,600 seconds.


<a id="nestedatt--pipeline_config"></a>
### Nested Schema for `pipeline_config`

Optional:

- `functions` (List of String) A list of ``Function`` objects.


<a id="nestedatt--runtime"></a>
### Nested Schema for `runtime`

Optional:

- `name` (String) The ``name`` of the runtime to use. Currently, the only allowed value is ``APPSYNC_JS``.
- `runtime_version` (String) The ``version`` of the runtime to use. Currently, the only allowed version is ``1.0.0``.


<a id="nestedatt--sync_config"></a>
### Nested Schema for `sync_config`

Optional:

- `conflict_detection` (String) The Conflict Detection strategy to use.
  +  *VERSION*: Detect conflicts based on object versions for this resolver.
  +  *NONE*: Do not detect conflicts when invoking this resolver.
- `conflict_handler` (String) The Conflict Resolution strategy to perform in the event of a conflict.
  +  *OPTIMISTIC_CONCURRENCY*: Resolve conflicts by rejecting mutations when versions don't match the latest version at the server.
  +  *AUTOMERGE*: Resolve conflicts with the Automerge conflict resolution strategy.
  +  *LAMBDA*: Resolve conflicts with an LAMlong function supplied in the ``LambdaConflictHandlerConfig``.
- `lambda_conflict_handler_config` (Attributes) The ``LambdaConflictHandlerConfig`` when configuring ``LAMBDA`` as the Conflict Handler. (see [below for nested schema](#nestedatt--sync_config--lambda_conflict_handler_config))

<a id="nestedatt--sync_config--lambda_conflict_handler_config"></a>
### Nested Schema for `sync_config.lambda_conflict_handler_config`

Optional:

- `lambda_conflict_handler_arn` (String) The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_appsync_resolver.example
  id = "resolver_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_appsync_resolver.example "resolver_arn"
```
