// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

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
	registry.AddResourceTypeFactory("awscc_ses_contact_list", contactListResourceType)
}

// contactListResourceType returns the Terraform awscc_ses_contact_list resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SES::ContactList resource type.
func contactListResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"contact_list_name": {
			// Property: ContactListName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the contact list.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the contact list.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ContactListName is a force-new attribute.
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the contact list.",
			//   "maxLength": 500,
			//   "type": "string"
			// }
			Description: "The description of the contact list.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the contact list.",
			//   "items": {
			//     "additionalProperties": false,
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
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the contact list.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"topics": {
			// Property: Topics
			// CloudFormation resource type schema:
			// {
			//   "description": "The topics associated with the contact list.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "DefaultSubscriptionStatus": {
			//         "type": "string"
			//       },
			//       "Description": {
			//         "description": "The description of the topic.",
			//         "maxLength": 500,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "DisplayName": {
			//         "description": "The display name of the topic.",
			//         "maxLength": 128,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "TopicName": {
			//         "description": "The name of the topic.",
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "TopicName",
			//       "DisplayName",
			//       "DefaultSubscriptionStatus"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The topics associated with the contact list.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"default_subscription_status": {
						// Property: DefaultSubscriptionStatus
						Type:     types.StringType,
						Required: true,
					},
					"description": {
						// Property: Description
						Description: "The description of the topic.",
						Type:        types.StringType,
						Optional:    true,
					},
					"display_name": {
						// Property: DisplayName
						Description: "The display name of the topic.",
						Type:        types.StringType,
						Required:    true,
					},
					"topic_name": {
						// Property: TopicName
						Description: "The name of the topic.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 20,
				},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::SES::ContactList.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::ContactList").WithTerraformTypeName("awscc_ses_contact_list").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ses_contact_list", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}