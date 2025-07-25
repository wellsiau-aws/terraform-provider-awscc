
---
page_title: "awscc_cognito_user_pool_risk_configuration_attachment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Cognito::UserPoolRiskConfigurationAttachment
---

# awscc_cognito_user_pool_risk_configuration_attachment (Resource)

Resource Type definition for AWS::Cognito::UserPoolRiskConfigurationAttachment

## Example Usage

### Cognito User Pool Risk Configuration

Configures advanced security features for a Cognito User Pool client including account takeover protection, compromised credentials detection, and IP-based risk exception rules.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create a Cognito User Pool
resource "aws_cognito_user_pool" "example" {
  name = "example-user-pool"

  admin_create_user_config {
    allow_admin_create_user_only = true
  }

  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create a Cognito User Pool Client using AWSCC
resource "awscc_cognito_user_pool_client" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  client_name  = "example-client"
  explicit_auth_flows = [
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_SRP_AUTH"
  ]

  prevent_user_existence_errors = "ENABLED"
  read_attributes               = ["email"]
  write_attributes              = ["email"]
}

# Configure the risk configuration attachment
resource "awscc_cognito_user_pool_risk_configuration_attachment" "example" {
  user_pool_id = aws_cognito_user_pool.example.id
  client_id    = awscc_cognito_user_pool_client.example.id

  account_takeover_risk_configuration = {
    actions = {
      high_action = {
        event_action = "BLOCK"
        notify       = false
      }
      medium_action = {
        event_action = "MFA_IF_CONFIGURED"
        notify       = false
      }
      low_action = {
        event_action = "NO_ACTION"
        notify       = false
      }
    }
  }

  compromised_credentials_risk_configuration = {
    actions = {
      event_action = "BLOCK"
    }
    event_filter = ["SIGN_IN", "PASSWORD_CHANGE"]
  }

  risk_exception_configuration = {
    blocked_ip_range_list = ["10.0.0.0/24"]
    skipped_ip_range_list = ["192.168.0.0/24"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `client_id` (String)
- `user_pool_id` (String)

### Optional

- `account_takeover_risk_configuration` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration))
- `compromised_credentials_risk_configuration` (Attributes) (see [below for nested schema](#nestedatt--compromised_credentials_risk_configuration))
- `risk_exception_configuration` (Attributes) (see [below for nested schema](#nestedatt--risk_exception_configuration))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--account_takeover_risk_configuration"></a>
### Nested Schema for `account_takeover_risk_configuration`

Optional:

- `actions` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--actions))
- `notify_configuration` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--notify_configuration))

<a id="nestedatt--account_takeover_risk_configuration--actions"></a>
### Nested Schema for `account_takeover_risk_configuration.actions`

Optional:

- `high_action` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--actions--high_action))
- `low_action` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--actions--low_action))
- `medium_action` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--actions--medium_action))

<a id="nestedatt--account_takeover_risk_configuration--actions--high_action"></a>
### Nested Schema for `account_takeover_risk_configuration.actions.high_action`

Optional:

- `event_action` (String)
- `notify` (Boolean)


<a id="nestedatt--account_takeover_risk_configuration--actions--low_action"></a>
### Nested Schema for `account_takeover_risk_configuration.actions.low_action`

Optional:

- `event_action` (String)
- `notify` (Boolean)


<a id="nestedatt--account_takeover_risk_configuration--actions--medium_action"></a>
### Nested Schema for `account_takeover_risk_configuration.actions.medium_action`

Optional:

- `event_action` (String)
- `notify` (Boolean)



<a id="nestedatt--account_takeover_risk_configuration--notify_configuration"></a>
### Nested Schema for `account_takeover_risk_configuration.notify_configuration`

Optional:

- `block_email` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--notify_configuration--block_email))
- `from` (String)
- `mfa_email` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--notify_configuration--mfa_email))
- `no_action_email` (Attributes) (see [below for nested schema](#nestedatt--account_takeover_risk_configuration--notify_configuration--no_action_email))
- `reply_to` (String)
- `source_arn` (String)

<a id="nestedatt--account_takeover_risk_configuration--notify_configuration--block_email"></a>
### Nested Schema for `account_takeover_risk_configuration.notify_configuration.block_email`

Optional:

- `html_body` (String)
- `subject` (String)
- `text_body` (String)


<a id="nestedatt--account_takeover_risk_configuration--notify_configuration--mfa_email"></a>
### Nested Schema for `account_takeover_risk_configuration.notify_configuration.mfa_email`

Optional:

- `html_body` (String)
- `subject` (String)
- `text_body` (String)


<a id="nestedatt--account_takeover_risk_configuration--notify_configuration--no_action_email"></a>
### Nested Schema for `account_takeover_risk_configuration.notify_configuration.no_action_email`

Optional:

- `html_body` (String)
- `subject` (String)
- `text_body` (String)




<a id="nestedatt--compromised_credentials_risk_configuration"></a>
### Nested Schema for `compromised_credentials_risk_configuration`

Optional:

- `actions` (Attributes) (see [below for nested schema](#nestedatt--compromised_credentials_risk_configuration--actions))
- `event_filter` (List of String)

<a id="nestedatt--compromised_credentials_risk_configuration--actions"></a>
### Nested Schema for `compromised_credentials_risk_configuration.actions`

Optional:

- `event_action` (String)



<a id="nestedatt--risk_exception_configuration"></a>
### Nested Schema for `risk_exception_configuration`

Optional:

- `blocked_ip_range_list` (List of String)
- `skipped_ip_range_list` (List of String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cognito_user_pool_risk_configuration_attachment.example
  id = "user_pool_id|client_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cognito_user_pool_risk_configuration_attachment.example "user_pool_id|client_id"
```
