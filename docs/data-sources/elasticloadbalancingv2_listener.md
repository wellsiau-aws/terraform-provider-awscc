---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_elasticloadbalancingv2_listener Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ElasticLoadBalancingV2::Listener
---

# awscc_elasticloadbalancingv2_listener (Data Source)

Data Source schema for AWS::ElasticLoadBalancingV2::Listener



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `alpn_policy` (List of String) [TLS listener] The name of the Application-Layer Protocol Negotiation (ALPN) policy.
- `certificates` (Attributes List) The default SSL server certificate for a secure listener. You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
 To create a certificate list for a secure listener, use [AWS::ElasticLoadBalancingV2::ListenerCertificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenercertificate.html). (see [below for nested schema](#nestedatt--certificates))
- `default_actions` (Attributes List) The actions for the default rule. You cannot define a condition for a default rule.
 To create additional rules for an Application Load Balancer, use [AWS::ElasticLoadBalancingV2::ListenerRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html). (see [below for nested schema](#nestedatt--default_actions))
- `listener_arn` (String)
- `listener_attributes` (Attributes Set) The listener attributes. (see [below for nested schema](#nestedatt--listener_attributes))
- `load_balancer_arn` (String) The Amazon Resource Name (ARN) of the load balancer.
- `mutual_authentication` (Attributes) The mutual authentication configuration information. (see [below for nested schema](#nestedatt--mutual_authentication))
- `port` (Number) The port on which the load balancer is listening. You can't specify a port for a Gateway Load Balancer.
- `protocol` (String) The protocol for connections from clients to the load balancer. For Application Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You can't specify a protocol for a Gateway Load Balancer.
- `ssl_policy` (String) [HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported. For more information, see [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/describe-ssl-policies.html) in the *Application Load Balancers Guide* and [Security policies](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/describe-ssl-policies.html) in the *Network Load Balancers Guide*.
 Updating the security policy can result in interruptions if the load balancer is handling a high volume of traffic. To decrease the possibility of an interruption if your load balancer is handling a high volume of traffic, create an additional load balancer or request an LCU reservation.

<a id="nestedatt--certificates"></a>
### Nested Schema for `certificates`

Read-Only:

- `certificate_arn` (String) The Amazon Resource Name (ARN) of the certificate.


<a id="nestedatt--default_actions"></a>
### Nested Schema for `default_actions`

Read-Only:

- `authenticate_cognito_config` (Attributes) [HTTPS listeners] Information for using Amazon Cognito to authenticate users. Specify only when ``Type`` is ``authenticate-cognito``. (see [below for nested schema](#nestedatt--default_actions--authenticate_cognito_config))
- `authenticate_oidc_config` (Attributes) [HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC). Specify only when ``Type`` is ``authenticate-oidc``. (see [below for nested schema](#nestedatt--default_actions--authenticate_oidc_config))
- `fixed_response_config` (Attributes) [Application Load Balancer] Information for creating an action that returns a custom HTTP response. Specify only when ``Type`` is ``fixed-response``. (see [below for nested schema](#nestedatt--default_actions--fixed_response_config))
- `forward_config` (Attributes) Information for creating an action that distributes requests among one or more target groups. For Network Load Balancers, you can specify a single target group. Specify only when ``Type`` is ``forward``. If you specify both ``ForwardConfig`` and ``TargetGroupArn``, you can specify only one target group using ``ForwardConfig`` and it must be the same target group specified in ``TargetGroupArn``. (see [below for nested schema](#nestedatt--default_actions--forward_config))
- `order` (Number) The order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first.
- `redirect_config` (Attributes) [Application Load Balancer] Information for creating a redirect action. Specify only when ``Type`` is ``redirect``. (see [below for nested schema](#nestedatt--default_actions--redirect_config))
- `target_group_arn` (String) The Amazon Resource Name (ARN) of the target group. Specify only when ``Type`` is ``forward`` and you want to route to a single target group. To route to one or more target groups, use ``ForwardConfig`` instead.
- `type` (String) The type of action.

<a id="nestedatt--default_actions--authenticate_cognito_config"></a>
### Nested Schema for `default_actions.authenticate_cognito_config`

Read-Only:

- `authentication_request_extra_params` (Map of String) The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
- `on_unauthenticated_request` (String) The behavior if the user is not authenticated. The following are possible values:
  +  deny```` - Return an HTTP 401 Unauthorized error.
  +  allow```` - Allow the request to be forwarded to the target.
  +  authenticate```` - Redirect the request to the IdP authorization endpoint. This is the default value.
- `scope` (String) The set of user claims to be requested from the IdP. The default is ``openid``.
 To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
- `session_cookie_name` (String) The name of the cookie used to maintain session information. The default is AWSELBAuthSessionCookie.
- `session_timeout` (String) The maximum duration of the authentication session, in seconds. The default is 604800 seconds (7 days).
- `user_pool_arn` (String) The Amazon Resource Name (ARN) of the Amazon Cognito user pool.
- `user_pool_client_id` (String) The ID of the Amazon Cognito user pool client.
- `user_pool_domain` (String) The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.


<a id="nestedatt--default_actions--authenticate_oidc_config"></a>
### Nested Schema for `default_actions.authenticate_oidc_config`

Read-Only:

- `authentication_request_extra_params` (Map of String) The query parameters (up to 10) to include in the redirect request to the authorization endpoint.
- `authorization_endpoint` (String) The authorization endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.
- `client_id` (String) The OAuth 2.0 client identifier.
- `client_secret` (String) The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set ``UseExistingClientSecret`` to true.
- `issuer` (String) The OIDC issuer identifier of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.
- `on_unauthenticated_request` (String) The behavior if the user is not authenticated. The following are possible values:
  +  deny```` - Return an HTTP 401 Unauthorized error.
  +  allow```` - Allow the request to be forwarded to the target.
  +  authenticate```` - Redirect the request to the IdP authorization endpoint. This is the default value.
- `scope` (String) The set of user claims to be requested from the IdP. The default is ``openid``.
 To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.
- `session_cookie_name` (String) The name of the cookie used to maintain session information. The default is AWSELBAuthSessionCookie.
- `session_timeout` (String) The maximum duration of the authentication session, in seconds. The default is 604800 seconds (7 days).
- `token_endpoint` (String) The token endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.
- `use_existing_client_secret` (Boolean) Indicates whether to use the existing client secret when modifying a rule. If you are creating a rule, you can omit this parameter or set it to false.
- `user_info_endpoint` (String) The user info endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.


<a id="nestedatt--default_actions--fixed_response_config"></a>
### Nested Schema for `default_actions.fixed_response_config`

Read-Only:

- `content_type` (String) The content type.
 Valid Values: text/plain | text/css | text/html | application/javascript | application/json
- `message_body` (String) The message.
- `status_code` (String) The HTTP response code (2XX, 4XX, or 5XX).


<a id="nestedatt--default_actions--forward_config"></a>
### Nested Schema for `default_actions.forward_config`

Read-Only:

- `target_group_stickiness_config` (Attributes) Information about the target group stickiness for a rule. (see [below for nested schema](#nestedatt--default_actions--forward_config--target_group_stickiness_config))
- `target_groups` (Attributes List) Information about how traffic will be distributed between multiple target groups in a forward rule. (see [below for nested schema](#nestedatt--default_actions--forward_config--target_groups))

<a id="nestedatt--default_actions--forward_config--target_group_stickiness_config"></a>
### Nested Schema for `default_actions.forward_config.target_group_stickiness_config`

Read-Only:

- `duration_seconds` (Number) The time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
- `enabled` (Boolean) Indicates whether target group stickiness is enabled.


<a id="nestedatt--default_actions--forward_config--target_groups"></a>
### Nested Schema for `default_actions.forward_config.target_groups`

Read-Only:

- `target_group_arn` (String) The Amazon Resource Name (ARN) of the target group.
- `weight` (Number) The weight. The range is 0 to 999.



<a id="nestedatt--default_actions--redirect_config"></a>
### Nested Schema for `default_actions.redirect_config`

Read-Only:

- `host` (String) The hostname. This component is not percent-encoded. The hostname can contain #{host}.
- `path` (String) The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
- `port` (String) The port. You can specify a value from 1 to 65535 or #{port}.
- `protocol` (String) The protocol. You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You can't redirect HTTPS to HTTP.
- `query` (String) The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
- `status_code` (String) The HTTP redirect code. The redirect is either permanent (HTTP 301) or temporary (HTTP 302).



<a id="nestedatt--listener_attributes"></a>
### Nested Schema for `listener_attributes`

Read-Only:

- `key` (String) The name of the attribute.
 The following attribute is supported by Network Load Balancers, and Gateway Load Balancers.
  +  ``tcp.idle_timeout.seconds`` - The tcp idle timeout value, in seconds. The valid range is 60-6000 seconds. The default is 350 seconds.
  
 The following attributes are only supported by Application Load Balancers.
  +  ``routing.http.request.x_amzn_mtls_clientcert_serial_number.header_name`` - Enables you to modify the header name of the *X-Amzn-Mtls-Clientcert-Serial-Number* HTTP request header.
  +  ``routing.http.request.x_amzn_mtls_clientcert_issuer.header_name`` - Enables you to modify the header name of the *X-Amzn-Mtls-Clientcert-Issuer* HTTP request header.
  +  ``routing.http.request.x_amzn_mtls_clientcert_subject.header_name`` - Enables you to modify the header name of the *X-Amzn-Mtls-Clientcert-Subject* HTTP request header.
  +  ``routing.http.request.x_amzn_mtls_clientcert_validity.header_name`` - Enables you to modify the header name of the *X-Amzn-Mtls-Clientcert-Validity* HTTP request header.
  +  ``routing.http.request.x_amzn_mtls_clientcert_leaf.header_name`` - Enables you to modify the header name of the *X-Amzn-Mtls-Clientcert-Leaf* HTTP request header.
  +  ``routing.http.request.x_amzn_mtls_clientcert.header_name`` - Enables you to modify the header name of the *X-Amzn-Mtls-Clientcert* HTTP request header.
  +  ``routing.http.request.x_amzn_tls_version.header_name`` - Enables you to modify the header name of the *X-Amzn-Tls-Version* HTTP request header.
  +  ``routing.http.request.x_amzn_tls_cipher_suite.header_name`` - Enables you to modify the header name of the *X-Amzn-Tls-Cipher-Suite* HTTP request header.
  +  ``routing.http.response.server.enabled`` - Enables you to allow or remove the HTTP response server header.
  +  ``routing.http.response.strict_transport_security.header_value`` - Informs browsers that the site should only be accessed using HTTPS, and that any future attempts to access it using HTTP should automatically be converted to HTTPS.
  +  ``routing.http.response.access_control_allow_origin.header_value`` - Specifies which origins are allowed to access the server.
  +  ``routing.http.response.access_control_allow_methods.header_value`` - Returns which HTTP methods are allowed when accessing the server from a different origin.
  +  ``routing.http.response.access_control_allow_headers.header_value`` - Specifies which headers can be used during the request.
  +  ``routing.http.response.access_control_allow_credentials.header_value`` - Indicates whether the browser should include credentials such as cookies or authentication when making requests.
  +  ``routing.http.response.access_control_expose_headers.header_value`` - Returns which headers the browser can expose to the requesting client.
  +  ``routing.http.response.access_control_max_age.header_value`` - Specifies how long the results of a preflight request can be cached, in seconds.
  +  ``routing.http.response.content_security_policy.header_value`` - Specifies restrictions enforced by the browser to help minimize the risk of certain types of security threats.
  +  ``routing.http.response.x_content_type_options.header_value`` - Indicates whether the MIME types advertised in the *Content-Type* headers should be followed and not be changed.
  +  ``routing.http.response.x_frame_options.header_value`` - Indicates whether the browser is allowed to render a page in a *frame*, *iframe*, *embed* or *object*.
- `value` (String) The value of the attribute.


<a id="nestedatt--mutual_authentication"></a>
### Nested Schema for `mutual_authentication`

Read-Only:

- `advertise_trust_store_ca_names` (String) Indicates whether trust store CA certificate names are advertised.
- `ignore_client_certificate_expiry` (Boolean) Indicates whether expired client certificates are ignored.
- `mode` (String) The client certificate handling method. Options are ``off``, ``passthrough`` or ``verify``. The default value is ``off``.
- `trust_store_arn` (String) The Amazon Resource Name (ARN) of the trust store.
