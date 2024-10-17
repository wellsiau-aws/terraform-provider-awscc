---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_quicksight_folder Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::QuickSight::Folder
---

# awscc_quicksight_folder (Data Source)

Data Source schema for AWS::QuickSight::Folder



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) <p>The Amazon Resource Name (ARN) for the folder.</p>
- `aws_account_id` (String)
- `created_time` (String) <p>The time that the folder was created.</p>
- `folder_id` (String)
- `folder_type` (String)
- `last_updated_time` (String) <p>The time that the folder was last updated.</p>
- `name` (String)
- `parent_folder_arn` (String)
- `permissions` (Attributes List) (see [below for nested schema](#nestedatt--permissions))
- `sharing_model` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Read-Only:

- `actions` (List of String) <p>The IAM action to grant or revoke permissions on.</p>
- `principal` (String) <p>The Amazon Resource Name (ARN) of the principal. This can be one of the
            following:</p>
         <ul>
            <li>
               <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>
            </li>
            <li>
               <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>
            </li>
            <li>
               <p>The ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight
                    ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.
                    (This is less common.) </p>
            </li>
         </ul>


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) <p>Tag key.</p>
- `value` (String) <p>Tag value.</p>