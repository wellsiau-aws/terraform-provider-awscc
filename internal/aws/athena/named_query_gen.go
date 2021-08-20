// Code generated by generators/resource/main.go; DO NOT EDIT.

package athena

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
	registry.AddResourceTypeFactory("awscc_athena_named_query", namedQueryResourceType)
}

// namedQueryResourceType returns the Terraform awscc_athena_named_query resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Athena::NamedQuery resource type.
func namedQueryResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"database": {
			// Property: Database
			// CloudFormation resource type schema:
			// {
			//   "description": "The database to which the query belongs.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The database to which the query belongs.",
			Type:        types.StringType,
			Required:    true,
			// Database is a force-new attribute.
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The query description.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The query description.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Description is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The query name.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The query name.",
			Type:        types.StringType,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"named_query_id": {
			// Property: NamedQueryId
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique ID of the query.",
			//   "type": "string"
			// }
			Description: "The unique ID of the query.",
			Type:        types.StringType,
			Computed:    true,
		},
		"query_string": {
			// Property: QueryString
			// CloudFormation resource type schema:
			// {
			//   "description": "The contents of the query with all query statements.",
			//   "maxLength": 262144,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The contents of the query with all query statements.",
			Type:        types.StringType,
			Required:    true,
			// QueryString is a force-new attribute.
		},
		"work_group": {
			// Property: WorkGroup
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the workgroup that contains the named query.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the workgroup that contains the named query.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// WorkGroup is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Athena::NamedQuery",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Athena::NamedQuery").WithTerraformTypeName("awscc_athena_named_query").WithTerraformSchema(schema)

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_athena_named_query", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
