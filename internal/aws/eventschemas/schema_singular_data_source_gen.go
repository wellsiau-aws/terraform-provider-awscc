// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package eventschemas

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_eventschemas_schema", schemaDataSource)
}

// schemaDataSource returns the Terraform awscc_eventschemas_schema data source.
// This Terraform data source corresponds to the CloudFormation AWS::EventSchemas::Schema resource.
func schemaDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Content
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source of the schema definition.",
		//	  "type": "string"
		//	}
		"content": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source of the schema definition.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the schema.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastModified
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The last modified time of the schema.",
		//	  "type": "string"
		//	}
		"last_modified": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The last modified time of the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RegistryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the schema registry.",
		//	  "type": "string"
		//	}
		"registry_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the schema registry.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the schema.",
		//	  "type": "string"
		//	}
		"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the schema.",
		//	  "type": "string"
		//	}
		"schema_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version number of the schema.",
		//	  "type": "string"
		//	}
		"schema_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version number of the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags associated with the resource.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags associated with the resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of schema. Valid types include OpenApi3 and JSONSchemaDraft4.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionCreatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date the schema version was created.",
		//	  "type": "string"
		//	}
		"version_created_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date the schema version was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EventSchemas::Schema",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EventSchemas::Schema").WithTerraformTypeName("awscc_eventschemas_schema")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"content":              "Content",
		"description":          "Description",
		"key":                  "Key",
		"last_modified":        "LastModified",
		"registry_name":        "RegistryName",
		"schema_arn":           "SchemaArn",
		"schema_name":          "SchemaName",
		"schema_version":       "SchemaVersion",
		"tags":                 "Tags",
		"type":                 "Type",
		"value":                "Value",
		"version_created_date": "VersionCreatedDate",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}