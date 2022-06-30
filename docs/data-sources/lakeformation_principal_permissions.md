---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_lakeformation_principal_permissions Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::LakeFormation::PrincipalPermissions
---

# awscc_lakeformation_principal_permissions (Data Source)

Data Source schema for AWS::LakeFormation::PrincipalPermissions



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `catalog` (String)
- `permissions` (List of String)
- `permissions_with_grant_option` (List of String)
- `principal` (Attributes) (see [below for nested schema](#nestedatt--principal))
- `principal_identifier` (String)
- `resource` (Attributes) (see [below for nested schema](#nestedatt--resource))
- `resource_identifier` (String)

<a id="nestedatt--principal"></a>
### Nested Schema for `principal`

Read-Only:

- `data_lake_principal_identifier` (String)


<a id="nestedatt--resource"></a>
### Nested Schema for `resource`

Read-Only:

- `catalog` (Map of String)
- `data_cells_filter` (Attributes) (see [below for nested schema](#nestedatt--resource--data_cells_filter))
- `data_location` (Attributes) (see [below for nested schema](#nestedatt--resource--data_location))
- `database` (Attributes) (see [below for nested schema](#nestedatt--resource--database))
- `lf_tag` (Attributes) (see [below for nested schema](#nestedatt--resource--lf_tag))
- `lf_tag_policy` (Attributes) (see [below for nested schema](#nestedatt--resource--lf_tag_policy))
- `table` (Attributes) (see [below for nested schema](#nestedatt--resource--table))
- `table_with_columns` (Attributes) (see [below for nested schema](#nestedatt--resource--table_with_columns))

<a id="nestedatt--resource--data_cells_filter"></a>
### Nested Schema for `resource.data_cells_filter`

Read-Only:

- `database_name` (String)
- `name` (String)
- `table_catalog_id` (String)
- `table_name` (String)


<a id="nestedatt--resource--data_location"></a>
### Nested Schema for `resource.data_location`

Read-Only:

- `catalog_id` (String)
- `resource_arn` (String)


<a id="nestedatt--resource--database"></a>
### Nested Schema for `resource.database`

Read-Only:

- `catalog_id` (String)
- `name` (String)


<a id="nestedatt--resource--lf_tag"></a>
### Nested Schema for `resource.lf_tag`

Read-Only:

- `catalog_id` (String)
- `tag_key` (String)
- `tag_values` (List of String)


<a id="nestedatt--resource--lf_tag_policy"></a>
### Nested Schema for `resource.lf_tag_policy`

Read-Only:

- `catalog_id` (String)
- `expression` (Attributes List) (see [below for nested schema](#nestedatt--resource--lf_tag_policy--expression))
- `resource_type` (String)

<a id="nestedatt--resource--lf_tag_policy--expression"></a>
### Nested Schema for `resource.lf_tag_policy.expression`

Read-Only:

- `tag_key` (String)
- `tag_values` (List of String)



<a id="nestedatt--resource--table"></a>
### Nested Schema for `resource.table`

Read-Only:

- `catalog_id` (String)
- `database_name` (String)
- `name` (String)
- `table_wildcard` (Map of String)


<a id="nestedatt--resource--table_with_columns"></a>
### Nested Schema for `resource.table_with_columns`

Read-Only:

- `catalog_id` (String)
- `column_names` (List of String)
- `column_wildcard` (Attributes) (see [below for nested schema](#nestedatt--resource--table_with_columns--column_wildcard))
- `database_name` (String)
- `name` (String)

<a id="nestedatt--resource--table_with_columns--column_wildcard"></a>
### Nested Schema for `resource.table_with_columns.column_wildcard`

Read-Only:

- `excluded_column_names` (List of String)

