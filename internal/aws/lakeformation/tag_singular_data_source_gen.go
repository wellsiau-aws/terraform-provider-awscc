// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lakeformation_tag", tagDataSource)
}

// tagDataSource returns the Terraform awscc_lakeformation_tag data source.
// This Terraform data source corresponds to the CloudFormation AWS::LakeFormation::Tag resource.
func tagDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CatalogId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TagKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The key-name for the LF-tag.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^([{a-zA-Z}{\\s}{0-9}_.:\\/=+\\-@%]*)$",
		//	  "type": "string"
		//	}
		"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The key-name for the LF-tag.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TagValues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of possible values an attribute can take.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 256,
		//	    "minLength": 0,
		//	    "pattern": "^([{a-zA-Z}{\\s}{0-9}_.:\\*\\/=+\\-@%]*)$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1000,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of possible values an attribute can take.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::LakeFormation::Tag",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::Tag").WithTerraformTypeName("awscc_lakeformation_tag")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog_id": "CatalogId",
		"tag_key":    "TagKey",
		"tag_values": "TagValues",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
