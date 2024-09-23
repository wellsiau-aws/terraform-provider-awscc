// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package medialive

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
	registry.AddDataSourceFactory("awscc_medialive_event_bridge_rule_template", eventBridgeRuleTemplateDataSource)
}

// eventBridgeRuleTemplateDataSource returns the Terraform awscc_medialive_event_bridge_rule_template data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaLive::EventBridgeRuleTemplate resource.
func eventBridgeRuleTemplateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An eventbridge rule template's ARN (Amazon Resource Name)",
		//	  "pattern": "^arn:.+:medialive:.+:eventbridge-rule-template:.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An eventbridge rule template's ARN (Amazon Resource Name)",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Placeholder documentation for __timestampIso8601",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Placeholder documentation for __timestampIso8601",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource's optional description.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource's optional description.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EventTargets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Placeholder documentation for __listOfEventBridgeRuleTemplateTarget",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The target to which to send matching events.",
		//	    "properties": {
		//	      "Arn": {
		//	        "description": "Target ARNs must be either an SNS topic or CloudWatch log group.",
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "pattern": "^arn.+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Arn"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"event_targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Target ARNs must be either an SNS topic or CloudWatch log group.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Placeholder documentation for __listOfEventBridgeRuleTemplateTarget",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EventType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of event to match with the rule.",
		//	  "enum": [
		//	    "MEDIALIVE_MULTIPLEX_ALERT",
		//	    "MEDIALIVE_MULTIPLEX_STATE_CHANGE",
		//	    "MEDIALIVE_CHANNEL_ALERT",
		//	    "MEDIALIVE_CHANNEL_INPUT_CHANGE",
		//	    "MEDIALIVE_CHANNEL_STATE_CHANGE",
		//	    "MEDIAPACKAGE_INPUT_NOTIFICATION",
		//	    "MEDIAPACKAGE_KEY_PROVIDER_NOTIFICATION",
		//	    "MEDIAPACKAGE_HARVEST_JOB_NOTIFICATION",
		//	    "SIGNAL_MAP_ACTIVE_ALARM",
		//	    "MEDIACONNECT_ALERT",
		//	    "MEDIACONNECT_SOURCE_HEALTH",
		//	    "MEDIACONNECT_OUTPUT_HEALTH",
		//	    "MEDIACONNECT_FLOW_STATUS_CHANGE"
		//	  ],
		//	  "type": "string"
		//	}
		"event_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of event to match with the rule.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`",
		//	  "maxLength": 11,
		//	  "minLength": 7,
		//	  "pattern": "^(aws-)?[0-9]{7}$",
		//	  "type": "string"
		//	}
		"group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An eventbridge rule template group's identifier. Can be either be its id or current name.",
		//	  "pattern": "^[^\\s]+$",
		//	  "type": "string"
		//	}
		"group_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An eventbridge rule template group's identifier. Can be either be its id or current name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An eventbridge rule template's id. AWS provided templates have ids that start with `aws-`",
		//	  "maxLength": 11,
		//	  "minLength": 7,
		//	  "pattern": "^(aws-)?[0-9]{7}$",
		//	  "type": "string"
		//	}
		"event_bridge_rule_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An eventbridge rule template's id. AWS provided templates have ids that start with `aws-`",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Identifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Placeholder documentation for __string",
		//	  "type": "string"
		//	}
		"identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Placeholder documentation for __string",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModifiedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Placeholder documentation for __timestampIso8601",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"modified_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Placeholder documentation for __timestampIso8601",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^[^\\s]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A resource's name. Names must be unique within the scope of a resource type in a specific region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Represents the tags associated with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "Placeholder documentation for __string",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Represents the tags associated with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaLive::EventBridgeRuleTemplate",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaLive::EventBridgeRuleTemplate").WithTerraformTypeName("awscc_medialive_event_bridge_rule_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                           "Arn",
		"created_at":                    "CreatedAt",
		"description":                   "Description",
		"event_bridge_rule_template_id": "Id",
		"event_targets":                 "EventTargets",
		"event_type":                    "EventType",
		"group_id":                      "GroupId",
		"group_identifier":              "GroupIdentifier",
		"identifier":                    "Identifier",
		"modified_at":                   "ModifiedAt",
		"name":                          "Name",
		"tags":                          "Tags",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}