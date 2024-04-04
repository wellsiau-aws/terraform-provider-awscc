// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_groundstation_config", configDataSource)
}

// configDataSource returns the Terraform awscc_groundstation_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::GroundStation::Config resource.
func configDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AntennaDownlinkConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SpectrumConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Bandwidth": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Units": {
		//	                  "enum": [
		//	                    "GHz",
		//	                    "MHz",
		//	                    "kHz"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "type": "number"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "CenterFrequency": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Units": {
		//	                  "enum": [
		//	                    "GHz",
		//	                    "MHz",
		//	                    "kHz"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "type": "number"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "Polarization": {
		//	              "enum": [
		//	                "LEFT_HAND",
		//	                "RIGHT_HAND",
		//	                "NONE"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "AntennaDownlinkDemodDecodeConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DecodeConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "UnvalidatedJSON": {
		//	              "pattern": "",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "DemodulationConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "UnvalidatedJSON": {
		//	              "pattern": "",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "SpectrumConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Bandwidth": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Units": {
		//	                  "enum": [
		//	                    "GHz",
		//	                    "MHz",
		//	                    "kHz"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "type": "number"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "CenterFrequency": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Units": {
		//	                  "enum": [
		//	                    "GHz",
		//	                    "MHz",
		//	                    "kHz"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "type": "number"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "Polarization": {
		//	              "enum": [
		//	                "LEFT_HAND",
		//	                "RIGHT_HAND",
		//	                "NONE"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "AntennaUplinkConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SpectrumConfig": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "CenterFrequency": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Units": {
		//	                  "enum": [
		//	                    "GHz",
		//	                    "MHz",
		//	                    "kHz"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "type": "number"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "Polarization": {
		//	              "enum": [
		//	                "LEFT_HAND",
		//	                "RIGHT_HAND",
		//	                "NONE"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "TargetEirp": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Units": {
		//	              "enum": [
		//	                "dBW"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "type": "number"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "TransmitDisabled": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "DataflowEndpointConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DataflowEndpointName": {
		//	          "type": "string"
		//	        },
		//	        "DataflowEndpointRegion": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "S3RecordingConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "BucketArn": {
		//	          "type": "string"
		//	        },
		//	        "Prefix": {
		//	          "pattern": "^([a-zA-Z0-9_\\-=/]|\\{satellite_id\\}|\\{config\\-name}|\\{s3\\-config-id}|\\{year\\}|\\{month\\}|\\{day\\}){1,900}$",
		//	          "type": "string"
		//	        },
		//	        "RoleArn": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "TrackingConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Autotrack": {
		//	          "enum": [
		//	            "REQUIRED",
		//	            "PREFERRED",
		//	            "REMOVED"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UplinkEchoConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AntennaUplinkConfigArn": {
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"config_data": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AntennaDownlinkConfig
				"antenna_downlink_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SpectrumConfig
						"spectrum_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Bandwidth
								"bandwidth": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Units
										"units": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: CenterFrequency
								"center_frequency": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Units
										"units": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Polarization
								"polarization": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AntennaDownlinkDemodDecodeConfig
				"antenna_downlink_demod_decode_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DecodeConfig
						"decode_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: UnvalidatedJSON
								"unvalidated_json": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DemodulationConfig
						"demodulation_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: UnvalidatedJSON
								"unvalidated_json": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: SpectrumConfig
						"spectrum_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Bandwidth
								"bandwidth": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Units
										"units": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: CenterFrequency
								"center_frequency": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Units
										"units": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Polarization
								"polarization": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AntennaUplinkConfig
				"antenna_uplink_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SpectrumConfig
						"spectrum_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CenterFrequency
								"center_frequency": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Units
										"units": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Polarization
								"polarization": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TargetEirp
						"target_eirp": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Units
								"units": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TransmitDisabled
						"transmit_disabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DataflowEndpointConfig
				"dataflow_endpoint_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DataflowEndpointName
						"dataflow_endpoint_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: DataflowEndpointRegion
						"dataflow_endpoint_region": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: S3RecordingConfig
				"s3_recording_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketArn
						"bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Prefix
						"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: RoleArn
						"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TrackingConfig
				"tracking_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Autotrack
						"autotrack": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: UplinkEchoConfig
				"uplink_echo_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AntennaUplinkConfigArn
						"antenna_uplink_config_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"config_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
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
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::GroundStation::Config",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::Config").WithTerraformTypeName("awscc_groundstation_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"antenna_downlink_config":              "AntennaDownlinkConfig",
		"antenna_downlink_demod_decode_config": "AntennaDownlinkDemodDecodeConfig",
		"antenna_uplink_config":                "AntennaUplinkConfig",
		"antenna_uplink_config_arn":            "AntennaUplinkConfigArn",
		"arn":                                  "Arn",
		"autotrack":                            "Autotrack",
		"bandwidth":                            "Bandwidth",
		"bucket_arn":                           "BucketArn",
		"center_frequency":                     "CenterFrequency",
		"config_data":                          "ConfigData",
		"config_id":                            "Id",
		"dataflow_endpoint_config":             "DataflowEndpointConfig",
		"dataflow_endpoint_name":               "DataflowEndpointName",
		"dataflow_endpoint_region":             "DataflowEndpointRegion",
		"decode_config":                        "DecodeConfig",
		"demodulation_config":                  "DemodulationConfig",
		"enabled":                              "Enabled",
		"key":                                  "Key",
		"name":                                 "Name",
		"polarization":                         "Polarization",
		"prefix":                               "Prefix",
		"role_arn":                             "RoleArn",
		"s3_recording_config":                  "S3RecordingConfig",
		"spectrum_config":                      "SpectrumConfig",
		"tags":                                 "Tags",
		"target_eirp":                          "TargetEirp",
		"tracking_config":                      "TrackingConfig",
		"transmit_disabled":                    "TransmitDisabled",
		"type":                                 "Type",
		"units":                                "Units",
		"unvalidated_json":                     "UnvalidatedJSON",
		"uplink_echo_config":                   "UplinkEchoConfig",
		"value":                                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}