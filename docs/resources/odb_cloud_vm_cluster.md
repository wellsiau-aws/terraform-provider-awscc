---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_odb_cloud_vm_cluster Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::ODB::CloudVmCluster resource creates a Cloud VM Cluster
---

# awscc_odb_cloud_vm_cluster (Resource)

The AWS::ODB::CloudVmCluster resource creates a Cloud VM Cluster



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cloud_exadata_infrastructure_id` (String) The unique identifier of the Exadata infrastructure that this VM cluster belongs to.
- `cluster_name` (String) The name of the Grid Infrastructure (GI) cluster.
- `cpu_core_count` (Number) The number of CPU cores enabled on the VM cluster.
- `data_collection_options` (Attributes) The set of diagnostic collection options enabled for the VM cluster. (see [below for nested schema](#nestedatt--data_collection_options))
- `data_storage_size_in_t_bs` (Number) The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster.
- `db_node_storage_size_in_g_bs` (Number) The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster.
- `db_servers` (List of String) The list of database servers for the VM cluster.
- `display_name` (String) The user-friendly name for the VM cluster.
- `gi_version` (String) The software version of the Oracle Grid Infrastructure (GI) for the VM cluster.
- `hostname` (String) The host name for the VM cluster.
- `is_local_backup_enabled` (Boolean) Indicates whether database backups to local Exadata storage is enabled for the VM cluster.
- `is_sparse_diskgroup_enabled` (Boolean) Indicates whether the VM cluster is configured with a sparse disk group.
- `license_model` (String) The Oracle license model applied to the VM cluster.
- `memory_size_in_g_bs` (Number) The amount of memory, in gigabytes (GB), that's allocated for the VM cluster.
- `odb_network_id` (String) The unique identifier of the ODB network for the VM cluster.
- `scan_listener_port_tcp` (Number) Property description not available.
- `ssh_public_keys` (List of String) The public key portion of one or more key pairs used for SSH access to the VM cluster.
- `system_version` (String) The operating system version of the image chosen for the VM cluster.
- `tags` (Attributes List) Tags to assign to the Vm Cluster. (see [below for nested schema](#nestedatt--tags))
- `time_zone` (String) The time zone of the VM cluster.

### Read-Only

- `cloud_vm_cluster_arn` (String) The Amazon Resource Name (ARN) of the VM cluster.
- `cloud_vm_cluster_id` (String) The unique identifier of the VM cluster.
- `compute_model` (String) The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled.
- `disk_redundancy` (String) The type of redundancy configured for the VM cluster. NORMAL is 2-way redundancy. HIGH is 3-way redundancy.
- `domain` (String) The domain of the VM cluster.
- `id` (String) Uniquely identifies the resource.
- `listener_port` (Number) The port number configured for the listener on the VM cluster.
- `node_count` (Number) The number of nodes in the VM cluster.
- `oci_resource_anchor_name` (String) The name of the OCI resource anchor for the VM cluster.
- `oci_url` (String) The HTTPS link to the VM cluster in OCI.
- `ocid` (String) The OCID of the VM cluster.
- `scan_dns_name` (String) The FQDN of the DNS record for the Single Client Access Name (SCAN) IP addresses that are associated with the VM cluster.
- `scan_ip_ids` (List of String) The OCID of the SCAN IP addresses that are associated with the VM cluster.
- `shape` (String) The hardware model name of the Exadata infrastructure that's running the VM cluster.
- `storage_size_in_g_bs` (Number) The amount of local node storage, in gigabytes (GB), that's allocated to the VM cluster.
- `vip_ids` (List of String) The virtual IP (VIP) addresses that are associated with the VM cluster. Oracle's Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the VM cluster to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.

<a id="nestedatt--data_collection_options"></a>
### Nested Schema for `data_collection_options`

Optional:

- `is_diagnostics_events_enabled` (Boolean) Indicates whether diagnostic collection is enabled for the VM cluster.
- `is_health_monitoring_enabled` (Boolean) Indicates whether health monitoring is enabled for the VM cluster.
- `is_incident_logs_enabled` (Boolean) Indicates whether incident logs are enabled for the cloud VM cluster.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, @, -, and ".
- `value` (String) The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_odb_cloud_vm_cluster.example
  id = "cloud_vm_cluster_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_odb_cloud_vm_cluster.example "cloud_vm_cluster_arn"
```
