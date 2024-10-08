// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lambda_code_signing_config", codeSigningConfigDataSource)
}

// codeSigningConfigDataSource returns the Terraform awscc_lambda_code_signing_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lambda::CodeSigningConfig resource.
func codeSigningConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllowedPublishers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
		//	  "properties": {
		//	    "SigningProfileVersionArns": {
		//	      "description": "List of Signing profile version Arns",
		//	      "items": {
		//	        "maxLength": 1024,
		//	        "minLength": 12,
		//	        "pattern": "arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 20,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "SigningProfileVersionArns"
		//	  ],
		//	  "type": "object"
		//	}
		"allowed_publishers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SigningProfileVersionArns
				"signing_profile_version_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "List of Signing profile version Arns",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CodeSigningConfigArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique Arn for CodeSigningConfig resource",
		//	  "pattern": "arn:(aws[a-zA-Z-]*)?:lambda:[a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\\d{1}:\\d{12}:code-signing-config:csc-[a-z0-9]{17}",
		//	  "type": "string"
		//	}
		"code_signing_config_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique Arn for CodeSigningConfig resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CodeSigningConfigId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for CodeSigningConfig resource",
		//	  "pattern": "csc-[a-zA-Z0-9-_\\.]{17}",
		//	  "type": "string"
		//	}
		"code_signing_config_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for CodeSigningConfig resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CodeSigningPolicies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Policies to control how to act if a signature is invalid",
		//	  "properties": {
		//	    "UntrustedArtifactOnDeployment": {
		//	      "default": "Warn",
		//	      "description": "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
		//	      "enum": [
		//	        "Warn",
		//	        "Enforce"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "UntrustedArtifactOnDeployment"
		//	  ],
		//	  "type": "object"
		//	}
		"code_signing_policies": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: UntrustedArtifactOnDeployment
				"untrusted_artifact_on_deployment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Policies to control how to act if a signature is invalid",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the CodeSigningConfig",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the CodeSigningConfig",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags to apply to CodeSigningConfig resource",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
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
			Description: "A list of tags to apply to CodeSigningConfig resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lambda::CodeSigningConfig",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::CodeSigningConfig").WithTerraformTypeName("awscc_lambda_code_signing_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_publishers":               "AllowedPublishers",
		"code_signing_config_arn":          "CodeSigningConfigArn",
		"code_signing_config_id":           "CodeSigningConfigId",
		"code_signing_policies":            "CodeSigningPolicies",
		"description":                      "Description",
		"key":                              "Key",
		"signing_profile_version_arns":     "SigningProfileVersionArns",
		"tags":                             "Tags",
		"untrusted_artifact_on_deployment": "UntrustedArtifactOnDeployment",
		"value":                            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
