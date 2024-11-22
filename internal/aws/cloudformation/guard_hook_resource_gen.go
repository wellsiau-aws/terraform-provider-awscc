// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_cloudformation_guard_hook", guardHookResource)
}

// guardHookResource returns the Terraform awscc_cloudformation_guard_hook resource.
// This Terraform resource corresponds to the CloudFormation AWS::CloudFormation::GuardHook resource.
func guardHookResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The typename alias for the hook.",
		//	  "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$",
		//	  "type": "string"
		//	}
		"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The typename alias for the hook.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExecutionRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The execution role ARN assumed by hooks to read Guard rules from S3 and write Guard outputs to S3.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:.+:iam::[0-9]{12}:role/.+",
		//	  "type": "string"
		//	}
		"execution_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The execution role ARN assumed by hooks to read Guard rules from S3 and write Guard outputs to S3.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:.+:iam::[0-9]{12}:role/.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FailureMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "WARN",
		//	  "description": "Attribute to specify CloudFormation behavior on hook failure.",
		//	  "enum": [
		//	    "FAIL",
		//	    "WARN"
		//	  ],
		//	  "type": "string"
		//	}
		"failure_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Attribute to specify CloudFormation behavior on hook failure.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("WARN"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"FAIL",
					"WARN",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HookArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the activated hook",
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$",
		//	  "type": "string"
		//	}
		"hook_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the activated hook",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HookStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "DISABLED",
		//	  "description": "Attribute to specify which stacks this hook applies to or should get invoked for",
		//	  "enum": [
		//	    "ENABLED",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"hook_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Attribute to specify which stacks this hook applies to or should get invoked for",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("DISABLED"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ENABLED",
					"DISABLED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogBucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "S3 Bucket where the guard validate report will be uploaded to",
		//	  "type": "string"
		//	}
		"log_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "S3 Bucket where the guard validate report will be uploaded to",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Options
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "InputParams": {
		//	      "additionalProperties": false,
		//	      "description": "S3 Source Location for the Guard files.",
		//	      "properties": {
		//	        "Uri": {
		//	          "description": "S3 uri of Guard files.",
		//	          "type": "string"
		//	        },
		//	        "VersionId": {
		//	          "description": "S3 object version",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Uri"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  }
		//	}
		"options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InputParams
				"input_params": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Uri
						"uri": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "S3 uri of Guard files.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: VersionId
						"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "S3 object version",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "S3 Source Location for the Guard files.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RuleLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "S3 Source Location for the Guard files.",
		//	  "properties": {
		//	    "Uri": {
		//	      "description": "S3 uri of Guard files.",
		//	      "type": "string"
		//	    },
		//	    "VersionId": {
		//	      "description": "S3 object version",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Uri"
		//	  ],
		//	  "type": "object"
		//	}
		"rule_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Uri
				"uri": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "S3 uri of Guard files.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: VersionId
				"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "S3 object version",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "S3 Source Location for the Guard files.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: StackFilters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Filters to allow hooks to target specific stack attributes",
		//	  "properties": {
		//	    "FilteringCriteria": {
		//	      "default": "ALL",
		//	      "description": "Attribute to specify the filtering behavior. ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match",
		//	      "enum": [
		//	        "ALL",
		//	        "ANY"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "StackNames": {
		//	      "additionalProperties": false,
		//	      "description": "List of stack names as filters",
		//	      "properties": {
		//	        "Exclude": {
		//	          "description": "List of stack names that the hook is going to be excluded from",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "CloudFormation Stack name",
		//	            "maxLength": 128,
		//	            "pattern": "^[a-zA-Z][-a-zA-Z0-9]*$",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Include": {
		//	          "description": "List of stack names that the hook is going to target",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "CloudFormation Stack name",
		//	            "maxLength": 128,
		//	            "pattern": "^[a-zA-Z][-a-zA-Z0-9]*$",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "StackRoles": {
		//	      "additionalProperties": false,
		//	      "description": "List of stack roles that are performing the stack operations.",
		//	      "properties": {
		//	        "Exclude": {
		//	          "description": "List of stack roles that the hook is going to be excluded from",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "IAM Role ARN",
		//	            "maxLength": 256,
		//	            "pattern": "arn:.+:iam::[0-9]{12}:role/.+",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "Include": {
		//	          "description": "List of stack roles that the hook is going to target",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "description": "IAM Role ARN",
		//	            "maxLength": 256,
		//	            "pattern": "arn:.+:iam::[0-9]{12}:role/.+",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "FilteringCriteria"
		//	  ],
		//	  "type": "object"
		//	}
		"stack_filters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FilteringCriteria
				"filtering_criteria": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Attribute to specify the filtering behavior. ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match",
					Optional:    true,
					Computed:    true,
					Default:     stringdefault.StaticString("ALL"),
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ALL",
							"ANY",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StackNames
				"stack_names": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Exclude
						"exclude": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "List of stack names that the hook is going to be excluded from",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 50),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthAtMost(128),
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z][-a-zA-Z0-9]*$"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Include
						"include": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "List of stack names that the hook is going to target",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 50),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthAtMost(128),
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z][-a-zA-Z0-9]*$"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "List of stack names as filters",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StackRoles
				"stack_roles": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Exclude
						"exclude": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "List of stack roles that the hook is going to be excluded from",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 50),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthAtMost(256),
									stringvalidator.RegexMatches(regexp.MustCompile("arn:.+:iam::[0-9]{12}:role/.+"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Include
						"include": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "List of stack roles that the hook is going to target",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 50),
								setvalidator.ValueStringsAre(
									stringvalidator.LengthAtMost(256),
									stringvalidator.RegexMatches(regexp.MustCompile("arn:.+:iam::[0-9]{12}:role/.+"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "List of stack roles that are performing the stack operations.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Filters to allow hooks to target specific stack attributes",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetFilters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Attribute to specify which targets should invoke the hook",
		//	  "properties": {
		//	    "Actions": {
		//	      "description": "List of actions that the hook is going to target",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Target actions are the type of operation hooks will be executed at.",
		//	        "enum": [
		//	          "CREATE",
		//	          "UPDATE",
		//	          "DELETE"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "InvocationPoints": {
		//	      "description": "List of invocation points that the hook is going to target",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Invocation points are the point in provisioning workflow where hooks will be executed.",
		//	        "enum": [
		//	          "PRE_PROVISION"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TargetNames": {
		//	      "description": "List of type names that the hook is going to target",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Type name of hook target. Hook targets are the destination where hooks will be invoked against.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "Targets": {
		//	      "description": "List of hook targets",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Hook targets are the destination where hooks will be invoked against.",
		//	        "properties": {
		//	          "Action": {
		//	            "description": "Target actions are the type of operation hooks will be executed at.",
		//	            "enum": [
		//	              "CREATE",
		//	              "UPDATE",
		//	              "DELETE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "InvocationPoint": {
		//	            "description": "Invocation points are the point in provisioning workflow where hooks will be executed.",
		//	            "enum": [
		//	              "PRE_PROVISION"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "TargetName": {
		//	            "description": "Type name of hook target. Hook targets are the destination where hooks will be invoked against.",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "pattern": "",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "TargetName",
		//	          "Action",
		//	          "InvocationPoint"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"target_filters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Actions
				"actions": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "List of actions that the hook is going to target",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 50),
						setvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"CREATE",
								"UPDATE",
								"DELETE",
							),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: InvocationPoints
				"invocation_points": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "List of invocation points that the hook is going to target",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 50),
						setvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"PRE_PROVISION",
							),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TargetNames
				"target_names": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "List of type names that the hook is going to target",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 50),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 256),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Targets
				"targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Action
							"action": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Target actions are the type of operation hooks will be executed at.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"CREATE",
										"UPDATE",
										"DELETE",
									),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: InvocationPoint
							"invocation_point": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Invocation points are the point in provisioning workflow where hooks will be executed.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"PRE_PROVISION",
									),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: TargetName
							"target_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Type name of hook target. Hook targets are the destination where hooks will be invoked against.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 256),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of hook targets",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 50),
						listvalidator.UniqueValues(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Attribute to specify which targets should invoke the hook",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetOperations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Which operations should this Hook run against? Resource changes, stacks or change sets.",
		//	  "items": {
		//	    "description": "Which operations should this Hook run against? Resource changes, stacks or change sets.",
		//	    "enum": [
		//	      "RESOURCE",
		//	      "STACK",
		//	      "CHANGE_SET"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"target_operations": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Which operations should this Hook run against? Resource changes, stacks or change sets.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"RESOURCE",
						"STACK",
						"CHANGE_SET",
					),
				),
			}, /*END VALIDATORS*/
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
		Description: "This is a CloudFormation resource for activating the first-party AWS::Hooks::GuardHook.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::GuardHook").WithTerraformTypeName("awscc_cloudformation_guard_hook")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":             "Action",
		"actions":            "Actions",
		"alias":              "Alias",
		"exclude":            "Exclude",
		"execution_role":     "ExecutionRole",
		"failure_mode":       "FailureMode",
		"filtering_criteria": "FilteringCriteria",
		"hook_arn":           "HookArn",
		"hook_status":        "HookStatus",
		"include":            "Include",
		"input_params":       "InputParams",
		"invocation_point":   "InvocationPoint",
		"invocation_points":  "InvocationPoints",
		"log_bucket":         "LogBucket",
		"options":            "Options",
		"rule_location":      "RuleLocation",
		"stack_filters":      "StackFilters",
		"stack_names":        "StackNames",
		"stack_roles":        "StackRoles",
		"target_filters":     "TargetFilters",
		"target_name":        "TargetName",
		"target_names":       "TargetNames",
		"target_operations":  "TargetOperations",
		"targets":            "Targets",
		"uri":                "Uri",
		"version_id":         "VersionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}