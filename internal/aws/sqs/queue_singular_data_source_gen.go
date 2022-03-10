// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sqs_queue", queueDataSourceType)
}

// queueDataSourceType returns the Terraform awscc_sqs_queue data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SQS::Queue resource type.
func queueDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the queue.",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the queue.",
			Type:        types.StringType,
			Computed:    true,
		},
		"content_based_deduplication": {
			// Property: ContentBasedDeduplication
			// CloudFormation resource type schema:
			// {
			//   "description": "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.",
			//   "type": "boolean"
			// }
			Description: "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"deduplication_scope": {
			// Property: DeduplicationScope
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.",
			//   "type": "string"
			// }
			Description: "Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.",
			Type:        types.StringType,
			Computed:    true,
		},
		"delay_seconds": {
			// Property: DelaySeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.",
			//   "type": "integer"
			// }
			Description: "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"fifo_queue": {
			// Property: FifoQueue
			// CloudFormation resource type schema:
			// {
			//   "description": "If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.",
			//   "type": "boolean"
			// }
			Description: "If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"fifo_throughput_limit": {
			// Property: FifoThroughputLimit
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.",
			//   "type": "string"
			// }
			Description: "Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.",
			Type:        types.StringType,
			Computed:    true,
		},
		"kms_data_key_reuse_period_seconds": {
			// Property: KmsDataKeyReusePeriodSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).",
			//   "type": "integer"
			// }
			Description: "The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"kms_master_key_id": {
			// Property: KmsMasterKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.",
			//   "type": "string"
			// }
			Description: "The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.",
			Type:        types.StringType,
			Computed:    true,
		},
		"maximum_message_size": {
			// Property: MaximumMessageSize
			// CloudFormation resource type schema:
			// {
			//   "description": "The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).",
			//   "type": "integer"
			// }
			Description: "The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"message_retention_period": {
			// Property: MessageRetentionPeriod
			// CloudFormation resource type schema:
			// {
			//   "description": "The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).",
			//   "type": "integer"
			// }
			Description: "The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"queue_name": {
			// Property: QueueName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.",
			//   "type": "string"
			// }
			Description: "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.",
			Type:        types.StringType,
			Computed:    true,
		},
		"queue_url": {
			// Property: QueueUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "URL of the source queue.",
			//   "type": "string"
			// }
			Description: "URL of the source queue.",
			Type:        types.StringType,
			Computed:    true,
		},
		"receive_message_wait_time_seconds": {
			// Property: ReceiveMessageWaitTimeSeconds
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.",
			//   "type": "integer"
			// }
			Description: "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"redrive_allow_policy": {
			// Property: RedriveAllowPolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.",
			//   "type": "string"
			// }
			Description: "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.",
			Type:        types.StringType,
			Computed:    true,
		},
		"redrive_policy": {
			// Property: RedrivePolicy
			// CloudFormation resource type schema:
			// {
			//   "description": "A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.",
			//   "type": "string"
			// }
			Description: "A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags that you attach to this queue.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			Description: "The tags that you attach to this queue.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"visibility_timeout": {
			// Property: VisibilityTimeout
			// CloudFormation resource type schema:
			// {
			//   "description": "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.",
			//   "type": "integer"
			// }
			Description: "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.",
			Type:        types.Int64Type,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SQS::Queue",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SQS::Queue").WithTerraformTypeName("awscc_sqs_queue")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                               "Arn",
		"content_based_deduplication":       "ContentBasedDeduplication",
		"deduplication_scope":               "DeduplicationScope",
		"delay_seconds":                     "DelaySeconds",
		"fifo_queue":                        "FifoQueue",
		"fifo_throughput_limit":             "FifoThroughputLimit",
		"key":                               "Key",
		"kms_data_key_reuse_period_seconds": "KmsDataKeyReusePeriodSeconds",
		"kms_master_key_id":                 "KmsMasterKeyId",
		"maximum_message_size":              "MaximumMessageSize",
		"message_retention_period":          "MessageRetentionPeriod",
		"queue_name":                        "QueueName",
		"queue_url":                         "QueueUrl",
		"receive_message_wait_time_seconds": "ReceiveMessageWaitTimeSeconds",
		"redrive_allow_policy":              "RedriveAllowPolicy",
		"redrive_policy":                    "RedrivePolicy",
		"tags":                              "Tags",
		"value":                             "Value",
		"visibility_timeout":                "VisibilityTimeout",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}