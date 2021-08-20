// Code generated by generators/resource/main.go; DO NOT EDIT.

package config

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
	registry.AddResourceTypeFactory("awscc_config_organization_conformance_pack", organizationConformancePackResourceType)
}

// organizationConformancePackResourceType returns the Terraform awscc_config_organization_conformance_pack resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Config::OrganizationConformancePack resource type.
func organizationConformancePackResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"conformance_pack_input_parameters": {
			// Property: ConformancePackInputParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of ConformancePackInputParameter objects.",
			//   "items": {
			//     "description": "Input parameters in the form of key-value pairs for the conformance pack.",
			//     "properties": {
			//       "ParameterName": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       },
			//       "ParameterValue": {
			//         "maxLength": 4096,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ParameterName",
			//       "ParameterValue"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 60,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of ConformancePackInputParameter objects.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameter_name": {
						// Property: ParameterName
						Type:     types.StringType,
						Required: true,
					},
					"parameter_value": {
						// Property: ParameterValue
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 60,
				},
			),
			Optional: true,
		},
		"delivery_s3_bucket": {
			// Property: DeliveryS3Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "AWS Config stores intermediate files while processing conformance pack template.",
			//   "maxLength": 63,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "AWS Config stores intermediate files while processing conformance pack template.",
			Type:        types.StringType,
			Optional:    true,
		},
		"delivery_s3_key_prefix": {
			// Property: DeliveryS3KeyPrefix
			// CloudFormation resource type schema:
			// {
			//   "description": "The prefix for the delivery S3 bucket.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The prefix for the delivery S3 bucket.",
			Type:        types.StringType,
			Optional:    true,
		},
		"excluded_accounts": {
			// Property: ExcludedAccounts
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 1000,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"organization_conformance_pack_name": {
			// Property: OrganizationConformancePackName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the organization conformance pack.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the organization conformance pack.",
			Type:        types.StringType,
			Required:    true,
			// OrganizationConformancePackName is a force-new attribute.
		},
		"template_body": {
			// Property: TemplateBody
			// CloudFormation resource type schema:
			// {
			//   "description": "A string containing full conformance pack template body.",
			//   "maxLength": 51200,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A string containing full conformance pack template body.",
			Type:        types.StringType,
			Optional:    true,
			// TemplateBody is a write-only attribute.
		},
		"template_s3_uri": {
			// Property: TemplateS3Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "Location of file containing the template body.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Location of file containing the template body.",
			Type:        types.StringType,
			Optional:    true,
			// TemplateS3Uri is a write-only attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Config::OrganizationConformancePack.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::OrganizationConformancePack").WithTerraformTypeName("awscc_config_organization_conformance_pack").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/TemplateBody",
		"/properties/TemplateS3Uri",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_config_organization_conformance_pack", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
