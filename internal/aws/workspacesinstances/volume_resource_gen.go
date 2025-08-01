// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package workspacesinstances

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_workspacesinstances_volume", volumeResource)
}

// volumeResource returns the Terraform awscc_workspacesinstances_volume resource.
// This Terraform resource corresponds to the CloudFormation AWS::WorkspacesInstances::Volume resource.
func volumeResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone in which to create the volume",
		//	  "pattern": "^[a-z]{2}-[a-z]+-\\d[a-z](-[a-z0-9]+)?$",
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone in which to create the volume",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]{2}-[a-z]+-\\d[a-z](-[a-z0-9]+)?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Encrypted
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the volume should be encrypted",
		//	  "type": "boolean"
		//	}
		"encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the volume should be encrypted",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Iops
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of I/O operations per second (IOPS)",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"iops": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of I/O operations per second (IOPS)",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.AtLeast(0),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for Amazon EBS encryption",
		//	  "maxLength": 128,
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for Amazon EBS encryption",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SizeInGB
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The size of the volume, in GiBs",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"size_in_gb": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The size of the volume, in GiBs",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.AtLeast(0),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnapshotId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The snapshot from which to create the volume",
		//	  "pattern": "^snap-[0-9a-zA-Z]{1,63}$",
		//	  "type": "string"
		//	}
		"snapshot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The snapshot from which to create the volume",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^snap-[0-9a-zA-Z]{1,63}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TagSpecifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags passed to EBS volume",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ResourceType": {
		//	        "enum": [
		//	          "instance",
		//	          "volume",
		//	          "spot-instances-request",
		//	          "network-interface"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "The tags to apply to the resource",
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "description": "The key name of the tag",
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "description": "The value for the tag",
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 30,
		//	  "type": "array"
		//	}
		"tag_specifications": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ResourceType
					"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"instance",
								"volume",
								"spot-instances-request",
								"network-interface",
							),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Tags
					"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The key name of the tag",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 128),
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The value for the tag",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The tags to apply to the resource",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags passed to EBS volume",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(30),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Throughput
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The throughput to provision for a volume, with a maximum of 1,000 MiB/s",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"throughput": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The throughput to provision for a volume, with a maximum of 1,000 MiB/s",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.AtLeast(0),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VolumeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier for the volume",
		//	  "pattern": "^vol-[0-9a-zA-Z]{1,63}$",
		//	  "type": "string"
		//	}
		"volume_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier for the volume",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VolumeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The volume type",
		//	  "enum": [
		//	    "standard",
		//	    "io1",
		//	    "io2",
		//	    "gp2",
		//	    "sc1",
		//	    "st1",
		//	    "gp3"
		//	  ],
		//	  "type": "string"
		//	}
		"volume_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The volume type",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"standard",
					"io1",
					"io2",
					"gp2",
					"sc1",
					"st1",
					"gp3",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::WorkspacesInstances::Volume - Manages WorkSpaces Volume resources",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::WorkspacesInstances::Volume").WithTerraformTypeName("awscc_workspacesinstances_volume")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"availability_zone":  "AvailabilityZone",
		"encrypted":          "Encrypted",
		"iops":               "Iops",
		"key":                "Key",
		"kms_key_id":         "KmsKeyId",
		"resource_type":      "ResourceType",
		"size_in_gb":         "SizeInGB",
		"snapshot_id":        "SnapshotId",
		"tag_specifications": "TagSpecifications",
		"tags":               "Tags",
		"throughput":         "Throughput",
		"value":              "Value",
		"volume_id":          "VolumeId",
		"volume_type":        "VolumeType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
