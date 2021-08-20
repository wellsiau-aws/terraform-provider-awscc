// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

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
	registry.AddResourceTypeFactory("awscc_ecs_primary_task_set", primaryTaskSetResourceType)
}

// primaryTaskSetResourceType returns the Terraform awscc_ecs_primary_task_set resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ECS::PrimaryTaskSet resource type.
func primaryTaskSetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cluster": {
			// Property: Cluster
			// CloudFormation resource type schema:
			// {
			//   "description": "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			//   "type": "string"
			// }
			Description: "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			Type:        types.StringType,
			Required:    true,
			// Cluster is a force-new attribute.
		},
		"service": {
			// Property: Service
			// CloudFormation resource type schema:
			// {
			//   "description": "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			//   "type": "string"
			// }
			Description: "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			Type:        types.StringType,
			Required:    true,
			// Service is a force-new attribute.
		},
		"task_set_id": {
			// Property: TaskSetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID or full Amazon Resource Name (ARN) of the task set.",
			//   "type": "string"
			// }
			Description: "The ID or full Amazon Resource Name (ARN) of the task set.",
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
		Description: "A pseudo-resource that manages which of your ECS task sets is primary.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::PrimaryTaskSet").WithTerraformTypeName("awscc_ecs_primary_task_set").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ecs_primary_task_set", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
