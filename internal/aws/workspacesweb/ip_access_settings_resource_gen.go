// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package workspacesweb

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_workspacesweb_ip_access_settings", ipAccessSettingsResource)
}

// ipAccessSettingsResource returns the Terraform awscc_workspacesweb_ip_access_settings resource.
// This Terraform resource corresponds to the CloudFormation AWS::WorkSpacesWeb::IpAccessSettings resource.
func ipAccessSettingsResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdditionalEncryptionContext
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 131072,
		//	      "minLength": 0,
		//	      "pattern": "^[\\s\\S]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"additional_encryption_context": // Pattern: ""
		schema.MapAttribute{             /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
				mapplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssociatedPortalArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 20,
		//	    "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"associated_portal_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: timetypes.RFC3339Type{},
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomerManagedKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:kms:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:key\\/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"customer_managed_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w+=\\/,.@-]+:kms:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:key\\/[a-zA-Z0-9-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpAccessSettingsArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[\\w+=\\/,.@-]+:[a-zA-Z0-9\\-]+:[a-zA-Z0-9\\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\\/[a-fA-F0-9\\-]{36})+$",
		//	  "type": "string"
		//	}
		"ip_access_settings_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Description": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^.+$",
		//	        "type": "string"
		//	      },
		//	      "IpRange": {
		//	        "description": "A single IP address or an IP address range in CIDR notation",
		//	        "pattern": "^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}(?:/([0-9]|[12][0-9]|3[0-2])|)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "IpRange"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 100,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"ip_rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^.+$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: IpRange
					"ip_range": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A single IP address or an IP address range in CIDR notation",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.RegexMatches(regexp.MustCompile("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}(?:/([0-9]|[12][0-9]|3[0-2])|)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 100),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
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
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Definition of AWS::WorkSpacesWeb::IpAccessSettings Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::WorkSpacesWeb::IpAccessSettings").WithTerraformTypeName("awscc_workspacesweb_ip_access_settings")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_encryption_context": "AdditionalEncryptionContext",
		"associated_portal_arns":        "AssociatedPortalArns",
		"creation_date":                 "CreationDate",
		"customer_managed_key":          "CustomerManagedKey",
		"description":                   "Description",
		"display_name":                  "DisplayName",
		"ip_access_settings_arn":        "IpAccessSettingsArn",
		"ip_range":                      "IpRange",
		"ip_rules":                      "IpRules",
		"key":                           "Key",
		"tags":                          "Tags",
		"value":                         "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
