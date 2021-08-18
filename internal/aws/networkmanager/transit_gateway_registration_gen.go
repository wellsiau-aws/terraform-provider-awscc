// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkmanager

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
	registry.AddResourceTypeFactory("awscc_networkmanager_transit_gateway_registration", transitGatewayRegistrationResourceType)
}

// transitGatewayRegistrationResourceType returns the Terraform awscc_networkmanager_transit_gateway_registration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkManager::TransitGatewayRegistration resource type.
func transitGatewayRegistrationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"global_network_id": {
			// Property: GlobalNetworkId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the global network.",
			//   "type": "string"
			// }
			Description: "The ID of the global network.",
			Type:        types.StringType,
			Required:    true,
			// GlobalNetworkId is a force-new attribute.
		},
		"transit_gateway_arn": {
			// Property: TransitGatewayArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the transit gateway.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the transit gateway.",
			Type:        types.StringType,
			Required:    true,
			// TransitGatewayArn is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::NetworkManager::TransitGatewayRegistration type registers a transit gateway in your global network. The transit gateway can be in any AWS Region, but it must be owned by the same AWS account that owns the global network. You cannot register a transit gateway in more than one global network.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::TransitGatewayRegistration").WithTerraformTypeName("awscc_networkmanager_transit_gateway_registration").WithTerraformSchema(schema)

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_networkmanager_transit_gateway_registration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
