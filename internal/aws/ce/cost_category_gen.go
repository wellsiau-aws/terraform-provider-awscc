// Code generated by generators/resource/main.go; DO NOT EDIT.

package ce

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
	registry.AddResourceTypeFactory("awscc_ce_cost_category", costCategoryResourceType)
}

// costCategoryResourceType returns the Terraform awscc_ce_cost_category resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CE::CostCategory resource type.
func costCategoryResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Cost category ARN",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Cost category ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"default_value": {
			// Property: DefaultValue
			// CloudFormation resource type schema:
			// {
			//   "description": "The default value for the cost category",
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The default value for the cost category",
			Type:        types.StringType,
			Optional:    true,
		},
		"effective_start": {
			// Property: EffectiveStart
			// CloudFormation resource type schema:
			// {
			//   "description": "ISO 8601 date time with offset format",
			//   "maxLength": 25,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "ISO 8601 date time with offset format",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// Name is a force-new attribute.
		},
		"rule_version": {
			// Property: RuleVersion
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "CostCategoryExpression.v1"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"rules": {
			// Property: Rules
			// CloudFormation resource type schema:
			// {
			//   "description": "JSON array format of Expression in Billing and Cost Management API",
			//   "type": "string"
			// }
			Description: "JSON array format of Expression in Billing and Cost Management API",
			Type:        types.StringType,
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::CostCategory").WithTerraformTypeName("awscc_ce_cost_category").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ce_cost_category", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}