// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cognito

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_cognito_managed_login_branding", managedLoginBrandingResource)
}

// managedLoginBrandingResource returns the Terraform awscc_cognito_managed_login_branding resource.
// This Terraform resource corresponds to the CloudFormation AWS::Cognito::ManagedLoginBranding resource.
func managedLoginBrandingResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Assets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Bytes": {
		//	        "maxLength": 1000000,
		//	        "type": "string"
		//	      },
		//	      "Category": {
		//	        "enum": [
		//	          "FAVICON_ICO",
		//	          "FAVICON_SVG",
		//	          "EMAIL_GRAPHIC",
		//	          "SMS_GRAPHIC",
		//	          "AUTH_APP_GRAPHIC",
		//	          "PASSWORD_GRAPHIC",
		//	          "PASSKEY_GRAPHIC",
		//	          "PAGE_HEADER_LOGO",
		//	          "PAGE_HEADER_BACKGROUND",
		//	          "PAGE_FOOTER_LOGO",
		//	          "PAGE_FOOTER_BACKGROUND",
		//	          "PAGE_BACKGROUND",
		//	          "FORM_BACKGROUND",
		//	          "FORM_LOGO",
		//	          "IDP_BUTTON_ICON"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "ColorMode": {
		//	        "enum": [
		//	          "LIGHT",
		//	          "DARK",
		//	          "DYNAMIC"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Extension": {
		//	        "enum": [
		//	          "ICO",
		//	          "JPEG",
		//	          "PNG",
		//	          "SVG",
		//	          "WEBP"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "ResourceId": {
		//	        "maxLength": 40,
		//	        "minLength": 1,
		//	        "pattern": "^[\\w\\- ]+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Category",
		//	      "ColorMode",
		//	      "Extension"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"assets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Bytes
					"bytes": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(1000000),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Category
					"category": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"FAVICON_ICO",
								"FAVICON_SVG",
								"EMAIL_GRAPHIC",
								"SMS_GRAPHIC",
								"AUTH_APP_GRAPHIC",
								"PASSWORD_GRAPHIC",
								"PASSKEY_GRAPHIC",
								"PAGE_HEADER_LOGO",
								"PAGE_HEADER_BACKGROUND",
								"PAGE_FOOTER_LOGO",
								"PAGE_FOOTER_BACKGROUND",
								"PAGE_BACKGROUND",
								"FORM_BACKGROUND",
								"FORM_LOGO",
								"IDP_BUTTON_ICON",
							),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ColorMode
					"color_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"LIGHT",
								"DARK",
								"DYNAMIC",
							),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Extension
					"extension": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"ICO",
								"JPEG",
								"PNG",
								"SVG",
								"WEBP",
							),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ResourceId
					"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 40),
							stringvalidator.RegexMatches(regexp.MustCompile("^[\\w\\- ]+$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClientId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// ClientId is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ManagedLoginBrandingId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[4][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$",
		//	  "type": "string"
		//	}
		"managed_login_branding_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReturnMergedResources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"return_merged_resources": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// ReturnMergedResources is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Settings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "object"
		//	}
		"settings": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType: jsontypes.NormalizedType{},
			Optional:   true,
			Computed:   true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UseCognitoProvidedValues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"use_cognito_provided_values": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"user_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::Cognito::ManagedLoginBranding",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Cognito::ManagedLoginBranding").WithTerraformTypeName("awscc_cognito_managed_login_branding")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assets":                      "Assets",
		"bytes":                       "Bytes",
		"category":                    "Category",
		"client_id":                   "ClientId",
		"color_mode":                  "ColorMode",
		"extension":                   "Extension",
		"managed_login_branding_id":   "ManagedLoginBrandingId",
		"resource_id":                 "ResourceId",
		"return_merged_resources":     "ReturnMergedResources",
		"settings":                    "Settings",
		"use_cognito_provided_values": "UseCognitoProvidedValues",
		"user_pool_id":                "UserPoolId",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ClientId",
		"/properties/ReturnMergedResources",
	})
	opts = opts.WithCreateTimeoutInMinutes(2).WithDeleteTimeoutInMinutes(2)

	opts = opts.WithUpdateTimeoutInMinutes(2)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
