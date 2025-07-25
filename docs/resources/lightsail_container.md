---
page_title: "awscc_lightsail_container Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Lightsail::Container
---

# awscc_lightsail_container (Resource)

Resource Type definition for AWS::Lightsail::Container

## Example Usage

```terraform
resource "awscc_lightsail_container" "example" {
  power        = "nano"
  scale        = "1"
  service_name = "example-service"
  container_service_deployment = {
    containers = [{
      container_name = "example-container"
      image          = "public.ecr.aws/nginx/nginx:latest"
      ports = [{
        port     = 80
        protocol = "HTTP"
      }]
    }]
    public_endpoint = {
      container_name = "example-container"
      container_port = 80
    }
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `power` (String) The power specification for the container service.
- `scale` (Number) The scale specification for the container service.
- `service_name` (String) The name for the container service.

### Optional

- `container_service_deployment` (Attributes) Describes a container deployment configuration of an Amazon Lightsail container service. (see [below for nested schema](#nestedatt--container_service_deployment))
- `is_disabled` (Boolean) A Boolean value to indicate whether the container service is disabled.
- `private_registry_access` (Attributes) A Boolean value to indicate whether the container service has access to private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. (see [below for nested schema](#nestedatt--private_registry_access))
- `public_domain_names` (Attributes Set) The public domain names to use with the container service, such as example.com and www.example.com. (see [below for nested schema](#nestedatt--public_domain_names))
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `container_arn` (String)
- `id` (String) Uniquely identifies the resource.
- `principal_arn` (String) The principal ARN of the container service.
- `url` (String) The publicly accessible URL of the container service.

<a id="nestedatt--container_service_deployment"></a>
### Nested Schema for `container_service_deployment`

Optional:

- `containers` (Attributes Set) An object that describes the configuration for the containers of the deployment. (see [below for nested schema](#nestedatt--container_service_deployment--containers))
- `public_endpoint` (Attributes) An object that describes the endpoint of the deployment. (see [below for nested schema](#nestedatt--container_service_deployment--public_endpoint))

<a id="nestedatt--container_service_deployment--containers"></a>
### Nested Schema for `container_service_deployment.containers`

Optional:

- `command` (Set of String) The launch command for the container.
- `container_name` (String) The name of the container.
- `environment` (Attributes Set) The environment variables of the container. (see [below for nested schema](#nestedatt--container_service_deployment--containers--environment))
- `image` (String) The name of the image used for the container.
- `ports` (Attributes Set) The open firewall ports of the container. (see [below for nested schema](#nestedatt--container_service_deployment--containers--ports))

<a id="nestedatt--container_service_deployment--containers--environment"></a>
### Nested Schema for `container_service_deployment.containers.environment`

Optional:

- `value` (String)
- `variable` (String)


<a id="nestedatt--container_service_deployment--containers--ports"></a>
### Nested Schema for `container_service_deployment.containers.ports`

Optional:

- `port` (String)
- `protocol` (String)



<a id="nestedatt--container_service_deployment--public_endpoint"></a>
### Nested Schema for `container_service_deployment.public_endpoint`

Optional:

- `container_name` (String) The name of the container for the endpoint.
- `container_port` (Number) The port of the container to which traffic is forwarded to.
- `health_check_config` (Attributes) An object that describes the health check configuration of the container. (see [below for nested schema](#nestedatt--container_service_deployment--public_endpoint--health_check_config))

<a id="nestedatt--container_service_deployment--public_endpoint--health_check_config"></a>
### Nested Schema for `container_service_deployment.public_endpoint.health_check_config`

Optional:

- `healthy_threshold` (Number) The number of consecutive health checks successes required before moving the container to the Healthy state. The default value is 2.
- `interval_seconds` (Number) The approximate interval, in seconds, between health checks of an individual container. You can specify between 5 and 300 seconds. The default value is 5.
- `path` (String) The path on the container on which to perform the health check. The default value is /.
- `success_codes` (String) The HTTP codes to use when checking for a successful response from a container. You can specify values between 200 and 499. You can specify multiple values (for example, 200,202) or a range of values (for example, 200-299).
- `timeout_seconds` (Number) The amount of time, in seconds, during which no response means a failed health check. You can specify between 2 and 60 seconds. The default value is 2.
- `unhealthy_threshold` (Number) The number of consecutive health check failures required before moving the container to the Unhealthy state. The default value is 2.




<a id="nestedatt--private_registry_access"></a>
### Nested Schema for `private_registry_access`

Optional:

- `ecr_image_puller_role` (Attributes) An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories. (see [below for nested schema](#nestedatt--private_registry_access--ecr_image_puller_role))

<a id="nestedatt--private_registry_access--ecr_image_puller_role"></a>
### Nested Schema for `private_registry_access.ecr_image_puller_role`

Optional:

- `is_active` (Boolean) A Boolean value that indicates whether to activate the role.

Read-Only:

- `principal_arn` (String) The Amazon Resource Name (ARN) of the role, if it is activated.



<a id="nestedatt--public_domain_names"></a>
### Nested Schema for `public_domain_names`

Optional:

- `certificate_name` (String)
- `domain_names` (Set of String) An object that describes the configuration for the containers of the deployment.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_lightsail_container.example
  id = "service_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_lightsail_container.example "service_name"
```