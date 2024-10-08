// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_vpn_connection", vPNConnectionResource)
}

// vPNConnectionResource returns the Terraform awscc_ec2_vpn_connection resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::VPNConnection resource.
func vPNConnectionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CustomerGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the customer gateway at your end of the VPN connection.",
		//	  "type": "string"
		//	}
		"customer_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the customer gateway at your end of the VPN connection.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableAcceleration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicate whether to enable acceleration for the VPN connection.\n Default: ``false``",
		//	  "type": "boolean"
		//	}
		"enable_acceleration": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicate whether to enable acceleration for the VPN connection.\n Default: ``false``",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocalIpv4NetworkCidr
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"local_ipv_4_network_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocalIpv6NetworkCidr
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"local_ipv_6_network_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OutsideIpAddressType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"outside_ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RemoteIpv4NetworkCidr
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"remote_ipv_4_network_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RemoteIpv6NetworkCidr
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"remote_ipv_6_network_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StaticRoutesOnly
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the VPN connection uses static routes only. Static routes must be used for devices that don't support BGP.\n If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify ``true``.",
		//	  "type": "boolean"
		//	}
		"static_routes_only": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the VPN connection uses static routes only. Static routes must be used for devices that don't support BGP.\n If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify ``true``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the VPN connection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a tag. For more information, see [Resource tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag key.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the VPN connection.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway associated with the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway associated with the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TransportTransitGatewayAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"transport_transit_gateway_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TunnelInsideIpVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"tunnel_inside_ip_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of VPN connection.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of VPN connection.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpnConnectionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"vpn_connection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpnGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the virtual private gateway at the AWS side of the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
		//	  "type": "string"
		//	}
		"vpn_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the virtual private gateway at the AWS side of the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpnTunnelOptionsSpecifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tunnel options for the VPN connection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The tunnel options for a single VPN tunnel.",
		//	    "properties": {
		//	      "PreSharedKey": {
		//	        "description": "The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.\n Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).",
		//	        "type": "string"
		//	      },
		//	      "TunnelInsideCidr": {
		//	        "description": "The range of inside IP addresses for the tunnel. Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway. \n Constraints: A size /30 CIDR block from the ``169.254.0.0/16`` range. The following CIDR blocks are reserved and cannot be used:\n  +   ``169.254.0.0/30`` \n  +   ``169.254.1.0/30`` \n  +   ``169.254.2.0/30`` \n  +   ``169.254.3.0/30`` \n  +   ``169.254.4.0/30`` \n  +   ``169.254.5.0/30`` \n  +   ``169.254.169.252/30``",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"vpn_tunnel_options_specifications": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PreSharedKey
					"pre_shared_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.\n Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TunnelInsideCidr
					"tunnel_inside_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The range of inside IP addresses for the tunnel. Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway. \n Constraints: A size /30 CIDR block from the ``169.254.0.0/16`` range. The following CIDR blocks are reserved and cannot be used:\n  +   ``169.254.0.0/30`` \n  +   ``169.254.1.0/30`` \n  +   ``169.254.2.0/30`` \n  +   ``169.254.3.0/30`` \n  +   ``169.254.4.0/30`` \n  +   ``169.254.5.0/30`` \n  +   ``169.254.169.252/30``",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tunnel options for the VPN connection.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.\n To specify a VPN connection between a transit gateway and customer gateway, use the ``TransitGatewayId`` and ``CustomerGatewayId`` properties.\n To specify a VPN connection between a virtual private gateway and customer gateway, use the ``VpnGatewayId`` and ``CustomerGatewayId`` properties.\n For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPNConnection").WithTerraformTypeName("awscc_ec2_vpn_connection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"customer_gateway_id":                     "CustomerGatewayId",
		"enable_acceleration":                     "EnableAcceleration",
		"key":                                     "Key",
		"local_ipv_4_network_cidr":                "LocalIpv4NetworkCidr",
		"local_ipv_6_network_cidr":                "LocalIpv6NetworkCidr",
		"outside_ip_address_type":                 "OutsideIpAddressType",
		"pre_shared_key":                          "PreSharedKey",
		"remote_ipv_4_network_cidr":               "RemoteIpv4NetworkCidr",
		"remote_ipv_6_network_cidr":               "RemoteIpv6NetworkCidr",
		"static_routes_only":                      "StaticRoutesOnly",
		"tags":                                    "Tags",
		"transit_gateway_id":                      "TransitGatewayId",
		"transport_transit_gateway_attachment_id": "TransportTransitGatewayAttachmentId",
		"tunnel_inside_cidr":                      "TunnelInsideCidr",
		"tunnel_inside_ip_version":                "TunnelInsideIpVersion",
		"type":                                    "Type",
		"value":                                   "Value",
		"vpn_connection_id":                       "VpnConnectionId",
		"vpn_gateway_id":                          "VpnGatewayId",
		"vpn_tunnel_options_specifications":       "VpnTunnelOptionsSpecifications",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
