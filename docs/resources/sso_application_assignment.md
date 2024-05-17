---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_sso_application_assignment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for SSO application access grant to a user or group.
---

# awscc_sso_application_assignment (Resource)

Resource Type definition for SSO application access grant to a user or group.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_arn` (String) The ARN of the application.
- `principal_id` (String) An identifier for an object in IAM Identity Center, such as a user or group
- `principal_type` (String) The entity type for which the assignment will be created.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_sso_application_assignment.example <resource ID>
```