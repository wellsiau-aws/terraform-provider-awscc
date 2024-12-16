// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkmanager_direct_connect_gateway_attachment", directConnectGatewayAttachmentDataSource)
}

// directConnectGatewayAttachmentDataSource returns the Terraform awscc_networkmanager_direct_connect_gateway_attachment data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkManager::DirectConnectGatewayAttachment resource.
func directConnectGatewayAttachmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the attachment.",
		//	  "type": "string"
		//	}
		"attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AttachmentPolicyRuleNumber
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The policy rule number associated with the attachment.",
		//	  "type": "integer"
		//	}
		"attachment_policy_rule_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The policy rule number associated with the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AttachmentType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Attachment type.",
		//	  "type": "string"
		//	}
		"attachment_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Attachment type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of a core network for the Direct Connect Gateway attachment.",
		//	  "type": "string"
		//	}
		"core_network_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of a core network for the Direct Connect Gateway attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of a core network for the Direct Connect Gateway attachment.",
		//	  "type": "string"
		//	}
		"core_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of a core network for the Direct Connect Gateway attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creation time of the attachment.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Creation time of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DirectConnectGatewayArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Direct Connect Gateway.",
		//	  "type": "string"
		//	}
		"direct_connect_gateway_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Direct Connect Gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EdgeLocations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Regions where the edges are located.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"edge_locations": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Regions where the edges are located.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkFunctionGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the network function group attachment.",
		//	  "type": "string"
		//	}
		"network_function_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the network function group attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Owner account of the attachment.",
		//	  "type": "string"
		//	}
		"owner_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Owner account of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProposedNetworkFunctionGroupChange
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The attachment to move from one network function group to another.",
		//	  "properties": {
		//	    "AttachmentPolicyRuleNumber": {
		//	      "description": "The rule number in the policy document that applies to this change.",
		//	      "type": "integer"
		//	    },
		//	    "NetworkFunctionGroupName": {
		//	      "description": "The name of the network function group to change.",
		//	      "type": "string"
		//	    },
		//	    "Tags": {
		//	      "description": "The key-value tags that changed for the network function group.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A key-value pair to associate with a resource.",
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"proposed_network_function_group_change": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AttachmentPolicyRuleNumber
				"attachment_policy_rule_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The rule number in the policy document that applies to this change.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: NetworkFunctionGroupName
				"network_function_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the network function group to change.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The key-value tags that changed for the network function group.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The attachment to move from one network function group to another.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProposedSegmentChange
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The attachment to move from one segment to another.",
		//	  "properties": {
		//	    "AttachmentPolicyRuleNumber": {
		//	      "description": "The rule number in the policy document that applies to this change.",
		//	      "type": "integer"
		//	    },
		//	    "SegmentName": {
		//	      "description": "The name of the segment to change.",
		//	      "type": "string"
		//	    },
		//	    "Tags": {
		//	      "description": "The key-value tags that changed for the segment.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A key-value pair to associate with a resource.",
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"proposed_segment_change": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AttachmentPolicyRuleNumber
				"attachment_policy_rule_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The rule number in the policy document that applies to this change.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SegmentName
				"segment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the segment to change.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "The key-value tags that changed for the segment.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The attachment to move from one segment to another.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Resource.",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SegmentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the segment attachment..",
		//	  "type": "string"
		//	}
		"segment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the segment attachment..",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "State of the attachment.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags for the attachment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags for the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Last update time of the attachment.",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Last update time of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::NetworkManager::DirectConnectGatewayAttachment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::DirectConnectGatewayAttachment").WithTerraformTypeName("awscc_networkmanager_direct_connect_gateway_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachment_id":                          "AttachmentId",
		"attachment_policy_rule_number":          "AttachmentPolicyRuleNumber",
		"attachment_type":                        "AttachmentType",
		"core_network_arn":                       "CoreNetworkArn",
		"core_network_id":                        "CoreNetworkId",
		"created_at":                             "CreatedAt",
		"direct_connect_gateway_arn":             "DirectConnectGatewayArn",
		"edge_locations":                         "EdgeLocations",
		"key":                                    "Key",
		"network_function_group_name":            "NetworkFunctionGroupName",
		"owner_account_id":                       "OwnerAccountId",
		"proposed_network_function_group_change": "ProposedNetworkFunctionGroupChange",
		"proposed_segment_change":                "ProposedSegmentChange",
		"resource_arn":                           "ResourceArn",
		"segment_name":                           "SegmentName",
		"state":                                  "State",
		"tags":                                   "Tags",
		"updated_at":                             "UpdatedAt",
		"value":                                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
