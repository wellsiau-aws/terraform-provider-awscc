---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_nimblestudio_studio Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::NimbleStudio::Studio
---

# awscc_nimblestudio_studio (Resource)

Resource Type definition for AWS::NimbleStudio::Studio



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_role_arn` (String)
- `display_name` (String)
- `studio_name` (String)
- `user_role_arn` (String)

### Optional

- `studio_encryption_configuration` (Attributes) (see [below for nested schema](#nestedatt--studio_encryption_configuration))
- `tags` (Map of String)

### Read-Only

- `home_region` (String)
- `id` (String) Uniquely identifies the resource.
- `sso_client_id` (String)
- `studio_id` (String)
- `studio_url` (String)

<a id="nestedatt--studio_encryption_configuration"></a>
### Nested Schema for `studio_encryption_configuration`

Optional:

- `key_arn` (String)
- `key_type` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_nimblestudio_studio.example
  id = "studio_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_nimblestudio_studio.example "studio_id"
```
