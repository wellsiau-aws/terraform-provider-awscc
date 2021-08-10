// Code generated by generators/resource/main.go; DO NOT EDIT.

package accessanalyzer

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_accessanalyzer_analyzer", analyzer)
}

// analyzer returns the Terraform aws_accessanalyzer_analyzer resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AccessAnalyzer::Analyzer resource type.
func analyzer(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"analyzer_name": {
			// Property: AnalyzerName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Analyzer name",
			     "maxLength": 1024,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Analyzer name",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// AnalyzerName is a force-new attribute.
		},
		"archive_rules": {
			// Property: ArchiveRules
			// CloudFormation resource type schema:
			/*
			   {
			     "insertionOrder": false,
			     "items": {
			       "description": "An Access Analyzer archive rule. Archive rules automatically archive new findings that meet the criteria you define when you create the rule.",
			       "properties": {
			         "Filter": {
			           "insertionOrder": false,
			           "items": {
			             "properties": {
			               "Contains": {
			                 "insertionOrder": false,
			                 "items": {
			                   "type": "string"
			                 },
			                 "type": "array"
			               },
			               "Eq": {
			                 "insertionOrder": false,
			                 "items": {
			                   "type": "string"
			                 },
			                 "type": "array"
			               },
			               "Exists": {
			                 "type": "boolean"
			               },
			               "Neq": {
			                 "insertionOrder": false,
			                 "items": {
			                   "type": "string"
			                 },
			                 "type": "array"
			               },
			               "Property": {
			                 "type": "string"
			               }
			             },
			             "$ref": "#/definitions/Filter",
			             "required": [
			               "Property"
			             ],
			             "type": "object"
			           },
			           "minItems": 1,
			           "type": "array"
			         },
			         "RuleName": {
			           "description": "The archive rule name",
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/ArchiveRule",
			       "required": [
			         "Filter",
			         "RuleName"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"filter": {
						// Property: Filter
						// CloudFormation resource type schema:
						/*
						   {
						     "insertionOrder": false,
						     "items": {
						       "properties": {
						         "Contains": {
						           "insertionOrder": false,
						           "items": {
						             "type": "string"
						           },
						           "type": "array"
						         },
						         "Eq": {
						           "insertionOrder": false,
						           "items": {
						             "type": "string"
						           },
						           "type": "array"
						         },
						         "Exists": {
						           "type": "boolean"
						         },
						         "Neq": {
						           "insertionOrder": false,
						           "items": {
						             "type": "string"
						           },
						           "type": "array"
						         },
						         "Property": {
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/Filter",
						       "required": [
						         "Property"
						       ],
						       "type": "object"
						     },
						     "minItems": 1,
						     "type": "array"
						   }
						*/
						// Multiset.
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"contains": {
									// Property: Contains
									// CloudFormation resource type schema:
									/*
									   {
									     "insertionOrder": false,
									     "items": {
									       "type": "string"
									     },
									     "type": "array"
									   }
									*/
									// Multiset.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"eq": {
									// Property: Eq
									// CloudFormation resource type schema:
									/*
									   {
									     "insertionOrder": false,
									     "items": {
									       "type": "string"
									     },
									     "type": "array"
									   }
									*/
									// Multiset.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"exists": {
									// Property: Exists
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "boolean"
									   }
									*/
									Type:     types.BoolType,
									Optional: true,
								},
								"neq": {
									// Property: Neq
									// CloudFormation resource type schema:
									/*
									   {
									     "insertionOrder": false,
									     "items": {
									       "type": "string"
									     },
									     "type": "array"
									   }
									*/
									// Multiset.
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"property": {
									// Property: Property
									// CloudFormation resource type schema:
									/*
									   {
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Required: true,
					},
					"rule_name": {
						// Property: RuleName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The archive rule name",
						     "type": "string"
						   }
						*/
						Description: "The archive rule name",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Amazon Resource Name (ARN) of the analyzer",
			     "maxLength": 1600,
			     "minLength": 1,
			     "type": "string"
			   }
			*/
			Description: "Amazon Resource Name (ARN) of the analyzer",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An array of key-value pairs to apply to this resource.",
			     "insertionOrder": false,
			     "items": {
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 127,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 255,
			           "minLength": 1,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": true
			   }
			*/
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 127,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 255,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The type of the analyzer, must be ACCOUNT or ORGANIZATION",
			     "maxLength": 1024,
			     "minLength": 0,
			     "type": "string"
			   }
			*/
			Description: "The type of the analyzer, must be ACCOUNT or ORGANIZATION",
			Type:        types.StringType,
			Required:    true,
			// Type is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AccessAnalyzer::Analyzer").WithTerraformTypeName("aws_accessanalyzer_analyzer").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_accessanalyzer_analyzer", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
