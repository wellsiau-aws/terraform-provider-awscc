// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_amplify_branch", branchDataSource)
}

// branchDataSource returns the Terraform awscc_amplify_branch data source.
// This Terraform data source corresponds to the CloudFormation AWS::Amplify::Branch resource.
func branchDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20,
		//	  "minLength": 1,
		//	  "pattern": "d[a-z0-9]+",
		//	  "type": "string"
		//	}
		"app_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Backend
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "StackArn": {
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"backend": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: StackArn
				"stack_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BasicAuthConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EnableBasicAuth": {
		//	      "type": "boolean"
		//	    },
		//	    "Password": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Username": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Username",
		//	    "Password"
		//	  ],
		//	  "type": "object"
		//	}
		"basic_auth_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableBasicAuth
				"enable_basic_auth": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Password
				"password": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Username
				"username": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BranchName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "(?s).+",
		//	  "type": "string"
		//	}
		"branch_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BuildSpec
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 25000,
		//	  "minLength": 1,
		//	  "pattern": "(?s).+",
		//	  "type": "string"
		//	}
		"build_spec": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1000,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnableAutoBuild
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_auto_build": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnablePerformanceMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_performance_mode": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnablePullRequestPreview
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"enable_pull_request_preview": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentVariables
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Name": {
		//	        "maxLength": 255,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 5500,
		//	        "pattern": "(?s).*",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		// Property: Framework
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"framework": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PullRequestEnvironmentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20,
		//	  "pattern": "(?s).*",
		//	  "type": "string"
		//	}
		"pull_request_environment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Stage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "EXPERIMENTAL",
		//	    "BETA",
		//	    "PULL_REQUEST",
		//	    "PRODUCTION",
		//	    "DEVELOPMENT"
		//	  ],
		//	  "type": "string"
		//	}
		"stage": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "insertionOrder": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
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
		Description: "Data Source schema for AWS::Amplify::Branch",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Branch").WithTerraformTypeName("awscc_amplify_branch")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                        "AppId",
		"arn":                           "Arn",
		"backend":                       "Backend",
		"basic_auth_config":             "BasicAuthConfig",
		"branch_name":                   "BranchName",
		"build_spec":                    "BuildSpec",
		"description":                   "Description",
		"enable_auto_build":             "EnableAutoBuild",
		"enable_basic_auth":             "EnableBasicAuth",
		"enable_performance_mode":       "EnablePerformanceMode",
		"enable_pull_request_preview":   "EnablePullRequestPreview",
		"environment_variables":         "EnvironmentVariables",
		"framework":                     "Framework",
		"key":                           "Key",
		"name":                          "Name",
		"password":                      "Password",
		"pull_request_environment_name": "PullRequestEnvironmentName",
		"stack_arn":                     "StackArn",
		"stage":                         "Stage",
		"tags":                          "Tags",
		"username":                      "Username",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
