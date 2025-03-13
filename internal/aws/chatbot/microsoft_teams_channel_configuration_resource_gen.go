// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package chatbot

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_chatbot_microsoft_teams_channel_configuration", microsoftTeamsChannelConfigurationResource)
}

// microsoftTeamsChannelConfigurationResource returns the Terraform awscc_chatbot_microsoft_teams_channel_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::Chatbot::MicrosoftTeamsChannelConfiguration resource.
func microsoftTeamsChannelConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the configuration",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the configuration",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the configuration",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[A-Za-z0-9-_]+$",
		//	  "type": "string"
		//	}
		"configuration_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the configuration",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9-_]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomizationResourceArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARNs of Custom Actions to associate with notifications in the provided chat channel.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^arn:aws:chatbot:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:custom-action/[a-zA-Z0-9_-]{1,64}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"customization_resource_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "ARNs of Custom Actions to associate with notifications in the provided chat channel.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:chatbot:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:custom-action/[a-zA-Z0-9_-]{1,64}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GuardrailPolicies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"guardrail_policies": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IamRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the IAM role that defines the permissions for AWS Chatbot",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the IAM role that defines the permissions for AWS Chatbot",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoggingLevel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "NONE",
		//	  "description": "Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs",
		//	  "pattern": "^(ERROR|INFO|NONE)$",
		//	  "type": "string"
		//	}
		"logging_level": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("NONE"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^(ERROR|INFO|NONE)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"sns_topic_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to add to the configuration",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
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
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to add to the configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TeamId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the Microsoft Teams team",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[0-9A-Fa-f]{8}(?:-[0-9A-Fa-f]{4}){3}-[0-9A-Fa-f]{12}$",
		//	  "type": "string"
		//	}
		"team_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the Microsoft Teams team",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(36, 36),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9A-Fa-f]{8}(?:-[0-9A-Fa-f]{4}){3}-[0-9A-Fa-f]{12}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TeamsChannelId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the Microsoft Teams channel",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^([a-zA-Z0-9-_=+/.,])*%3[aA]([a-zA-Z0-9-_=+/.,])*%40([a-zA-Z0-9-_=+/.,])*$",
		//	  "type": "string"
		//	}
		"teams_channel_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the Microsoft Teams channel",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^([a-zA-Z0-9-_=+/.,])*%3[aA]([a-zA-Z0-9-_=+/.,])*%40([a-zA-Z0-9-_=+/.,])*$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: TeamsChannelName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Microsoft Teams channel",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^(.*)$",
		//	  "type": "string"
		//	}
		"teams_channel_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Microsoft Teams channel",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^(.*)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TeamsTenantId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the Microsoft Teams tenant",
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "^[0-9A-Fa-f]{8}(?:-[0-9A-Fa-f]{4}){3}-[0-9A-Fa-f]{12}$",
		//	  "type": "string"
		//	}
		"teams_tenant_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the Microsoft Teams tenant",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(36, 36),
				stringvalidator.RegexMatches(regexp.MustCompile("^[0-9A-Fa-f]{8}(?:-[0-9A-Fa-f]{4}){3}-[0-9A-Fa-f]{12}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserRoleRequired
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Enables use of a user role requirement in your chat configuration",
		//	  "type": "boolean"
		//	}
		"user_role_required": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Enables use of a user role requirement in your chat configuration",
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Chatbot::MicrosoftTeamsChannelConfiguration").WithTerraformTypeName("awscc_chatbot_microsoft_teams_channel_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"configuration_name":          "ConfigurationName",
		"customization_resource_arns": "CustomizationResourceArns",
		"guardrail_policies":          "GuardrailPolicies",
		"iam_role_arn":                "IamRoleArn",
		"key":                         "Key",
		"logging_level":               "LoggingLevel",
		"sns_topic_arns":              "SnsTopicArns",
		"tags":                        "Tags",
		"team_id":                     "TeamId",
		"teams_channel_id":            "TeamsChannelId",
		"teams_channel_name":          "TeamsChannelName",
		"teams_tenant_id":             "TeamsTenantId",
		"user_role_required":          "UserRoleRequired",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
