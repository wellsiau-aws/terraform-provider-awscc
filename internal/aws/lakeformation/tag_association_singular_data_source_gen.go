// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lakeformation_tag_association", tagAssociationDataSourceType)
}

// tagAssociationDataSourceType returns the Terraform awscc_lakeformation_tag_association data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::LakeFormation::TagAssociation resource type.
func tagAssociationDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"lf_tags": {
			// Property: LFTags
			// CloudFormation resource type schema:
			// {
			//   "description": "List of Lake Formation Tags to associate with the Lake Formation Resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "CatalogId": {
			//         "maxLength": 12,
			//         "minLength": 12,
			//         "type": "string"
			//       },
			//       "TagKey": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "TagValues": {
			//         "insertionOrder": false,
			//         "items": {
			//           "maxLength": 256,
			//           "minLength": 0,
			//           "type": "string"
			//         },
			//         "maxItems": 50,
			//         "minItems": 1,
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "CatalogId",
			//       "TagKey",
			//       "TagValues"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "List of Lake Formation Tags to associate with the Lake Formation Resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"catalog_id": {
						// Property: CatalogId
						Type:     types.StringType,
						Computed: true,
					},
					"tag_key": {
						// Property: TagKey
						Type:     types.StringType,
						Computed: true,
					},
					"tag_values": {
						// Property: TagValues
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"resource": {
			// Property: Resource
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Resource to tag with the Lake Formation Tags",
			//   "properties": {
			//     "Catalog": {
			//       "additionalProperties": false,
			//       "type": "object"
			//     },
			//     "Database": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CatalogId": {
			//           "maxLength": 12,
			//           "minLength": 12,
			//           "type": "string"
			//         },
			//         "Name": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "CatalogId",
			//         "Name"
			//       ],
			//       "type": "object"
			//     },
			//     "Table": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CatalogId": {
			//           "maxLength": 12,
			//           "minLength": 12,
			//           "type": "string"
			//         },
			//         "DatabaseName": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Name": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "TableWildcard": {
			//           "additionalProperties": false,
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "CatalogId",
			//         "DatabaseName"
			//       ],
			//       "type": "object"
			//     },
			//     "TableWithColumns": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CatalogId": {
			//           "maxLength": 12,
			//           "minLength": 12,
			//           "type": "string"
			//         },
			//         "ColumnNames": {
			//           "insertionOrder": false,
			//           "items": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "DatabaseName": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Name": {
			//           "maxLength": 255,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "CatalogId",
			//         "DatabaseName",
			//         "Name",
			//         "ColumnNames"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Resource to tag with the Lake Formation Tags",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"catalog": {
						// Property: Catalog
						Type:     types.MapType{ElemType: types.StringType},
						Computed: true,
					},
					"database": {
						// Property: Database
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"table": {
						// Property: Table
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"database_name": {
									// Property: DatabaseName
									Type:     types.StringType,
									Computed: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
								"table_wildcard": {
									// Property: TableWildcard
									Type:     types.MapType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"table_with_columns": {
						// Property: TableWithColumns
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"column_names": {
									// Property: ColumnNames
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"database_name": {
									// Property: DatabaseName
									Type:     types.StringType,
									Computed: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"resource_identifier": {
			// Property: ResourceIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique string identifying the resource. Used as primary identifier, which ideally should be a string",
			//   "type": "string"
			// }
			Description: "Unique string identifying the resource. Used as primary identifier, which ideally should be a string",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags_identifier": {
			// Property: TagsIdentifier
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string",
			//   "type": "string"
			// }
			Description: "Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::LakeFormation::TagAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::TagAssociation").WithTerraformTypeName("awscc_lakeformation_tag_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":             "Catalog",
		"catalog_id":          "CatalogId",
		"column_names":        "ColumnNames",
		"database":            "Database",
		"database_name":       "DatabaseName",
		"lf_tags":             "LFTags",
		"name":                "Name",
		"resource":            "Resource",
		"resource_identifier": "ResourceIdentifier",
		"table":               "Table",
		"table_wildcard":      "TableWildcard",
		"table_with_columns":  "TableWithColumns",
		"tag_key":             "TagKey",
		"tag_values":          "TagValues",
		"tags_identifier":     "TagsIdentifier",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
