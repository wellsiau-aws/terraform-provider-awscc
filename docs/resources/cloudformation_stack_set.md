---
page_title: "awscc_cloudformation_stack_set Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances
---

# awscc_cloudformation_stack_set (Resource)

StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances

## Example Usage

### Basic StackSet Usage
To Create a basic StackSet (Self-Managed Permissions)
```terraform
resource "awscc_cloudformation_stack_set" "stackset" {
  stack_set_name = "network-stackset"

  permission_model = "SELF_MANAGED"

  template_body = jsonencode({

    Resources = {
      myVpc = {
        Type = "AWS::EC2::VPC"
        Properties = {
          CidrBlock = "10.0.0.0/16"
          Tags = [
            {
              Key   = "Name"
              Value = "Primary_CF_VPC"
            }
          ]
        }
      }
    }
  })
}
```

### Advanced StackSet Usage

> **_NOTE:_** Please make sure you Enable all features in AWS Organizations and Activate trusted access with AWS Organizations before deploying a Stackset with service-managed permissions.
Please refer to the following documentation for more infomration: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-activate-trusted-access.html

To Create a StackSet with Service-Managed Permissions 
```terraform
resource "awscc_cloudformation_stack_set" "stackset" {
  stack_set_name   = "network-stackset"
  permission_model = "SERVICE_MANAGED"
  auto_deployment = {
    enabled                          = true
    retain_stacks_on_account_removal = false
  }

  template_body = jsonencode({
    Resources = {
      myVpc = {
        Type = "AWS::EC2::VPC"
        Properties = {
          CidrBlock = "10.0.0.0/16"
          Tags = [
            {
              Key   = "Name"
              Value = "Primary_CF_VPC"
            }
          ]
        }
      }
    }
  })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `permission_model` (String) Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.
- `stack_set_name` (String) The name to associate with the stack set. The name must be unique in the Region where you create your stack set.

### Optional

- `administration_role_arn` (String) The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.
- `auto_deployment` (Attributes) Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED. (see [below for nested schema](#nestedatt--auto_deployment))
- `call_as` (String) Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.
- `capabilities` (Set of String) In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.
- `description` (String) A description of the stack set. You can use the description to identify the stack set's purpose or other important information.
- `execution_role_name` (String) The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.
- `managed_execution` (Attributes) Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations. (see [below for nested schema](#nestedatt--managed_execution))
- `operation_preferences` (Attributes) The user-specified preferences for how AWS CloudFormation performs a stack set operation. (see [below for nested schema](#nestedatt--operation_preferences))
- `parameters` (Attributes Set) The input parameters for the stack set template. (see [below for nested schema](#nestedatt--parameters))
- `stack_instances_group` (Attributes Set) A group of stack instances with parameters in some specific accounts and regions. (see [below for nested schema](#nestedatt--stack_instances_group))
- `tags` (Attributes Set) The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified. (see [below for nested schema](#nestedatt--tags))
- `template_body` (String) The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.
- `template_url` (String) Location of file containing the template body. The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket.

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `stack_set_id` (String) The ID of the stack set that you're creating.

<a id="nestedatt--auto_deployment"></a>
### Nested Schema for `auto_deployment`

Optional:

- `enabled` (Boolean) If set to true, StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions. If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.
- `retain_stacks_on_account_removal` (Boolean) If set to true, stack resources are retained when an account is removed from a target organization or OU. If set to false, stack resources are deleted. Specify only if Enabled is set to True.


<a id="nestedatt--managed_execution"></a>
### Nested Schema for `managed_execution`

Optional:

- `active` (Boolean) When true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.


<a id="nestedatt--operation_preferences"></a>
### Nested Schema for `operation_preferences`

Optional:

- `concurrency_mode` (String) Specifies how the concurrency level behaves during the operation execution.
- `failure_tolerance_count` (Number)
- `failure_tolerance_percentage` (Number)
- `max_concurrent_count` (Number)
- `max_concurrent_percentage` (Number)
- `region_concurrency_type` (String) The concurrency type of deploying StackSets operations in regions, could be in parallel or one region at a time
- `region_order` (List of String)


<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `parameter_key` (String) The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.
- `parameter_value` (String) The input value associated with the parameter.


<a id="nestedatt--stack_instances_group"></a>
### Nested Schema for `stack_instances_group`

Optional:

- `deployment_targets` (Attributes) The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions. (see [below for nested schema](#nestedatt--stack_instances_group--deployment_targets))
- `parameter_overrides` (Attributes Set) A list of stack set parameters whose values you want to override in the selected stack instances. (see [below for nested schema](#nestedatt--stack_instances_group--parameter_overrides))
- `regions` (Set of String) The names of one or more Regions where you want to create stack instances using the specified AWS account(s).

<a id="nestedatt--stack_instances_group--deployment_targets"></a>
### Nested Schema for `stack_instances_group.deployment_targets`

Optional:

- `account_filter_type` (String) The filter type you want to apply on organizational units and accounts.
- `accounts` (Set of String) AWS accounts that you want to create stack instances in the specified Region(s) for.
- `accounts_url` (String) Returns the value of the AccountsUrl property.
- `organizational_unit_ids` (Set of String) The organization root ID or organizational unit (OU) IDs to which StackSets deploys.


<a id="nestedatt--stack_instances_group--parameter_overrides"></a>
### Nested Schema for `stack_instances_group.parameter_overrides`

Optional:

- `parameter_key` (String) The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.
- `parameter_value` (String) The input value associated with the parameter.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.
- `value` (String) A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cloudformation_stack_set.example
  id = "stack_set_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cloudformation_stack_set.example "stack_set_id"
```