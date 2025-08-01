---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigatewayv2_routing_rule Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Schema for AWS::ApiGatewayV2::RoutingRule
---

# awscc_apigatewayv2_routing_rule (Resource)

Schema for AWS::ApiGatewayV2::RoutingRule



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `actions` (Attributes List) (see [below for nested schema](#nestedatt--actions))
- `conditions` (Attributes List) (see [below for nested schema](#nestedatt--conditions))
- `domain_name_arn` (String) The amazon resource name (ARN) of the domain name resource.
- `priority` (Number)

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `routing_rule_arn` (String) Amazon Resource Name (ARN) of the resource.
- `routing_rule_id` (String) RoutingRule Id generated by service

<a id="nestedatt--actions"></a>
### Nested Schema for `actions`

Required:

- `invoke_api` (Attributes) (see [below for nested schema](#nestedatt--actions--invoke_api))

<a id="nestedatt--actions--invoke_api"></a>
### Nested Schema for `actions.invoke_api`

Required:

- `api_id` (String)
- `stage` (String)

Optional:

- `strip_base_path` (Boolean)



<a id="nestedatt--conditions"></a>
### Nested Schema for `conditions`

Optional:

- `match_base_paths` (Attributes) (see [below for nested schema](#nestedatt--conditions--match_base_paths))
- `match_headers` (Attributes) (see [below for nested schema](#nestedatt--conditions--match_headers))

<a id="nestedatt--conditions--match_base_paths"></a>
### Nested Schema for `conditions.match_base_paths`

Optional:

- `any_of` (List of String)


<a id="nestedatt--conditions--match_headers"></a>
### Nested Schema for `conditions.match_headers`

Optional:

- `any_of` (Attributes List) (see [below for nested schema](#nestedatt--conditions--match_headers--any_of))

<a id="nestedatt--conditions--match_headers--any_of"></a>
### Nested Schema for `conditions.match_headers.any_of`

Optional:

- `header` (String)
- `value_glob` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_apigatewayv2_routing_rule.example
  id = "routing_rule_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_apigatewayv2_routing_rule.example "routing_rule_arn"
```
