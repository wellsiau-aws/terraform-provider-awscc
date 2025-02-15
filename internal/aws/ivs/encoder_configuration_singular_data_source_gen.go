// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ivs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ivs_encoder_configuration", encoderConfigurationDataSource)
}

// encoderConfigurationDataSource returns the Terraform awscc_ivs_encoder_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::IVS::EncoderConfiguration resource.
func encoderConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Encoder configuration identifier.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:encoder-configuration/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Encoder configuration identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Encoder configuration name.",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z0-9-_]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Encoder configuration name.",
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
		//	      "Key",
		//	      "Value"
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
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Video
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps",
		//	  "properties": {
		//	    "Bitrate": {
		//	      "default": 2500000,
		//	      "description": "Bitrate for generated output, in bps. Default: 2500000.",
		//	      "maximum": 8500000,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "Framerate": {
		//	      "default": 30,
		//	      "description": "Video frame rate, in fps. Default: 30.",
		//	      "maximum": 60,
		//	      "minimum": 1,
		//	      "type": "number"
		//	    },
		//	    "Height": {
		//	      "default": 720,
		//	      "description": "Video-resolution height. This must be an even number. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 720.",
		//	      "maximum": 1920,
		//	      "minimum": 2,
		//	      "type": "integer"
		//	    },
		//	    "Width": {
		//	      "default": 1280,
		//	      "description": "Video-resolution width. This must be an even number. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 1280.",
		//	      "maximum": 1920,
		//	      "minimum": 2,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"video": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Bitrate
				"bitrate": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Bitrate for generated output, in bps. Default: 2500000.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Framerate
				"framerate": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "Video frame rate, in fps. Default: 30.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Height
				"height": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Video-resolution height. This must be an even number. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 720.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Width
				"width": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Video-resolution width. This must be an even number. Note that the maximum value is determined by width times height, such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 1280.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IVS::EncoderConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::EncoderConfiguration").WithTerraformTypeName("awscc_ivs_encoder_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":       "Arn",
		"bitrate":   "Bitrate",
		"framerate": "Framerate",
		"height":    "Height",
		"key":       "Key",
		"name":      "Name",
		"tags":      "Tags",
		"value":     "Value",
		"video":     "Video",
		"width":     "Width",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
