// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

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
	registry.AddResourceTypeFactory("awscc_cloudformation_resource_default_version", resourceDefaultVersionResourceType)
}

// resourceDefaultVersionResourceType returns the Terraform awscc_cloudformation_resource_default_version resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFormation::ResourceDefaultVersion resource type.
func resourceDefaultVersionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion",
			Type:        types.StringType,
			Computed:    true,
		},
		"type_name": {
			// Property: TypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Type:        types.StringType,
			Optional:    true,
		},
		"type_version_arn": {
			// Property: TypeVersionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the type version.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the type version.",
			Type:        types.StringType,
			Optional:    true,
		},
		"version_id": {
			// Property: VersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of an existing version of the resource to set as the default.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of an existing version of the resource to set as the default.",
			Type:        types.StringType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The default version of a resource that has been registered in the CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::ResourceDefaultVersion").WithTerraformTypeName("awscc_cloudformation_resource_default_version").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_cloudformation_resource_default_version", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
