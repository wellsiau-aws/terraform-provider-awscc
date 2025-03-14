// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_documentation_part", documentationPartDataSource)
}

// documentationPartDataSource returns the Terraform awscc_apigateway_documentation_part data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::DocumentationPart resource.
func documentationPartDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DocumentationPartId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"documentation_part_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The ``Location`` property specifies the location of the Amazon API Gateway API entity that the documentation applies to. ``Location`` is a property of the [AWS::ApiGateway::DocumentationPart](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html) resource.\n For more information about each property, including constraints and valid values, see [DocumentationPart](https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationPartLocation.html) in the *Amazon API Gateway REST API Reference*.",
		//	  "properties": {
		//	    "Method": {
		//	      "description": "",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "description": "",
		//	      "type": "string"
		//	    },
		//	    "Path": {
		//	      "description": "",
		//	      "type": "string"
		//	    },
		//	    "StatusCode": {
		//	      "description": "",
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "description": "",
		//	      "enum": [
		//	        "API",
		//	        "AUTHORIZER",
		//	        "MODEL",
		//	        "RESOURCE",
		//	        "METHOD",
		//	        "PATH_PARAMETER",
		//	        "QUERY_PARAMETER",
		//	        "REQUEST_HEADER",
		//	        "REQUEST_BODY",
		//	        "RESPONSE",
		//	        "RESPONSE_HEADER",
		//	        "RESPONSE_BODY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Method
				"method": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Path
				"path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StatusCode
				"status_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The ``Location`` property specifies the location of the Amazon API Gateway API entity that the documentation applies to. ``Location`` is a property of the [AWS::ApiGateway::DocumentationPart](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html) resource.\n For more information about each property, including constraints and valid values, see [DocumentationPart](https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationPartLocation.html) in the *Amazon API Gateway REST API Reference*.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Properties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"properties": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::DocumentationPart",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationPart").WithTerraformTypeName("awscc_apigateway_documentation_part")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"documentation_part_id": "DocumentationPartId",
		"location":              "Location",
		"method":                "Method",
		"name":                  "Name",
		"path":                  "Path",
		"properties":            "Properties",
		"rest_api_id":           "RestApiId",
		"status_code":           "StatusCode",
		"type":                  "Type",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
