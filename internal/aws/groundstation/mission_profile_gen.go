// Code generated by generators/resource/main.go; DO NOT EDIT.

package groundstation

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_groundstation_mission_profile", missionProfileResourceType)
}

// missionProfileResourceType returns the Terraform awscc_groundstation_mission_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GroundStation::MissionProfile resource type.
func missionProfileResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"contact_post_pass_duration_seconds": {
			// Property: ContactPostPassDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Post-pass time needed after the contact.",
			//   "type": "integer"
			// }
			Description: "Post-pass time needed after the contact.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"contact_pre_pass_duration_seconds": {
			// Property: ContactPrePassDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Pre-pass time needed before the contact.",
			//   "type": "integer"
			// }
			Description: "Pre-pass time needed before the contact.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"dataflow_edges": {
			// Property: DataflowEdges
			// CloudFormation resource type schema:
			// {
			//   "description": "",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Destination": {
			//         "type": "string"
			//       },
			//       "Source": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"destination": {
						// Property: Destination
						Type:     types.StringType,
						Optional: true,
					},
					"source": {
						// Property: Source
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
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
		"minimum_viable_contact_duration_seconds": {
			// Property: MinimumViableContactDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
			//   "type": "integer"
			// }
			Description: "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
			Type:        types.NumberType,
			Required:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name used to identify a mission profile.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A name used to identify a mission profile.",
			Type:        types.StringType,
			Required:    true,
		},
		"region": {
			// Property: Region
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
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
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"tracking_config_arn": {
			// Property: TrackingConfigArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "AWS Ground Station Mission Profile resource type for CloudFormation.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::MissionProfile").WithTerraformTypeName("awscc_groundstation_mission_profile").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_groundstation_mission_profile", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}