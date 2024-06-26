// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotanalytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotanalytics_pipeline", pipelineDataSource)
}

// pipelineDataSource returns the Terraform awscc_iotanalytics_pipeline data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTAnalytics::Pipeline resource.
func pipelineDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"pipeline_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PipelineActivities
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AddAttributes": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attributes": {
		//	            "additionalProperties": false,
		//	            "patternProperties": {
		//	              "": {
		//	                "maxLength": 256,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Attributes",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Channel": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ChannelName": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "[a-zA-Z0-9_]+",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "ChannelName",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Datastore": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "DatastoreName": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "[a-zA-Z0-9_]+",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "DatastoreName",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "DeviceRegistryEnrich": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attribute": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "RoleArn": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "type": "string"
		//	          },
		//	          "ThingName": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Attribute",
		//	          "ThingName",
		//	          "RoleArn",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "DeviceShadowEnrich": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attribute": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "RoleArn": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "type": "string"
		//	          },
		//	          "ThingName": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Attribute",
		//	          "ThingName",
		//	          "RoleArn",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Filter": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Filter": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Filter",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Lambda": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "BatchSize": {
		//	            "maximum": 1000,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "LambdaName": {
		//	            "maxLength": 64,
		//	            "minLength": 1,
		//	            "pattern": "[a-zA-Z0-9_-]+",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "LambdaName",
		//	          "Name",
		//	          "BatchSize"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Math": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attribute": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Math": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Attribute",
		//	          "Math",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "RemoveAttributes": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attributes": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 50,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Attributes",
		//	          "Name"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "SelectAttributes": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attributes": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 50,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "Name": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Next": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Attributes"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 25,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"pipeline_activities": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AddAttributes
					"add_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attributes
							"attributes":        // Pattern: ""
							schema.MapAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Channel
					"channel": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ChannelName
							"channel_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Datastore
					"datastore": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DatastoreName
							"datastore_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DeviceRegistryEnrich
					"device_registry_enrich": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attribute
							"attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: RoleArn
							"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: ThingName
							"thing_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DeviceShadowEnrich
					"device_shadow_enrich": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attribute
							"attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: RoleArn
							"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: ThingName
							"thing_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Filter
					"filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Filter
							"filter": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Lambda
					"lambda": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: BatchSize
							"batch_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: LambdaName
							"lambda_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Math
					"math": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attribute
							"attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Math
							"math": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: RemoveAttributes
					"remove_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attributes
							"attributes": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SelectAttributes
					"select_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attributes
							"attributes": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Next
							"next": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PipelineName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_]+",
		//	  "type": "string"
		//	}
		"pipeline_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTAnalytics::Pipeline",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTAnalytics::Pipeline").WithTerraformTypeName("awscc_iotanalytics_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_attributes":         "AddAttributes",
		"attribute":              "Attribute",
		"attributes":             "Attributes",
		"batch_size":             "BatchSize",
		"channel":                "Channel",
		"channel_name":           "ChannelName",
		"datastore":              "Datastore",
		"datastore_name":         "DatastoreName",
		"device_registry_enrich": "DeviceRegistryEnrich",
		"device_shadow_enrich":   "DeviceShadowEnrich",
		"filter":                 "Filter",
		"key":                    "Key",
		"lambda":                 "Lambda",
		"lambda_name":            "LambdaName",
		"math":                   "Math",
		"name":                   "Name",
		"next":                   "Next",
		"pipeline_activities":    "PipelineActivities",
		"pipeline_id":            "Id",
		"pipeline_name":          "PipelineName",
		"remove_attributes":      "RemoveAttributes",
		"role_arn":               "RoleArn",
		"select_attributes":      "SelectAttributes",
		"tags":                   "Tags",
		"thing_name":             "ThingName",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
