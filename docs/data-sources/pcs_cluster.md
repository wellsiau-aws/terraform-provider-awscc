---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_pcs_cluster Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::PCS::Cluster
---

# awscc_pcs_cluster (Data Source)

Data Source schema for AWS::PCS::Cluster



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) The unique Amazon Resource Name (ARN) of the cluster.
- `cluster_id` (String) The generated unique ID of the cluster.
- `endpoints` (Attributes List) The list of endpoints available for interaction with the scheduler. (see [below for nested schema](#nestedatt--endpoints))
- `error_info` (Attributes List) The list of errors that occurred during cluster provisioning. (see [below for nested schema](#nestedatt--error_info))
- `name` (String) The name that identifies the cluster.
- `networking` (Attributes) The networking configuration for the cluster's control plane. (see [below for nested schema](#nestedatt--networking))
- `scheduler` (Attributes) The cluster management and job scheduling software associated with the cluster. (see [below for nested schema](#nestedatt--scheduler))
- `size` (String) The size of the cluster.
- `slurm_configuration` (Attributes) Additional options related to the Slurm scheduler. (see [below for nested schema](#nestedatt--slurm_configuration))
- `status` (String) The provisioning status of the cluster. The provisioning status doesn't indicate the overall health of the cluster.
- `tags` (Map of String) 1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.

<a id="nestedatt--endpoints"></a>
### Nested Schema for `endpoints`

Read-Only:

- `port` (String) The endpoint's connection port number.
- `private_ip_address` (String) The endpoint's private IP address.
- `public_ip_address` (String) The endpoint's public IP address.
- `type` (String) Indicates the type of endpoint running at the specific IP address.


<a id="nestedatt--error_info"></a>
### Nested Schema for `error_info`

Read-Only:

- `code` (String) The short-form error code.
- `message` (String) The detailed error information.


<a id="nestedatt--networking"></a>
### Nested Schema for `networking`

Read-Only:

- `security_group_ids` (List of String) The list of security group IDs associated with the Elastic Network Interface (ENI) created in subnets.
- `subnet_ids` (List of String) The list of subnet IDs where AWS PCS creates an Elastic Network Interface (ENI) to enable communication between managed controllers and AWS PCS resources. The subnet must have an available IP address, cannot reside in AWS Outposts, AWS Wavelength, or an AWS Local Zone. AWS PCS currently supports only 1 subnet in this list.


<a id="nestedatt--scheduler"></a>
### Nested Schema for `scheduler`

Read-Only:

- `type` (String) The software AWS PCS uses to manage cluster scaling and job scheduling.
- `version` (String) The version of the specified scheduling software that AWS PCS uses to manage cluster scaling and job scheduling.


<a id="nestedatt--slurm_configuration"></a>
### Nested Schema for `slurm_configuration`

Read-Only:

- `auth_key` (Attributes) The shared Slurm key for authentication, also known as the cluster secret. (see [below for nested schema](#nestedatt--slurm_configuration--auth_key))
- `scale_down_idle_time_in_seconds` (Number) The time before an idle node is scaled down.
- `slurm_custom_settings` (Attributes List) Additional Slurm-specific configuration that directly maps to Slurm settings. (see [below for nested schema](#nestedatt--slurm_configuration--slurm_custom_settings))

<a id="nestedatt--slurm_configuration--auth_key"></a>
### Nested Schema for `slurm_configuration.auth_key`

Read-Only:

- `secret_arn` (String) The Amazon Resource Name (ARN) of the the shared Slurm key.
- `secret_version` (String) The version of the shared Slurm key.


<a id="nestedatt--slurm_configuration--slurm_custom_settings"></a>
### Nested Schema for `slurm_configuration.slurm_custom_settings`

Read-Only:

- `parameter_name` (String) AWS PCS supports configuration of the following Slurm parameters for clusters: Prolog, Epilog, and SelectTypeParameters.
- `parameter_value` (String) The value for the configured Slurm setting.