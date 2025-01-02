// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package pcs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_pcs_queue", queueDataSource)
}

// queueDataSource returns the Terraform awscc_pcs_queue data source.
// This Terraform data source corresponds to the CloudFormation AWS::PCS::Queue resource.
func queueDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique Amazon Resource Name (ARN) of the queue.",
		//	  "pattern": "^(.*?)",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique Amazon Resource Name (ARN) of the queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the cluster of the queue.",
		//	  "type": "string"
		//	}
		"cluster_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the cluster of the queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ComputeNodeGroupConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of compute node group configurations associated with the queue. Queues assign jobs to associated compute node groups.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The compute node group configuration for a queue.",
		//	    "properties": {
		//	      "ComputeNodeGroupId": {
		//	        "description": "The compute node group ID for the compute node group configuration.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"compute_node_group_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ComputeNodeGroupId
					"compute_node_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The compute node group ID for the compute node group configuration.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of compute node group configurations associated with the queue. Queues assign jobs to associated compute node groups.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ErrorInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of errors that occurred during queue provisioning.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An error that occurred during resource provisioning.",
		//	    "properties": {
		//	      "Code": {
		//	        "description": "The short-form error code.",
		//	        "type": "string"
		//	      },
		//	      "Message": {
		//	        "description": "The detailed error information.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"error_info": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Code
					"code": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The short-form error code.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Message
					"message": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The detailed error information.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of errors that occurred during queue provisioning.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The generated unique ID of the queue.",
		//	  "type": "string"
		//	}
		"queue_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The generated unique ID of the queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name that identifies the queue.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name that identifies the queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provisioning status of the queue. The provisioning status doesn't indicate the overall health of the queue.",
		//	  "enum": [
		//	    "CREATING",
		//	    "ACTIVE",
		//	    "UPDATING",
		//	    "DELETING",
		//	    "CREATE_FAILED",
		//	    "DELETE_FAILED",
		//	    "UPDATE_FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provisioning status of the queue. The provisioning status doesn't indicate the overall health of the queue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::PCS::Queue",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PCS::Queue").WithTerraformTypeName("awscc_pcs_queue")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                               "Arn",
		"cluster_id":                        "ClusterId",
		"code":                              "Code",
		"compute_node_group_configurations": "ComputeNodeGroupConfigurations",
		"compute_node_group_id":             "ComputeNodeGroupId",
		"error_info":                        "ErrorInfo",
		"message":                           "Message",
		"name":                              "Name",
		"queue_id":                          "Id",
		"status":                            "Status",
		"tags":                              "Tags",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}