// Code generated by generators/resource/main.go; DO NOT EDIT.

package auditmanager

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_auditmanager_assessment", assessment)
}

// assessment returns the Terraform aws_auditmanager_assessment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AuditManager::Assessment resource type.
func assessment(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the assessment.",
			     "maxLength": 2048,
			     "minLength": 20,
			     "pattern": "",
			     "$ref": "#/definitions/AssessmentArn",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the assessment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"assessment_id": {
			// Property: AssessmentId
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 36,
			     "minLength": 36,
			     "pattern": "",
			     "$ref": "#/definitions/UUID",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"assessment_reports_destination": {
			// Property: AssessmentReportsDestination
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The destination in which evidence reports are stored for the specified assessment.",
			     "properties": {
			       "Destination": {
			         "description": "The URL of the specified Amazon S3 bucket.",
			         "$ref": "#/definitions/S3Url",
			         "type": "string"
			       },
			       "DestinationType": {
			         "description": "The destination type, such as Amazon S3.",
			         "enum": [
			           "S3"
			         ],
			         "$ref": "#/definitions/AssessmentReportDestinationType",
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/AssessmentReportsDestination",
			     "type": "object"
			   }
			*/
			Description: "The destination in which evidence reports are stored for the specified assessment.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"destination": {
						// Property: Destination
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The URL of the specified Amazon S3 bucket.",
						     "$ref": "#/definitions/S3Url",
						     "type": "string"
						   }
						*/
						Description: "The URL of the specified Amazon S3 bucket.",
						Type:        types.StringType,
						Optional:    true,
					},
					"destination_type": {
						// Property: DestinationType
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The destination type, such as Amazon S3.",
						     "enum": [
						       "S3"
						     ],
						     "$ref": "#/definitions/AssessmentReportDestinationType",
						     "type": "string"
						   }
						*/
						Description: "The destination type, such as Amazon S3.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"aws_account": {
			// Property: AwsAccount
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The AWS account associated with the assessment.",
			     "properties": {
			       "EmailAddress": {
			         "description": "The unique identifier for the email account.",
			         "maxLength": 320,
			         "minLength": 1,
			         "pattern": "",
			         "$ref": "#/definitions/EmailAddress",
			         "type": "string"
			       },
			       "Id": {
			         "description": "The identifier for the specified AWS account.",
			         "maxLength": 12,
			         "minLength": 12,
			         "pattern": "",
			         "$ref": "#/definitions/AccountId",
			         "type": "string"
			       },
			       "Name": {
			         "description": "The name of the specified AWS account.",
			         "maxLength": 50,
			         "minLength": 1,
			         "pattern": "",
			         "$ref": "#/definitions/AccountName",
			         "type": "string"
			       }
			     },
			     "$ref": "#/definitions/AWSAccount",
			     "type": "object"
			   }
			*/
			Description: "The AWS account associated with the assessment.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"email_address": {
						// Property: EmailAddress
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The unique identifier for the email account.",
						     "maxLength": 320,
						     "minLength": 1,
						     "pattern": "",
						     "$ref": "#/definitions/EmailAddress",
						     "type": "string"
						   }
						*/
						Description: "The unique identifier for the email account.",
						Type:        types.StringType,
						Optional:    true,
					},
					"id": {
						// Property: Id
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The identifier for the specified AWS account.",
						     "maxLength": 12,
						     "minLength": 12,
						     "pattern": "",
						     "$ref": "#/definitions/AccountId",
						     "type": "string"
						   }
						*/
						Description: "The identifier for the specified AWS account.",
						Type:        types.StringType,
						Optional:    true,
					},
					"name": {
						// Property: Name
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The name of the specified AWS account.",
						     "maxLength": 50,
						     "minLength": 1,
						     "pattern": "",
						     "$ref": "#/definitions/AccountName",
						     "type": "string"
						   }
						*/
						Description: "The name of the specified AWS account.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// AwsAccount is a force-new attribute.
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The sequence of characters that identifies when the event occurred.",
			     "$ref": "#/definitions/Timestamp",
			     "type": "number"
			   }
			*/
			Description: "The sequence of characters that identifies when the event occurred.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"delegations": {
			// Property: Delegations
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The list of delegations.",
			     "items": {
			       "additionalProperties": false,
			       "description": "The assignment of a control set to a delegate for review.",
			       "properties": {
			         "AssessmentId": {
			           "maxLength": 36,
			           "minLength": 36,
			           "pattern": "",
			           "$ref": "#/definitions/UUID",
			           "type": "string"
			         },
			         "AssessmentName": {
			           "description": "The name of the related assessment.",
			           "maxLength": 127,
			           "minLength": 1,
			           "pattern": "",
			           "$ref": "#/definitions/AssessmentName",
			           "type": "string"
			         },
			         "Comment": {
			           "description": "The comment related to the delegation.",
			           "maxLength": 350,
			           "pattern": "",
			           "$ref": "#/definitions/DelegationComment",
			           "type": "string"
			         },
			         "ControlSetId": {
			           "description": "The identifier for the specified control set.",
			           "maxLength": 300,
			           "minLength": 1,
			           "pattern": "",
			           "$ref": "#/definitions/ControlSetId",
			           "type": "string"
			         },
			         "CreatedBy": {
			           "description": "The IAM user or role that performed the action.",
			           "maxLength": 100,
			           "minLength": 1,
			           "pattern": "",
			           "$ref": "#/definitions/CreatedBy",
			           "type": "string"
			         },
			         "CreationTime": {
			           "description": "The sequence of characters that identifies when the event occurred.",
			           "$ref": "#/definitions/Timestamp",
			           "type": "number"
			         },
			         "Id": {
			           "maxLength": 36,
			           "minLength": 36,
			           "pattern": "",
			           "$ref": "#/definitions/UUID",
			           "type": "string"
			         },
			         "LastUpdated": {
			           "description": "The sequence of characters that identifies when the event occurred.",
			           "$ref": "#/definitions/Timestamp",
			           "type": "number"
			         },
			         "RoleArn": {
			           "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
			           "maxLength": 2048,
			           "minLength": 20,
			           "pattern": "",
			           "$ref": "#/definitions/IamArn",
			           "type": "string"
			         },
			         "RoleType": {
			           "description": " The IAM role type.",
			           "enum": [
			             "PROCESS_OWNER",
			             "RESOURCE_OWNER"
			           ],
			           "$ref": "#/definitions/RoleType",
			           "type": "string"
			         },
			         "Status": {
			           "description": "The status of the delegation.",
			           "enum": [
			             "IN_PROGRESS",
			             "UNDER_REVIEW",
			             "COMPLETE"
			           ],
			           "$ref": "#/definitions/DelegationStatus",
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Delegation",
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The list of delegations.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"assessment_id": {
						// Property: AssessmentId
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 36,
						     "minLength": 36,
						     "pattern": "",
						     "$ref": "#/definitions/UUID",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"assessment_name": {
						// Property: AssessmentName
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The name of the related assessment.",
						     "maxLength": 127,
						     "minLength": 1,
						     "pattern": "",
						     "$ref": "#/definitions/AssessmentName",
						     "type": "string"
						   }
						*/
						Description: "The name of the related assessment.",
						Type:        types.StringType,
						Optional:    true,
					},
					"comment": {
						// Property: Comment
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The comment related to the delegation.",
						     "maxLength": 350,
						     "pattern": "",
						     "$ref": "#/definitions/DelegationComment",
						     "type": "string"
						   }
						*/
						Description: "The comment related to the delegation.",
						Type:        types.StringType,
						Optional:    true,
					},
					"control_set_id": {
						// Property: ControlSetId
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The identifier for the specified control set.",
						     "maxLength": 300,
						     "minLength": 1,
						     "pattern": "",
						     "$ref": "#/definitions/ControlSetId",
						     "type": "string"
						   }
						*/
						Description: "The identifier for the specified control set.",
						Type:        types.StringType,
						Optional:    true,
					},
					"created_by": {
						// Property: CreatedBy
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The IAM user or role that performed the action.",
						     "maxLength": 100,
						     "minLength": 1,
						     "pattern": "",
						     "$ref": "#/definitions/CreatedBy",
						     "type": "string"
						   }
						*/
						Description: "The IAM user or role that performed the action.",
						Type:        types.StringType,
						Optional:    true,
					},
					"creation_time": {
						// Property: CreationTime
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The sequence of characters that identifies when the event occurred.",
						     "$ref": "#/definitions/Timestamp",
						     "type": "number"
						   }
						*/
						Description: "The sequence of characters that identifies when the event occurred.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"id": {
						// Property: Id
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 36,
						     "minLength": 36,
						     "pattern": "",
						     "$ref": "#/definitions/UUID",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Optional: true,
					},
					"last_updated": {
						// Property: LastUpdated
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The sequence of characters that identifies when the event occurred.",
						     "$ref": "#/definitions/Timestamp",
						     "type": "number"
						   }
						*/
						Description: "The sequence of characters that identifies when the event occurred.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"role_arn": {
						// Property: RoleArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
						     "maxLength": 2048,
						     "minLength": 20,
						     "pattern": "",
						     "$ref": "#/definitions/IamArn",
						     "type": "string"
						   }
						*/
						Description: "The Amazon Resource Name (ARN) of the IAM user or role.",
						Type:        types.StringType,
						Optional:    true,
					},
					"role_type": {
						// Property: RoleType
						// CloudFormation resource type schema:
						/*
						   {
						     "description": " The IAM role type.",
						     "enum": [
						       "PROCESS_OWNER",
						       "RESOURCE_OWNER"
						     ],
						     "$ref": "#/definitions/RoleType",
						     "type": "string"
						   }
						*/
						Description: " The IAM role type.",
						Type:        types.StringType,
						Optional:    true,
					},
					"status": {
						// Property: Status
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The status of the delegation.",
						     "enum": [
						       "IN_PROGRESS",
						       "UNDER_REVIEW",
						       "COMPLETE"
						     ],
						     "$ref": "#/definitions/DelegationStatus",
						     "type": "string"
						   }
						*/
						Description: "The status of the delegation.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The description of the specified assessment.",
			     "$ref": "#/definitions/AssessmentDescription",
			     "type": "string"
			   }
			*/
			Description: "The description of the specified assessment.",
			Type:        types.StringType,
			Optional:    true,
			// Description is a write-only attribute.
		},
		"framework_id": {
			// Property: FrameworkId
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The identifier for the specified framework.",
			     "maxLength": 36,
			     "minLength": 32,
			     "pattern": "",
			     "$ref": "#/definitions/FrameworkId",
			     "type": "string"
			   }
			*/
			Description: "The identifier for the specified framework.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// FrameworkId is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the related assessment.",
			     "maxLength": 127,
			     "minLength": 1,
			     "pattern": "",
			     "$ref": "#/definitions/AssessmentName",
			     "type": "string"
			   }
			*/
			Description: "The name of the related assessment.",
			Type:        types.StringType,
			Optional:    true,
			// Name is a write-only attribute.
		},
		"roles": {
			// Property: Roles
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The list of roles for the specified assessment.",
			     "items": {
			       "additionalProperties": false,
			       "description": "The wrapper that contains AWS Audit Manager role information, such as the role type and IAM ARN.",
			       "properties": {
			         "RoleArn": {
			           "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
			           "maxLength": 2048,
			           "minLength": 20,
			           "pattern": "",
			           "$ref": "#/definitions/IamArn",
			           "type": "string"
			         },
			         "RoleType": {
			           "description": " The IAM role type.",
			           "enum": [
			             "PROCESS_OWNER",
			             "RESOURCE_OWNER"
			           ],
			           "$ref": "#/definitions/RoleType",
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Role",
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The list of roles for the specified assessment.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"role_arn": {
						// Property: RoleArn
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The Amazon Resource Name (ARN) of the IAM user or role.",
						     "maxLength": 2048,
						     "minLength": 20,
						     "pattern": "",
						     "$ref": "#/definitions/IamArn",
						     "type": "string"
						   }
						*/
						Description: "The Amazon Resource Name (ARN) of the IAM user or role.",
						Type:        types.StringType,
						Optional:    true,
					},
					"role_type": {
						// Property: RoleType
						// CloudFormation resource type schema:
						/*
						   {
						     "description": " The IAM role type.",
						     "enum": [
						       "PROCESS_OWNER",
						       "RESOURCE_OWNER"
						     ],
						     "$ref": "#/definitions/RoleType",
						     "type": "string"
						   }
						*/
						Description: " The IAM role type.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "description": "The wrapper that contains the AWS accounts and AWS services in scope for the assessment.",
			     "properties": {
			       "AwsAccounts": {
			         "description": "The AWS accounts included in scope.",
			         "items": {
			           "additionalProperties": false,
			           "description": "The AWS account associated with the assessment.",
			           "properties": {
			             "EmailAddress": {
			               "description": "The unique identifier for the email account.",
			               "maxLength": 320,
			               "minLength": 1,
			               "pattern": "",
			               "$ref": "#/definitions/EmailAddress",
			               "type": "string"
			             },
			             "Id": {
			               "description": "The identifier for the specified AWS account.",
			               "maxLength": 12,
			               "minLength": 12,
			               "pattern": "",
			               "$ref": "#/definitions/AccountId",
			               "type": "string"
			             },
			             "Name": {
			               "description": "The name of the specified AWS account.",
			               "maxLength": 50,
			               "minLength": 1,
			               "pattern": "",
			               "$ref": "#/definitions/AccountName",
			               "type": "string"
			             }
			           },
			           "$ref": "#/definitions/AWSAccount",
			           "type": "object"
			         },
			         "type": "array"
			       },
			       "AwsServices": {
			         "description": "The AWS services included in scope.",
			         "items": {
			           "additionalProperties": false,
			           "description": "An AWS service such as Amazon S3, AWS CloudTrail, and so on.",
			           "properties": {
			             "ServiceName": {
			               "description": "The name of the AWS service.",
			               "$ref": "#/definitions/AWSServiceName",
			               "type": "string"
			             }
			           },
			           "$ref": "#/definitions/AWSService",
			           "type": "object"
			         },
			         "type": "array"
			       }
			     },
			     "$ref": "#/definitions/Scope",
			     "type": "object"
			   }
			*/
			Description: "The wrapper that contains the AWS accounts and AWS services in scope for the assessment.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"aws_accounts": {
						// Property: AwsAccounts
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The AWS accounts included in scope.",
						     "items": {
						       "additionalProperties": false,
						       "description": "The AWS account associated with the assessment.",
						       "properties": {
						         "EmailAddress": {
						           "description": "The unique identifier for the email account.",
						           "maxLength": 320,
						           "minLength": 1,
						           "pattern": "",
						           "$ref": "#/definitions/EmailAddress",
						           "type": "string"
						         },
						         "Id": {
						           "description": "The identifier for the specified AWS account.",
						           "maxLength": 12,
						           "minLength": 12,
						           "pattern": "",
						           "$ref": "#/definitions/AccountId",
						           "type": "string"
						         },
						         "Name": {
						           "description": "The name of the specified AWS account.",
						           "maxLength": 50,
						           "minLength": 1,
						           "pattern": "",
						           "$ref": "#/definitions/AccountName",
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/AWSAccount",
						       "type": "object"
						     },
						     "type": "array"
						   }
						*/
						Description: "The AWS accounts included in scope.",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"email_address": {
									// Property: EmailAddress
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The unique identifier for the email account.",
									     "maxLength": 320,
									     "minLength": 1,
									     "pattern": "",
									     "$ref": "#/definitions/EmailAddress",
									     "type": "string"
									   }
									*/
									Description: "The unique identifier for the email account.",
									Type:        types.StringType,
									Optional:    true,
								},
								"id": {
									// Property: Id
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The identifier for the specified AWS account.",
									     "maxLength": 12,
									     "minLength": 12,
									     "pattern": "",
									     "$ref": "#/definitions/AccountId",
									     "type": "string"
									   }
									*/
									Description: "The identifier for the specified AWS account.",
									Type:        types.StringType,
									Optional:    true,
								},
								"name": {
									// Property: Name
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The name of the specified AWS account.",
									     "maxLength": 50,
									     "minLength": 1,
									     "pattern": "",
									     "$ref": "#/definitions/AccountName",
									     "type": "string"
									   }
									*/
									Description: "The name of the specified AWS account.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"aws_services": {
						// Property: AwsServices
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The AWS services included in scope.",
						     "items": {
						       "additionalProperties": false,
						       "description": "An AWS service such as Amazon S3, AWS CloudTrail, and so on.",
						       "properties": {
						         "ServiceName": {
						           "description": "The name of the AWS service.",
						           "$ref": "#/definitions/AWSServiceName",
						           "type": "string"
						         }
						       },
						       "$ref": "#/definitions/AWSService",
						       "type": "object"
						     },
						     "type": "array"
						   }
						*/
						Description: "The AWS services included in scope.",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"service_name": {
									// Property: ServiceName
									// CloudFormation resource type schema:
									/*
									   {
									     "description": "The name of the AWS service.",
									     "$ref": "#/definitions/AWSServiceName",
									     "type": "string"
									   }
									*/
									Description: "The name of the AWS service.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The status of the specified assessment. ",
			     "enum": [
			       "ACTIVE",
			       "INACTIVE"
			     ],
			     "$ref": "#/definitions/AssessmentStatus",
			     "type": "string"
			   }
			*/
			Description: "The status of the specified assessment. ",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags associated with the assessment.",
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			           "maxLength": 256,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Key",
			         "Value"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Description: "The tags associated with the assessment.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						     "maxLength": 256,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "An entity that defines the scope of audit evidence collected by AWS Audit Manager.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AuditManager::Assessment").WithTerraformTypeName("aws_auditmanager_assessment").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Name",
		"/properties/Description",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_auditmanager_assessment", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
