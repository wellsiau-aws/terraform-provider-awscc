// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_transit_gateway_route", transitGatewayRouteDataSource)
}

// transitGatewayRouteDataSource returns the Terraform awscc_ec2_transit_gateway_route data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::TransitGatewayRoute resource.
func transitGatewayRouteDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Blackhole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to drop traffic that matches this route.",
		//	  "type": "boolean"
		//	}
		"blackhole": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to drop traffic that matches this route.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationCidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The CIDR range used for destination matches. Routing decisions are based on the most specific match.",
		//	  "type": "string"
		//	}
		"destination_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The CIDR range used for destination matches. Routing decisions are based on the most specific match.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of transit gateway attachment.",
		//	  "type": "string"
		//	}
		"transit_gateway_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of transit gateway attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayRouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of transit gateway route table.",
		//	  "type": "string"
		//	}
		"transit_gateway_route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of transit gateway route table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::TransitGatewayRoute",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayRoute").WithTerraformTypeName("awscc_ec2_transit_gateway_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"blackhole":                      "Blackhole",
		"destination_cidr_block":         "DestinationCidrBlock",
		"transit_gateway_attachment_id":  "TransitGatewayAttachmentId",
		"transit_gateway_route_table_id": "TransitGatewayRouteTableId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
