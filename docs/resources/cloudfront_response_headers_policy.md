---
page_title: "awscc_cloudfront_response_headers_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A response headers policy.
  A response headers policy contains information about a set of HTTP response headers.
  After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution. When it's attached to a cache behavior, the response headers policy affects the HTTP headers that CloudFront includes in HTTP responses to requests that match the cache behavior. CloudFront adds or removes response headers according to the configuration of the response headers policy.
  For more information, see Adding or removing HTTP headers in CloudFront responses https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html in the Amazon CloudFront Developer Guide.
---

# awscc_cloudfront_response_headers_policy (Resource)

A response headers policy.
 A response headers policy contains information about a set of HTTP response headers.
 After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution. When it's attached to a cache behavior, the response headers policy affects the HTTP headers that CloudFront includes in HTTP responses to requests that match the cache behavior. CloudFront adds or removes response headers according to the configuration of the response headers policy.
 For more information, see [Adding or removing HTTP headers in CloudFront responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html) in the *Amazon CloudFront Developer Guide*.

## Example Usage

The example below creates a CloudFront response headers policy.

```terraform
resource "awscc_cloudfront_response_headers_policy" "example" {
  response_headers_policy_config = {
    name    = "example-policy"
    comment = "test comment"
    cors_config = {
      access_control_allow_credentials = true

      access_control_allow_headers = {
        items = ["test"]
      }

      access_control_allow_methods = {
        items = ["GET"]
      }

      access_control_allow_origins = {
        items = ["test.example.comtest"]
      }

      origin_override = true
    }
  }
}
```

The example below creates a CloudFront response headers policy with a custom headers config.

```terraform
resource "awscc_cloudfront_response_headers_policy" "example" {
  response_headers_policy_config = {
    name = "example-policy"
    custom_headers_config = {
      items = [
        {
          header   = "X-Permitted-Cross-Domain-Policies"
          override = true
          value    = "none"
          }, {
          header   = "X-Test"
          override = true
          value    = "none"
        }
      ]
    }
  }
}
```

The example below creates a CloudFront response headers policy with a custom headers config and server timing headers config.

```terraform
resource "awscc_cloudfront_response_headers_policy" "example" {
  response_headers_policy_config = {
    name = "example-headers-policy"

    custom_headers_config = {
      items = [
        {
          header   = "X-Permitted-Cross-Domain-Policies"
          override = true
          value    = "none"
        }
      ]
    }
    server_timing_headers_config = {
      enabled       = true
      sampling_rate = 50
    }
  }
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `response_headers_policy_config` (Attributes) A response headers policy configuration. (see [below for nested schema](#nestedatt--response_headers_policy_config))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `last_modified_time` (String)
- `response_headers_policy_id` (String)

<a id="nestedatt--response_headers_policy_config"></a>
### Nested Schema for `response_headers_policy_config`

Required:

- `name` (String) A name to identify the response headers policy.
 The name must be unique for response headers policies in this AWS-account.

Optional:

- `comment` (String) A comment to describe the response headers policy.
 The comment cannot be longer than 128 characters.
- `cors_config` (Attributes) A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS). (see [below for nested schema](#nestedatt--response_headers_policy_config--cors_config))
- `custom_headers_config` (Attributes) A configuration for a set of custom HTTP response headers. (see [below for nested schema](#nestedatt--response_headers_policy_config--custom_headers_config))
- `remove_headers_config` (Attributes) A configuration for a set of HTTP headers to remove from the HTTP response. (see [below for nested schema](#nestedatt--response_headers_policy_config--remove_headers_config))
- `security_headers_config` (Attributes) A configuration for a set of security-related HTTP response headers. (see [below for nested schema](#nestedatt--response_headers_policy_config--security_headers_config))
- `server_timing_headers_config` (Attributes) A configuration for enabling the ``Server-Timing`` header in HTTP responses sent from CloudFront. (see [below for nested schema](#nestedatt--response_headers_policy_config--server_timing_headers_config))

<a id="nestedatt--response_headers_policy_config--cors_config"></a>
### Nested Schema for `response_headers_policy_config.cors_config`

Optional:

- `access_control_allow_credentials` (Boolean) A Boolean that CloudFront uses as the value for the ``Access-Control-Allow-Credentials`` HTTP response header.
 For more information about the ``Access-Control-Allow-Credentials`` HTTP response header, see [Access-Control-Allow-Credentials](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials) in the MDN Web Docs.
- `access_control_allow_headers` (Attributes) A list of HTTP header names that CloudFront includes as values for the ``Access-Control-Allow-Headers`` HTTP response header.
 For more information about the ``Access-Control-Allow-Headers`` HTTP response header, see [Access-Control-Allow-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--cors_config--access_control_allow_headers))
- `access_control_allow_methods` (Attributes) A list of HTTP methods that CloudFront includes as values for the ``Access-Control-Allow-Methods`` HTTP response header.
 For more information about the ``Access-Control-Allow-Methods`` HTTP response header, see [Access-Control-Allow-Methods](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--cors_config--access_control_allow_methods))
- `access_control_allow_origins` (Attributes) A list of origins (domain names) that CloudFront can use as the value for the ``Access-Control-Allow-Origin`` HTTP response header.
 For more information about the ``Access-Control-Allow-Origin`` HTTP response header, see [Access-Control-Allow-Origin](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--cors_config--access_control_allow_origins))
- `access_control_expose_headers` (Attributes) A list of HTTP headers that CloudFront includes as values for the ``Access-Control-Expose-Headers`` HTTP response header.
 For more information about the ``Access-Control-Expose-Headers`` HTTP response header, see [Access-Control-Expose-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--cors_config--access_control_expose_headers))
- `access_control_max_age_sec` (Number) A number that CloudFront uses as the value for the ``Access-Control-Max-Age`` HTTP response header.
 For more information about the ``Access-Control-Max-Age`` HTTP response header, see [Access-Control-Max-Age](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age) in the MDN Web Docs.
- `origin_override` (Boolean) A Boolean that determines whether CloudFront overrides HTTP response headers received from the origin with the ones specified in this response headers policy.

<a id="nestedatt--response_headers_policy_config--cors_config--access_control_allow_headers"></a>
### Nested Schema for `response_headers_policy_config.cors_config.access_control_allow_headers`

Optional:

- `items` (List of String) The list of HTTP header names. You can specify ``*`` to allow all headers.


<a id="nestedatt--response_headers_policy_config--cors_config--access_control_allow_methods"></a>
### Nested Schema for `response_headers_policy_config.cors_config.access_control_allow_methods`

Optional:

- `items` (List of String) The list of HTTP methods. Valid values are:
  +   ``GET`` 
  +   ``DELETE`` 
  +   ``HEAD`` 
  +   ``OPTIONS`` 
  +   ``PATCH`` 
  +   ``POST`` 
  +   ``PUT`` 
  +   ``ALL`` 
  
 ``ALL`` is a special value that includes all of the listed HTTP methods.


<a id="nestedatt--response_headers_policy_config--cors_config--access_control_allow_origins"></a>
### Nested Schema for `response_headers_policy_config.cors_config.access_control_allow_origins`

Optional:

- `items` (List of String) The list of origins (domain names). You can specify ``*`` to allow all origins.


<a id="nestedatt--response_headers_policy_config--cors_config--access_control_expose_headers"></a>
### Nested Schema for `response_headers_policy_config.cors_config.access_control_expose_headers`

Optional:

- `items` (List of String) The list of HTTP headers. You can specify ``*`` to expose all headers.



<a id="nestedatt--response_headers_policy_config--custom_headers_config"></a>
### Nested Schema for `response_headers_policy_config.custom_headers_config`

Optional:

- `items` (Attributes List) The list of HTTP response headers and their values. (see [below for nested schema](#nestedatt--response_headers_policy_config--custom_headers_config--items))

<a id="nestedatt--response_headers_policy_config--custom_headers_config--items"></a>
### Nested Schema for `response_headers_policy_config.custom_headers_config.items`

Optional:

- `header` (String) The HTTP response header name.
- `override` (Boolean) A Boolean that determines whether CloudFront overrides a response header with the same name received from the origin with the header specified here.
- `value` (String) The value for the HTTP response header.



<a id="nestedatt--response_headers_policy_config--remove_headers_config"></a>
### Nested Schema for `response_headers_policy_config.remove_headers_config`

Optional:

- `items` (Attributes Set) The list of HTTP header names. (see [below for nested schema](#nestedatt--response_headers_policy_config--remove_headers_config--items))

<a id="nestedatt--response_headers_policy_config--remove_headers_config--items"></a>
### Nested Schema for `response_headers_policy_config.remove_headers_config.items`

Optional:

- `header` (String) The HTTP header name.



<a id="nestedatt--response_headers_policy_config--security_headers_config"></a>
### Nested Schema for `response_headers_policy_config.security_headers_config`

Optional:

- `content_security_policy` (Attributes) The policy directives and their values that CloudFront includes as values for the ``Content-Security-Policy`` HTTP response header.
 For more information about the ``Content-Security-Policy`` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--security_headers_config--content_security_policy))
- `content_type_options` (Attributes) Determines whether CloudFront includes the ``X-Content-Type-Options`` HTTP response header with its value set to ``nosniff``.
 For more information about the ``X-Content-Type-Options`` HTTP response header, see [X-Content-Type-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--security_headers_config--content_type_options))
- `frame_options` (Attributes) Determines whether CloudFront includes the ``X-Frame-Options`` HTTP response header and the header's value.
 For more information about the ``X-Frame-Options`` HTTP response header, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--security_headers_config--frame_options))
- `referrer_policy` (Attributes) Determines whether CloudFront includes the ``Referrer-Policy`` HTTP response header and the header's value.
 For more information about the ``Referrer-Policy`` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--security_headers_config--referrer_policy))
- `strict_transport_security` (Attributes) Determines whether CloudFront includes the ``Strict-Transport-Security`` HTTP response header and the header's value.
 For more information about the ``Strict-Transport-Security`` HTTP response header, see [Security headers](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/understanding-response-headers-policies.html#understanding-response-headers-policies-security) in the *Amazon CloudFront Developer Guide* and [Strict-Transport-Security](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--security_headers_config--strict_transport_security))
- `xss_protection` (Attributes) Determines whether CloudFront includes the ``X-XSS-Protection`` HTTP response header and the header's value.
 For more information about the ``X-XSS-Protection`` HTTP response header, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs. (see [below for nested schema](#nestedatt--response_headers_policy_config--security_headers_config--xss_protection))

<a id="nestedatt--response_headers_policy_config--security_headers_config--content_security_policy"></a>
### Nested Schema for `response_headers_policy_config.security_headers_config.content_security_policy`

Optional:

- `content_security_policy` (String) The policy directives and their values that CloudFront includes as values for the ``Content-Security-Policy`` HTTP response header.
 For more information about the ``Content-Security-Policy`` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.
- `override` (Boolean) A Boolean that determines whether CloudFront overrides the ``Content-Security-Policy`` HTTP response header received from the origin with the one specified in this response headers policy.


<a id="nestedatt--response_headers_policy_config--security_headers_config--content_type_options"></a>
### Nested Schema for `response_headers_policy_config.security_headers_config.content_type_options`

Optional:

- `override` (Boolean) A Boolean that determines whether CloudFront overrides the ``X-Content-Type-Options`` HTTP response header received from the origin with the one specified in this response headers policy.


<a id="nestedatt--response_headers_policy_config--security_headers_config--frame_options"></a>
### Nested Schema for `response_headers_policy_config.security_headers_config.frame_options`

Optional:

- `frame_option` (String) The value of the ``X-Frame-Options`` HTTP response header. Valid values are ``DENY`` and ``SAMEORIGIN``.
 For more information about these values, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
- `override` (Boolean) A Boolean that determines whether CloudFront overrides the ``X-Frame-Options`` HTTP response header received from the origin with the one specified in this response headers policy.


<a id="nestedatt--response_headers_policy_config--security_headers_config--referrer_policy"></a>
### Nested Schema for `response_headers_policy_config.security_headers_config.referrer_policy`

Optional:

- `override` (Boolean) A Boolean that determines whether CloudFront overrides the ``Referrer-Policy`` HTTP response header received from the origin with the one specified in this response headers policy.
- `referrer_policy` (String) Determines whether CloudFront includes the ``Referrer-Policy`` HTTP response header and the header's value.
 For more information about the ``Referrer-Policy`` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.


<a id="nestedatt--response_headers_policy_config--security_headers_config--strict_transport_security"></a>
### Nested Schema for `response_headers_policy_config.security_headers_config.strict_transport_security`

Optional:

- `access_control_max_age_sec` (Number) A number that CloudFront uses as the value for the ``max-age`` directive in the ``Strict-Transport-Security`` HTTP response header.
- `include_subdomains` (Boolean) A Boolean that determines whether CloudFront includes the ``includeSubDomains`` directive in the ``Strict-Transport-Security`` HTTP response header.
- `override` (Boolean) A Boolean that determines whether CloudFront overrides the ``Strict-Transport-Security`` HTTP response header received from the origin with the one specified in this response headers policy.
- `preload` (Boolean) A Boolean that determines whether CloudFront includes the ``preload`` directive in the ``Strict-Transport-Security`` HTTP response header.


<a id="nestedatt--response_headers_policy_config--security_headers_config--xss_protection"></a>
### Nested Schema for `response_headers_policy_config.security_headers_config.xss_protection`

Optional:

- `mode_block` (Boolean) A Boolean that determines whether CloudFront includes the ``mode=block`` directive in the ``X-XSS-Protection`` header.
 For more information about this directive, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
- `override` (Boolean) A Boolean that determines whether CloudFront overrides the ``X-XSS-Protection`` HTTP response header received from the origin with the one specified in this response headers policy.
- `protection` (Boolean) A Boolean that determines the value of the ``X-XSS-Protection`` HTTP response header. When this setting is ``true``, the value of the ``X-XSS-Protection`` header is ``1``. When this setting is ``false``, the value of the ``X-XSS-Protection`` header is ``0``.
 For more information about these settings, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
- `report_uri` (String) A reporting URI, which CloudFront uses as the value of the ``report`` directive in the ``X-XSS-Protection`` header.
 You cannot specify a ``ReportUri`` when ``ModeBlock`` is ``true``.
 For more information about using a reporting URL, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.



<a id="nestedatt--response_headers_policy_config--server_timing_headers_config"></a>
### Nested Schema for `response_headers_policy_config.server_timing_headers_config`

Optional:

- `enabled` (Boolean) A Boolean that determines whether CloudFront adds the ``Server-Timing`` header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
- `sampling_rate` (Number) A number 0?100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the ``Server-Timing`` header to. When you set the sampling rate to 100, CloudFront adds the ``Server-Timing`` header to the HTTP response for every request that matches the cache behavior that this response headers policy is attached to. When you set it to 50, CloudFront adds the header to 50% of the responses for requests that match the cache behavior. You can set the sampling rate to any number 0?100 with up to four decimal places.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cloudfront_response_headers_policy.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cloudfront_response_headers_policy.example "id"
```