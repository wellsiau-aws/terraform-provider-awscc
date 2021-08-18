// Code generated by generators/resource/main.go; DO NOT EDIT.

package timestream

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
	registry.AddResourceTypeFactory("awscc_timestream_table", tableResourceType)
}

// tableResourceType returns the Terraform awscc_timestream_table resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Timestream::Table resource type.
func tableResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"database_name": {
			// Property: DatabaseName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the database which the table to be created belongs to.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the database which the table to be created belongs to.",
			Type:        types.StringType,
			Required:    true,
			// DatabaseName is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The table name exposed as a read-only attribute.",
			//   "type": "string"
			// }
			Description: "The table name exposed as a read-only attribute.",
			Type:        types.StringType,
			Computed:    true,
		},
		"retention_properties": {
			// Property: RetentionProperties
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The retention duration of the memory store and the magnetic store.",
			//   "properties": {
			//     "MagneticStoreRetentionPeriodInDays": {
			//       "description": "The duration for which data must be stored in the magnetic store.",
			//       "type": "string"
			//     },
			//     "MemoryStoreRetentionPeriodInHours": {
			//       "description": "The duration for which data must be stored in the memory store.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The retention duration of the memory store and the magnetic store.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"magnetic_store_retention_period_in_days": {
						// Property: MagneticStoreRetentionPeriodInDays
						Description: "The duration for which data must be stored in the magnetic store.",
						Type:        types.StringType,
						Optional:    true,
					},
					"memory_store_retention_period_in_hours": {
						// Property: MemoryStoreRetentionPeriodInHours
						Description: "The duration for which data must be stored in the memory store.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"table_name": {
			// Property: TableName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// TableName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "You can use the Resource Tags property to apply tags to resources, which can help you identify and categorize those resources.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array"
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 200,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::Timestream::Table resource creates a Timestream Table.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Timestream::Table").WithTerraformTypeName("awscc_timestream_table").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_timestream_table", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
