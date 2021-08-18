// Code generated by generators/resource/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_redshift_cluster", clusterResourceType)
}

// clusterResourceType returns the Terraform awscc_redshift_cluster resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Redshift::Cluster resource type.
func clusterResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"allow_version_upgrade": {
			// Property: AllowVersionUpgrade
			// CloudFormation resource type schema:
			// {
			//   "description": "Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True",
			//   "type": "boolean"
			// }
			Description: "Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True",
			Type:        types.BoolType,
			Optional:    true,
		},
		"automated_snapshot_retention_period": {
			// Property: AutomatedSnapshotRetentionPeriod
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1",
			//   "type": "integer"
			// }
			Description: "The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1",
			Type:        types.NumberType,
			Optional:    true,
		},
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			// {
			//   "description": "The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint",
			//   "type": "string"
			// }
			Description: "The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// AvailabilityZone is a force-new attribute.
		},
		"cluster_identifier": {
			// Property: ClusterIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
			//   "maxLength": 63,
			//   "type": "string"
			// }
			Description: "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ClusterIdentifier is a force-new attribute.
		},
		"cluster_parameter_group_name": {
			// Property: ClusterParameterGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the parameter group to be associated with this cluster.",
			//   "maxLength": 255,
			//   "type": "string"
			// }
			Description: "The name of the parameter group to be associated with this cluster.",
			Type:        types.StringType,
			Optional:    true,
		},
		"cluster_security_groups": {
			// Property: ClusterSecurityGroups
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of security groups to be associated with this cluster.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of security groups to be associated with this cluster.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"cluster_subnet_group_name": {
			// Property: ClusterSubnetGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of a cluster subnet group to be associated with this cluster.",
			//   "type": "string"
			// }
			Description: "The name of a cluster subnet group to be associated with this cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ClusterSubnetGroupName is a force-new attribute.
		},
		"cluster_type": {
			// Property: ClusterType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required",
			//   "type": "string"
			// }
			Description: "The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required",
			Type:        types.StringType,
			Required:    true,
		},
		"cluster_version": {
			// Property: ClusterVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.",
			//   "type": "string"
			// }
			Description: "The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.",
			Type:        types.StringType,
			Optional:    true,
		},
		"db_name": {
			// Property: DBName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.",
			//   "type": "string"
			// }
			Description: "The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.",
			Type:        types.StringType,
			Required:    true,
			// DBName is a force-new attribute.
		},
		"elastic_ip": {
			// Property: ElasticIp
			// CloudFormation resource type schema:
			// {
			//   "description": "The Elastic IP (EIP) address for the cluster.",
			//   "type": "string"
			// }
			Description: "The Elastic IP (EIP) address for the cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ElasticIp is a force-new attribute.
		},
		"encrypted": {
			// Property: Encrypted
			// CloudFormation resource type schema:
			// {
			//   "description": "If true, the data in the cluster is encrypted at rest.",
			//   "type": "boolean"
			// }
			Description: "If true, the data in the cluster is encrypted at rest.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			// Encrypted is a force-new attribute.
		},
		"endpoint": {
			// Property: Endpoint
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "Address": {
			//       "type": "string"
			//     },
			//     "Port": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"address": {
						// Property: Address
						Type:     types.StringType,
						Computed: true,
					},
					"port": {
						// Property: Port
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Optional: true,
		},
		"hsm_client_certificate_identifier": {
			// Property: HsmClientCertificateIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM",
			//   "type": "string"
			// }
			Description: "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM",
			Type:        types.StringType,
			Optional:    true,
		},
		"hsm_configuration_identifier": {
			// Property: HsmConfigurationIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.",
			//   "type": "string"
			// }
			Description: "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.",
			Type:        types.StringType,
			Optional:    true,
		},
		"iam_roles": {
			// Property: IamRoles
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 10 IAM roles in a single request",
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 10,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 10 IAM roles in a single request",
			// Ordered set.
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.",
			//   "type": "string"
			// }
			Description: "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// KmsKeyId is a force-new attribute.
		},
		"logging_properties": {
			// Property: LoggingProperties
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BucketName": {
			//       "type": "string"
			//     },
			//     "S3KeyPrefix": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "BucketName"
			//   ],
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"bucket_name": {
						// Property: BucketName
						Type:     types.StringType,
						Required: true,
					},
					"s3_key_prefix": {
						// Property: S3KeyPrefix
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"master_user_password": {
			// Property: MasterUserPassword
			// CloudFormation resource type schema:
			// {
			//   "description": "The password associated with the master user account for the cluster that is being created. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.",
			//   "maxLength": 64,
			//   "type": "string"
			// }
			Description: "The password associated with the master user account for the cluster that is being created. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.",
			Type:        types.StringType,
			Required:    true,
			// MasterUserPassword is a write-only attribute.
		},
		"master_username": {
			// Property: MasterUsername
			// CloudFormation resource type schema:
			// {
			//   "description": "The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter.",
			//   "maxLength": 128,
			//   "type": "string"
			// }
			Description: "The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter.",
			Type:        types.StringType,
			Required:    true,
			// MasterUsername is a force-new attribute.
		},
		"node_type": {
			// Property: NodeType
			// CloudFormation resource type schema:
			// {
			//   "description": "The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.4xlarge | ra3.16xlarge",
			//   "type": "string"
			// }
			Description: "The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.4xlarge | ra3.16xlarge",
			Type:        types.StringType,
			Required:    true,
		},
		"number_of_nodes": {
			// Property: NumberOfNodes
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.",
			//   "type": "integer"
			// }
			Description: "The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"owner_account": {
			// Property: OwnerAccount
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// OwnerAccount is a force-new attribute.
		},
		"port": {
			// Property: Port
			// CloudFormation resource type schema:
			// {
			//   "description": "The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings",
			//   "type": "integer"
			// }
			Description: "The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings",
			Type:        types.NumberType,
			Optional:    true,
			Computed:    true,
			// Port is a force-new attribute.
		},
		"preferred_maintenance_window": {
			// Property: PreferredMaintenanceWindow
			// CloudFormation resource type schema:
			// {
			//   "description": "The weekly time range (in UTC) during which automated cluster maintenance can occur.",
			//   "type": "string"
			// }
			Description: "The weekly time range (in UTC) during which automated cluster maintenance can occur.",
			Type:        types.StringType,
			Optional:    true,
		},
		"publicly_accessible": {
			// Property: PubliclyAccessible
			// CloudFormation resource type schema:
			// {
			//   "description": "If true, the cluster can be accessed from a public network.",
			//   "type": "boolean"
			// }
			Description: "If true, the cluster can be accessed from a public network.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"snapshot_cluster_identifier": {
			// Property: SnapshotClusterIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.",
			//   "type": "string"
			// }
			Description: "The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// SnapshotClusterIdentifier is a force-new attribute.
		},
		"snapshot_identifier": {
			// Property: SnapshotIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.",
			//   "type": "string"
			// }
			Description: "The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// SnapshotIdentifier is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of tags for the cluster parameter group.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The list of tags for the cluster parameter group.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"vpc_security_group_ids": {
			// Property: VpcSecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::Cluster").WithTerraformTypeName("awscc_redshift_cluster").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/MasterUserPassword",
	})
	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(2160)

	opts = opts.WithUpdateTimeoutInMinutes(2160)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_redshift_cluster", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
