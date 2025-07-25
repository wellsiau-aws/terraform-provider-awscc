// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package b2bi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_b2bi_partnership", partnershipDataSource)
}

// partnershipDataSource returns the Terraform awscc_b2bi_partnership data source.
// This Terraform data source corresponds to the CloudFormation AWS::B2BI::Partnership resource.
func partnershipDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Capabilities
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "maxLength": 64,
		//	    "minLength": 1,
		//	    "pattern": "^[a-zA-Z0-9_-]+$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"capabilities": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CapabilityOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "InboundEdi": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "X12": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "AcknowledgmentOptions": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "FunctionalAcknowledgment": {
		//	                  "enum": [
		//	                    "DO_NOT_GENERATE",
		//	                    "GENERATE_ALL_SEGMENTS",
		//	                    "GENERATE_WITHOUT_TRANSACTION_SET_RESPONSE_LOOP"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "TechnicalAcknowledgment": {
		//	                  "enum": [
		//	                    "DO_NOT_GENERATE",
		//	                    "GENERATE_ALL_SEGMENTS"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "FunctionalAcknowledgment",
		//	                "TechnicalAcknowledgment"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "OutboundEdi": {
		//	      "properties": {
		//	        "X12": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Common": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "ControlNumbers": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "StartingFunctionalGroupControlNumber": {
		//	                      "maximum": 999999999,
		//	                      "minimum": 1,
		//	                      "type": "number"
		//	                    },
		//	                    "StartingInterchangeControlNumber": {
		//	                      "maximum": 999999999,
		//	                      "minimum": 1,
		//	                      "type": "number"
		//	                    },
		//	                    "StartingTransactionSetControlNumber": {
		//	                      "maximum": 999999999,
		//	                      "minimum": 1,
		//	                      "type": "number"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "Delimiters": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "ComponentSeparator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[!\u0026'()*+,\\-./:;?=%@\\[\\]_{}|\u003c\u003e~^`\"]$",
		//	                      "type": "string"
		//	                    },
		//	                    "DataElementSeparator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[!\u0026'()*+,\\-./:;?=%@\\[\\]_{}|\u003c\u003e~^`\"]$",
		//	                      "type": "string"
		//	                    },
		//	                    "SegmentTerminator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[!\u0026'()*+,\\-./:;?=%@\\[\\]_{}|\u003c\u003e~^`\"]$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "FunctionalGroupHeaders": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "ApplicationReceiverCode": {
		//	                      "maxLength": 15,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9 ]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ApplicationSenderCode": {
		//	                      "maxLength": 15,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9 ]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ResponsibleAgencyCode": {
		//	                      "maxLength": 2,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "Gs05TimeFormat": {
		//	                  "enum": [
		//	                    "HHMM",
		//	                    "HHMMSS",
		//	                    "HHMMSSDD"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "InterchangeControlHeaders": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "AcknowledgmentRequestedCode": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ReceiverId": {
		//	                      "maxLength": 15,
		//	                      "minLength": 15,
		//	                      "pattern": "^[a-zA-Z0-9 ]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "ReceiverIdQualifier": {
		//	                      "maxLength": 2,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "RepetitionSeparator": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "type": "string"
		//	                    },
		//	                    "SenderId": {
		//	                      "maxLength": 15,
		//	                      "minLength": 15,
		//	                      "pattern": "^[a-zA-Z0-9 ]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "SenderIdQualifier": {
		//	                      "maxLength": 2,
		//	                      "minLength": 2,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    },
		//	                    "UsageIndicatorCode": {
		//	                      "maxLength": 1,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-zA-Z0-9]*$",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "ValidateEdi": {
		//	                  "type": "boolean"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "WrapOptions": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "LineLength": {
		//	                  "minimum": 1,
		//	                  "type": "number"
		//	                },
		//	                "LineTerminator": {
		//	                  "enum": [
		//	                    "CRLF",
		//	                    "LF",
		//	                    "CR"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "WrapBy": {
		//	                  "enum": [
		//	                    "SEGMENT",
		//	                    "ONE_LINE",
		//	                    "LINE_LENGTH"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"capability_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InboundEdi
				"inbound_edi": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: X12
						"x12": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AcknowledgmentOptions
								"acknowledgment_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: FunctionalAcknowledgment
										"functional_acknowledgment": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: TechnicalAcknowledgment
										"technical_acknowledgment": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OutboundEdi
				"outbound_edi": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: X12
						"x12": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Common
								"common": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ControlNumbers
										"control_numbers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: StartingFunctionalGroupControlNumber
												"starting_functional_group_control_number": schema.Float64Attribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: StartingInterchangeControlNumber
												"starting_interchange_control_number": schema.Float64Attribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: StartingTransactionSetControlNumber
												"starting_transaction_set_control_number": schema.Float64Attribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Delimiters
										"delimiters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: ComponentSeparator
												"component_separator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: DataElementSeparator
												"data_element_separator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: SegmentTerminator
												"segment_terminator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: FunctionalGroupHeaders
										"functional_group_headers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: ApplicationReceiverCode
												"application_receiver_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: ApplicationSenderCode
												"application_sender_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: ResponsibleAgencyCode
												"responsible_agency_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Gs05TimeFormat
										"gs_05_time_format": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: InterchangeControlHeaders
										"interchange_control_headers": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: AcknowledgmentRequestedCode
												"acknowledgment_requested_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: ReceiverId
												"receiver_id": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: ReceiverIdQualifier
												"receiver_id_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: RepetitionSeparator
												"repetition_separator": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: SenderId
												"sender_id": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: SenderIdQualifier
												"sender_id_qualifier": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
												// Property: UsageIndicatorCode
												"usage_indicator_code": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: ValidateEdi
										"validate_edi": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: WrapOptions
								"wrap_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: LineLength
										"line_length": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: LineTerminator
										"line_terminator": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: WrapBy
										"wrap_by": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: Email
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 254,
		//	  "minLength": 5,
		//	  "pattern": "^[\\w\\.\\-]+@[\\w\\.\\-]+$",
		//	  "type": "string"
		//	}
		"email": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 254,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PartnershipArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"partnership_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PartnershipId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"partnership_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Phone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 22,
		//	  "minLength": 7,
		//	  "pattern": "^\\+?([0-9 \\t\\-()\\/]{7,})(?:\\s*(?:#|x\\.?|ext\\.?|extension) \\t*(\\d+))?$",
		//	  "type": "string"
		//	}
		"phone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProfileId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
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
		// Property: TradingPartnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9_-]+$",
		//	  "type": "string"
		//	}
		"trading_partner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::B2BI::Partnership",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::B2BI::Partnership").WithTerraformTypeName("awscc_b2bi_partnership")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"acknowledgment_options":        "AcknowledgmentOptions",
		"acknowledgment_requested_code": "AcknowledgmentRequestedCode",
		"application_receiver_code":     "ApplicationReceiverCode",
		"application_sender_code":       "ApplicationSenderCode",
		"capabilities":                  "Capabilities",
		"capability_options":            "CapabilityOptions",
		"common":                        "Common",
		"component_separator":           "ComponentSeparator",
		"control_numbers":               "ControlNumbers",
		"created_at":                    "CreatedAt",
		"data_element_separator":        "DataElementSeparator",
		"delimiters":                    "Delimiters",
		"email":                         "Email",
		"functional_acknowledgment":     "FunctionalAcknowledgment",
		"functional_group_headers":      "FunctionalGroupHeaders",
		"gs_05_time_format":             "Gs05TimeFormat",
		"inbound_edi":                   "InboundEdi",
		"interchange_control_headers":   "InterchangeControlHeaders",
		"key":                           "Key",
		"line_length":                   "LineLength",
		"line_terminator":               "LineTerminator",
		"modified_at":                   "ModifiedAt",
		"name":                          "Name",
		"outbound_edi":                  "OutboundEdi",
		"partnership_arn":               "PartnershipArn",
		"partnership_id":                "PartnershipId",
		"phone":                         "Phone",
		"profile_id":                    "ProfileId",
		"receiver_id":                   "ReceiverId",
		"receiver_id_qualifier":         "ReceiverIdQualifier",
		"repetition_separator":          "RepetitionSeparator",
		"responsible_agency_code":       "ResponsibleAgencyCode",
		"segment_terminator":            "SegmentTerminator",
		"sender_id":                     "SenderId",
		"sender_id_qualifier":           "SenderIdQualifier",
		"starting_functional_group_control_number": "StartingFunctionalGroupControlNumber",
		"starting_interchange_control_number":      "StartingInterchangeControlNumber",
		"starting_transaction_set_control_number":  "StartingTransactionSetControlNumber",
		"tags":                     "Tags",
		"technical_acknowledgment": "TechnicalAcknowledgment",
		"trading_partner_id":       "TradingPartnerId",
		"usage_indicator_code":     "UsageIndicatorCode",
		"validate_edi":             "ValidateEdi",
		"value":                    "Value",
		"wrap_by":                  "WrapBy",
		"wrap_options":             "WrapOptions",
		"x12":                      "X12",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
