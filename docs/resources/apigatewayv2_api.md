---
page_title: "awscc_apigatewayv2_api Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::ApiGatewayV2::Api resource creates an API. WebSocket APIs and HTTP APIs are supported. For more information about WebSocket APIs, see About WebSocket APIs in API Gateway https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html in the API Gateway Developer Guide. For more information about HTTP APIs, see HTTP APIs https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api.html in the API Gateway Developer Guide.
---

# awscc_apigatewayv2_api (Resource)

The ``AWS::ApiGatewayV2::Api`` resource creates an API. WebSocket APIs and HTTP APIs are supported. For more information about WebSocket APIs, see [About WebSocket APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html) in the *API Gateway Developer Guide*. For more information about HTTP APIs, see [HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api.html) in the *API Gateway Developer Guide.*

## Example Usage

### Basic Websockets API

```terraform
resource "awscc_apigatewayv2_api" "example_api" {
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
  tags = {
    key   = "Modified By"
    value = "AWSCC"
  }
}
```

### Basic HTTP API

```terraform
resource "awscc_apigatewayv2_api" "example_http_api" {
  name          = "example-http-api"
  protocol_type = "HTTP"
  tags = {
    key   = "Modified By"
    value = "AWSCC"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key_selection_expression` (String) An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
- `base_path` (String) Specifies how to interpret the base path of the API during import. Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.
- `body` (String) The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
- `body_s3_location` (Attributes) The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources. (see [below for nested schema](#nestedatt--body_s3_location))
- `cors_configuration` (Attributes) A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information. (see [below for nested schema](#nestedatt--cors_configuration))
- `credentials_arn` (String) This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.
- `description` (String) The description of the API.
- `disable_execute_api_endpoint` (Boolean) Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
- `disable_schema_validation` (Boolean) Avoid validating models when creating a deployment. Supported only for WebSocket APIs.
- `fail_on_warnings` (Boolean) Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.
- `ip_address_type` (String)
- `name` (String) The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
- `protocol_type` (String) The API protocol. Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``.
- `route_key` (String) This property is part of quick create. If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.
- `route_selection_expression` (String) The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
- `tags` (Map of String) The collection of tags. Each tag element is associated with a given resource.
- `target` (String) This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
- `version` (String) A version identifier for the API.

### Read-Only

- `api_endpoint` (String)
- `api_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--body_s3_location"></a>
### Nested Schema for `body_s3_location`

Optional:

- `bucket` (String) The S3 bucket that contains the OpenAPI definition to import. Required if you specify a ``BodyS3Location`` for an API.
- `etag` (String) The Etag of the S3 object.
- `key` (String) The key of the S3 object. Required if you specify a ``BodyS3Location`` for an API.
- `version` (String) The version of the S3 object.


<a id="nestedatt--cors_configuration"></a>
### Nested Schema for `cors_configuration`

Optional:

- `allow_credentials` (Boolean) Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.
- `allow_headers` (List of String) Represents a collection of allowed headers. Supported only for HTTP APIs.
- `allow_methods` (List of String) Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.
- `allow_origins` (List of String) Represents a collection of allowed origins. Supported only for HTTP APIs.
- `expose_headers` (List of String) Represents a collection of exposed headers. Supported only for HTTP APIs.
- `max_age` (Number) The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_apigatewayv2_api.example
  id = "api_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_apigatewayv2_api.example "api_id"
```