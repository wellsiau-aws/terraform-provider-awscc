// Code generated by generators/resource/main.go; DO NOT EDIT.

package databrew

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_databrew_job", jobResourceType)
}

// jobResourceType returns the Terraform awscc_databrew_job resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataBrew::Job resource type.
func jobResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"data_catalog_outputs": {
			// Property: DataCatalogOutputs
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "CatalogId": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "DatabaseName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "DatabaseOptions": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "TableName": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "TempDirectory": {
			//             "additionalProperties": false,
			//             "description": "S3 Output location",
			//             "properties": {
			//               "Bucket": {
			//                 "type": "string"
			//               },
			//               "Key": {
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Bucket"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "TableName"
			//         ],
			//         "type": "object"
			//       },
			//       "Overwrite": {
			//         "type": "boolean"
			//       },
			//       "S3Options": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Location": {
			//             "additionalProperties": false,
			//             "description": "S3 Output location",
			//             "properties": {
			//               "Bucket": {
			//                 "type": "string"
			//               },
			//               "Key": {
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Bucket"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Location"
			//         ],
			//         "type": "object"
			//       },
			//       "TableName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "DatabaseName",
			//       "TableName"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"catalog_id": {
						// Property: CatalogId
						Type:     types.StringType,
						Optional: true,
					},
					"database_name": {
						// Property: DatabaseName
						Type:     types.StringType,
						Required: true,
					},
					"database_options": {
						// Property: DatabaseOptions
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"table_name": {
									// Property: TableName
									Type:     types.StringType,
									Required: true,
								},
								"temp_directory": {
									// Property: TempDirectory
									Description: "S3 Output location",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Required: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"overwrite": {
						// Property: Overwrite
						Type:     types.BoolType,
						Optional: true,
					},
					"s3_options": {
						// Property: S3Options
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"location": {
									// Property: Location
									Description: "S3 Output location",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Required: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"table_name": {
						// Property: TableName
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"database_outputs": {
			// Property: DatabaseOutputs
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "DatabaseOptions": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "TableName": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "TempDirectory": {
			//             "additionalProperties": false,
			//             "description": "S3 Output location",
			//             "properties": {
			//               "Bucket": {
			//                 "type": "string"
			//               },
			//               "Key": {
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Bucket"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "TableName"
			//         ],
			//         "type": "object"
			//       },
			//       "DatabaseOutputMode": {
			//         "description": "Database table name",
			//         "enum": [
			//           "NEW_TABLE"
			//         ],
			//         "type": "string"
			//       },
			//       "GlueConnectionName": {
			//         "description": "Glue connection name",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "GlueConnectionName",
			//       "DatabaseOptions"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"database_options": {
						// Property: DatabaseOptions
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"table_name": {
									// Property: TableName
									Type:     types.StringType,
									Required: true,
								},
								"temp_directory": {
									// Property: TempDirectory
									Description: "S3 Output location",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"bucket": {
												// Property: Bucket
												Type:     types.StringType,
												Required: true,
											},
											"key": {
												// Property: Key
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Required: true,
					},
					"database_output_mode": {
						// Property: DatabaseOutputMode
						Description: "Database table name",
						Type:        types.StringType,
						Optional:    true,
					},
					"glue_connection_name": {
						// Property: GlueConnectionName
						Description: "Glue connection name",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"dataset_name": {
			// Property: DatasetName
			// CloudFormation resource type schema:
			// {
			//   "description": "Dataset name",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Dataset name",
			Type:        types.StringType,
			Optional:    true,
		},
		"encryption_key_arn": {
			// Property: EncryptionKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Encryption Key Arn",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "Encryption Key Arn",
			Type:        types.StringType,
			Optional:    true,
		},
		"encryption_mode": {
			// Property: EncryptionMode
			// CloudFormation resource type schema:
			// {
			//   "description": "Encryption mode",
			//   "enum": [
			//     "SSE-KMS",
			//     "SSE-S3"
			//   ],
			//   "type": "string"
			// }
			Description: "Encryption mode",
			Type:        types.StringType,
			Optional:    true,
		},
		"job_sample": {
			// Property: JobSample
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Job Sample",
			//   "properties": {
			//     "Mode": {
			//       "description": "Sample configuration mode for profile jobs.",
			//       "enum": [
			//         "FULL_DATASET",
			//         "CUSTOM_ROWS"
			//       ],
			//       "type": "string"
			//     },
			//     "Size": {
			//       "description": "Sample configuration size for profile jobs.",
			//       "format": "int64",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Job Sample",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"mode": {
						// Property: Mode
						Description: "Sample configuration mode for profile jobs.",
						Type:        types.StringType,
						Optional:    true,
					},
					"size": {
						// Property: Size
						Description: "Sample configuration size for profile jobs.",
						Type:        types.NumberType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"log_subscription": {
			// Property: LogSubscription
			// CloudFormation resource type schema:
			// {
			//   "description": "Log subscription",
			//   "enum": [
			//     "ENABLE",
			//     "DISABLE"
			//   ],
			//   "type": "string"
			// }
			Description: "Log subscription",
			Type:        types.StringType,
			Optional:    true,
		},
		"max_capacity": {
			// Property: MaxCapacity
			// CloudFormation resource type schema:
			// {
			//   "description": "Max capacity",
			//   "type": "integer"
			// }
			Description: "Max capacity",
			Type:        types.NumberType,
			Optional:    true,
		},
		"max_retries": {
			// Property: MaxRetries
			// CloudFormation resource type schema:
			// {
			//   "description": "Max retries",
			//   "type": "integer"
			// }
			Description: "Max retries",
			Type:        types.NumberType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Job name",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Job name",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"output_location": {
			// Property: OutputLocation
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Output location",
			//   "properties": {
			//     "Bucket": {
			//       "type": "string"
			//     },
			//     "Key": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Bucket"
			//   ],
			//   "type": "object"
			// }
			Description: "Output location",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"bucket": {
						// Property: Bucket
						Type:     types.StringType,
						Required: true,
					},
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"outputs": {
			// Property: Outputs
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "CompressionFormat": {
			//         "enum": [
			//           "GZIP",
			//           "LZ4",
			//           "SNAPPY",
			//           "BZIP2",
			//           "DEFLATE",
			//           "LZO",
			//           "BROTLI",
			//           "ZSTD",
			//           "ZLIB"
			//         ],
			//         "type": "string"
			//       },
			//       "Format": {
			//         "enum": [
			//           "CSV",
			//           "JSON",
			//           "PARQUET",
			//           "GLUEPARQUET",
			//           "AVRO",
			//           "ORC",
			//           "XML"
			//         ],
			//         "type": "string"
			//       },
			//       "FormatOptions": {
			//         "additionalProperties": false,
			//         "description": "Format options for job Output",
			//         "properties": {
			//           "Csv": {
			//             "additionalProperties": false,
			//             "description": "Output Csv options",
			//             "properties": {
			//               "Delimiter": {
			//                 "maxLength": 1,
			//                 "minLength": 1,
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Location": {
			//         "additionalProperties": false,
			//         "description": "S3 Output location",
			//         "properties": {
			//           "Bucket": {
			//             "type": "string"
			//           },
			//           "Key": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Bucket"
			//         ],
			//         "type": "object"
			//       },
			//       "Overwrite": {
			//         "type": "boolean"
			//       },
			//       "PartitionColumns": {
			//         "insertionOrder": true,
			//         "items": {
			//           "type": "string"
			//         },
			//         "type": "array",
			//         "uniqueItems": true
			//       }
			//     },
			//     "required": [
			//       "Location"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"compression_format": {
						// Property: CompressionFormat
						Type:     types.StringType,
						Optional: true,
					},
					"format": {
						// Property: Format
						Type:     types.StringType,
						Optional: true,
					},
					"format_options": {
						// Property: FormatOptions
						Description: "Format options for job Output",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"csv": {
									// Property: Csv
									Description: "Output Csv options",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"delimiter": {
												// Property: Delimiter
												Type:     types.StringType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"location": {
						// Property: Location
						Description: "S3 Output location",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"bucket": {
									// Property: Bucket
									Type:     types.StringType,
									Required: true,
								},
								"key": {
									// Property: Key
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Required: true,
					},
					"overwrite": {
						// Property: Overwrite
						Type:     types.BoolType,
						Optional: true,
					},
					"partition_columns": {
						// Property: PartitionColumns
						// Ordered set.
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"profile_configuration": {
			// Property: ProfileConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ColumnStatisticsConfigurations": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Selectors": {
			//             "insertionOrder": true,
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Name": {
			//                   "maxLength": 255,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Regex": {
			//                   "maxLength": 255,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "minItems": 1,
			//             "type": "array"
			//           },
			//           "Statistics": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "IncludedStatistics": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "minItems": 1,
			//                 "type": "array"
			//               },
			//               "Overrides": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "properties": {
			//                     "Parameters": {
			//                       "additionalProperties": false,
			//                       "patternProperties": {
			//                         "": {
			//                           "type": "string"
			//                         }
			//                       },
			//                       "type": "object"
			//                     },
			//                     "Statistic": {
			//                       "maxLength": 128,
			//                       "minLength": 1,
			//                       "pattern": "",
			//                       "type": "string"
			//                     }
			//                   },
			//                   "required": [
			//                     "Statistic",
			//                     "Parameters"
			//                   ],
			//                   "type": "object"
			//                 },
			//                 "minItems": 1,
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Statistics"
			//         ],
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "DatasetStatisticsConfiguration": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "IncludedStatistics": {
			//           "insertionOrder": true,
			//           "items": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         },
			//         "Overrides": {
			//           "insertionOrder": true,
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Parameters": {
			//                 "additionalProperties": false,
			//                 "patternProperties": {
			//                   "": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "Statistic": {
			//                 "maxLength": 128,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Statistic",
			//               "Parameters"
			//             ],
			//             "type": "object"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ProfileColumns": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Name": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Regex": {
			//             "maxLength": 255,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"column_statistics_configurations": {
						// Property: ColumnStatisticsConfigurations
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"selectors": {
									// Property: Selectors
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"name": {
												// Property: Name
												Type:     types.StringType,
												Optional: true,
											},
											"regex": {
												// Property: Regex
												Type:     types.StringType,
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{
											MinItems: 1,
										},
									),
									Optional: true,
								},
								"statistics": {
									// Property: Statistics
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"included_statistics": {
												// Property: IncludedStatistics
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
											},
											"overrides": {
												// Property: Overrides
												Attributes: schema.ListNestedAttributes(
													map[string]schema.Attribute{
														"parameters": {
															// Property: Parameters
															// Pattern: ""
															Type:     types.MapType{ElemType: types.StringType},
															Required: true,
														},
														"statistic": {
															// Property: Statistic
															Type:     types.StringType,
															Required: true,
														},
													},
													schema.ListNestedAttributesOptions{
														MinItems: 1,
													},
												),
												Optional: true,
											},
										},
									),
									Required: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Optional: true,
					},
					"dataset_statistics_configuration": {
						// Property: DatasetStatisticsConfiguration
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"included_statistics": {
									// Property: IncludedStatistics
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"overrides": {
									// Property: Overrides
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"parameters": {
												// Property: Parameters
												// Pattern: ""
												Type:     types.MapType{ElemType: types.StringType},
												Required: true,
											},
											"statistic": {
												// Property: Statistic
												Type:     types.StringType,
												Required: true,
											},
										},
										schema.ListNestedAttributesOptions{
											MinItems: 1,
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"profile_columns": {
						// Property: ProfileColumns
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
								"regex": {
									// Property: Regex
									Type:     types.StringType,
									Optional: true,
								},
							},
							schema.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			// {
			//   "description": "Project name",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Project name",
			Type:        types.StringType,
			Optional:    true,
		},
		"recipe": {
			// Property: Recipe
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Name": {
			//       "description": "Recipe name",
			//       "type": "string"
			//     },
			//     "Version": {
			//       "description": "Recipe version",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Name"
			//   ],
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"name": {
						// Property: Name
						Description: "Recipe name",
						Type:        types.StringType,
						Required:    true,
					},
					"version": {
						// Property: Version
						Description: "Recipe version",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role arn",
			//   "type": "string"
			// }
			Description: "Role arn",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
		"timeout": {
			// Property: Timeout
			// CloudFormation resource type schema:
			// {
			//   "description": "Timeout",
			//   "type": "integer"
			// }
			Description: "Timeout",
			Type:        types.NumberType,
			Optional:    true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "Job type",
			//   "enum": [
			//     "PROFILE",
			//     "RECIPE"
			//   ],
			//   "type": "string"
			// }
			Description: "Job type",
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
		Description: "Resource schema for AWS::DataBrew::Job.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataBrew::Job").WithTerraformTypeName("awscc_databrew_job").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_databrew_job", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
