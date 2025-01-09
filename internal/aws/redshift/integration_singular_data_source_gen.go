// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_redshift_integration", integrationDataSource)
}

// integrationDataSource returns the Terraform awscc_redshift_integration data source.
// This Terraform data source corresponds to the CloudFormation AWS::Redshift::Integration resource.
func integrationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdditionalEncryptionContext
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An optional set of non-secret key–value pairs that contains additional contextual information about the data.",
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 131072,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"additional_encryption_context": // Pattern: ""
		schema.MapAttribute{             /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "An optional set of non-secret key–value pairs that contains additional contextual information about the data.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreateTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time (UTC) when the integration was created.",
		//	  "type": "string"
		//	}
		"create_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time (UTC) when the integration was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the integration.",
		//	  "type": "string"
		//	}
		"integration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the integration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IntegrationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the integration.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"integration_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the integration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KMSKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the database to use as the source for replication",
		//	  "type": "string"
		//	}
		"source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the database to use as the source for replication",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Redshift data warehouse to use as the target for replication",
		//	  "type": "string"
		//	}
		"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Redshift data warehouse to use as the target for replication",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Redshift::Integration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::Integration").WithTerraformTypeName("awscc_redshift_integration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_encryption_context": "AdditionalEncryptionContext",
		"create_time":                   "CreateTime",
		"integration_arn":               "IntegrationArn",
		"integration_name":              "IntegrationName",
		"key":                           "Key",
		"kms_key_id":                    "KMSKeyId",
		"source_arn":                    "SourceArn",
		"tags":                          "Tags",
		"target_arn":                    "TargetArn",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
