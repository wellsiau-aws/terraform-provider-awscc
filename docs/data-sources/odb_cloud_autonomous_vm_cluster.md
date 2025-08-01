---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_odb_cloud_autonomous_vm_cluster Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ODB::CloudAutonomousVmCluster
---

# awscc_odb_cloud_autonomous_vm_cluster (Data Source)

Data Source schema for AWS::ODB::CloudAutonomousVmCluster



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `autonomous_data_storage_percentage` (Number) The percentage of data storage currently in use for Autonomous Databases in the Autonomous VM cluster.
- `autonomous_data_storage_size_in_t_bs` (Number) The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB.
- `available_autonomous_data_storage_size_in_t_bs` (Number) The available data storage space for Autonomous Databases in the Autonomous VM cluster, in TB.
- `available_container_databases` (Number) The number of Autonomous CDBs that you can create with the currently available storage.
- `available_cpus` (Number) The number of CPU cores available for allocation to Autonomous Databases.
- `cloud_autonomous_vm_cluster_arn` (String) The Amazon Resource Name (ARN) for the Autonomous VM cluster.
- `cloud_autonomous_vm_cluster_id` (String) The unique identifier of the Autonomous VM cluster.
- `cloud_exadata_infrastructure_id` (String) The unique identifier of the Cloud Exadata Infrastructure containing this Autonomous VM cluster.
- `compute_model` (String) The compute model of the Autonomous VM cluster: ECPU or OCPU.
- `cpu_core_count` (Number) The total number of CPU cores in the Autonomous VM cluster.
- `cpu_core_count_per_node` (Number) The number of CPU cores enabled per node in the Autonomous VM cluster.
- `cpu_percentage` (Number) The percentage of total CPU cores currently in use in the Autonomous VM cluster.
- `data_storage_size_in_g_bs` (Number) The total data storage allocated to the Autonomous VM cluster, in GB.
- `data_storage_size_in_t_bs` (Number) The total data storage allocated to the Autonomous VM cluster, in TB.
- `db_node_storage_size_in_g_bs` (Number) The local node storage allocated to the Autonomous VM cluster, in gigabytes (GB).
- `db_servers` (List of String) The list of database servers associated with the Autonomous VM cluster.
- `description` (String) The user-provided description of the Autonomous VM cluster.
- `display_name` (String) The display name of the Autonomous VM cluster.
- `domain` (String) The domain name for the Autonomous VM cluster.
- `exadata_storage_in_t_bs_lowest_scaled_value` (Number) The minimum value to which you can scale down the Exadata storage, in TB.
- `hostname` (String) The hostname for the Autonomous VM cluster.
- `is_mtls_enabled_vm_cluster` (Boolean) Indicates whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster.
- `license_model` (String) The Oracle license model that applies to the Autonomous VM cluster. Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE.
- `maintenance_window` (Attributes) The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window. (see [below for nested schema](#nestedatt--maintenance_window))
- `max_acds_lowest_scaled_value` (Number) The minimum value to which you can scale down the maximum number of Autonomous CDBs.
- `memory_per_oracle_compute_unit_in_g_bs` (Number) The amount of memory allocated per Oracle Compute Unit, in GB.
- `memory_size_in_g_bs` (Number) The total amount of memory allocated to the Autonomous VM cluster, in gigabytes (GB).
- `node_count` (Number) The number of database server nodes in the Autonomous VM cluster.
- `non_provisionable_autonomous_container_databases` (Number) The number of Autonomous CDBs that can't be provisioned because of resource constraints.
- `oci_resource_anchor_name` (String) The name of the OCI resource anchor associated with this Autonomous VM cluster.
- `oci_url` (String) The URL for accessing the OCI console page for this Autonomous VM cluster.
- `ocid` (String) The Oracle Cloud Identifier (OCID) of the Autonomous VM cluster.
- `odb_network_id` (String) The unique identifier of the ODB network associated with this Autonomous VM cluster.
- `provisionable_autonomous_container_databases` (Number) The number of Autonomous CDBs that can be provisioned in the Autonomous VM cluster.
- `provisioned_autonomous_container_databases` (Number) The number of Autonomous CDBs currently provisioned in the Autonomous VM cluster.
- `provisioned_cpus` (Number) The number of CPU cores currently provisioned in the Autonomous VM cluster.
- `reclaimable_cpus` (Number) The number of CPU cores that can be reclaimed from terminated or scaled-down Autonomous Databases.
- `reserved_cpus` (Number) The number of CPU cores reserved for system operations and redundancy.
- `scan_listener_port_non_tls` (Number) The SCAN listener port for non-TLS (TCP) protocol. The default is 1521.
- `scan_listener_port_tls` (Number) The SCAN listener port for TLS (TCP) protocol. The default is 2484.
- `shape` (String) The shape of the Exadata infrastructure for the Autonomous VM cluster.
- `tags` (Attributes List) The tags associated with the Autonomous VM cluster. (see [below for nested schema](#nestedatt--tags))
- `time_zone` (String) The time zone of the Autonomous VM cluster.
- `total_container_databases` (Number) The total number of Autonomous Container Databases that can be created with the allocated local storage.

<a id="nestedatt--maintenance_window"></a>
### Nested Schema for `maintenance_window`

Read-Only:

- `days_of_week` (List of String) The days of the week when maintenance can be performed.
- `hours_of_day` (List of Number) The hours of the day when maintenance can be performed.
- `lead_time_in_weeks` (Number) The lead time in weeks before the maintenance window.
- `months` (List of String) The months when maintenance can be performed.
- `preference` (String) The preference for the maintenance window scheduling.
- `weeks_of_month` (List of Number) The weeks of the month when maintenance can be performed.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, @, -, and ".
- `value` (String) The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
