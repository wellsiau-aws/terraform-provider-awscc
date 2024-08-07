---
page_title: "awscc_rds_db_cluster_parameter_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::RDS::DBClusterParameterGroup resource creates a new Amazon RDS DB cluster parameter group.
  For information about configuring parameters for Amazon Aurora DB clusters, see Working with parameter groups https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html in the Amazon Aurora User Guide.
  If you apply a parameter group to a DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
  If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.
---

# awscc_rds_db_cluster_parameter_group (Resource)

The ``AWS::RDS::DBClusterParameterGroup`` resource creates a new Amazon RDS DB cluster parameter group.
 For information about configuring parameters for Amazon Aurora DB clusters, see [Working with parameter groups](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_WorkingWithParamGroups.html) in the *Amazon Aurora User Guide*.
  If you apply a parameter group to a DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
 If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.

## Example Usage

### Basic example
To create a RDS db cluster parameter group
```terraform
resource "awscc_rds_db_cluster_parameter_group" "this" {
  db_cluster_parameter_group_name = "rds-db-cluster-pg"
  description                     = "RDS DB cluster parameter group"
  family                          = "aurora5.6"

  parameters = {
    character_set_server = "utf8"
    character_set_client = "latin2"
    time_zone            = "US/Eastern"
  }

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) A friendly description for this DB cluster parameter group.
- `family` (String) The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
  The DB cluster parameter group family can't be changed when updating a DB cluster parameter group.
  To list all of the available parameter group families, use the following command:
  ``aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"`` 
 The output contains duplicates.
 For more information, see ``CreateDBClusterParameterGroup``.
- `parameters` (String) Provides a list of parameters for the DB cluster parameter group.

### Optional

- `db_cluster_parameter_group_name` (String) The name of the DB cluster parameter group.
 Constraints:
  +  Must not match the name of an existing DB cluster parameter group.
  
 If you don't specify a value for ``DBClusterParameterGroupName`` property, a name is automatically created for the DB cluster parameter group.
  This value is stored as a lowercase string.
- `tags` (Attributes List) An optional array of key-value pairs to apply to this DB cluster parameter group. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String) A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").

Optional:

- `value` (String) A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_rds_db_cluster_parameter_group.example <resource ID>
```
