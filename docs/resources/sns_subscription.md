
---
page_title: "awscc_sns_subscription Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SNS::Subscription
---

# awscc_sns_subscription (Resource)

Resource Type definition for AWS::SNS::Subscription

## Example Usage

### SNS Lambda Subscription with Message Filtering

Creates an SNS subscription that triggers a Lambda function with message filtering for specific event types (order_placed and order_cancelled) using MessageAttributes scope.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a SNS topic
resource "awscc_sns_topic" "example" {
  topic_name = "example-topic"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for Lambda
data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "lambda_role" {
  name               = "example-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
  tags = {
    "Modified By" = "AWSCC"
  }
}

# Create IAM policy for Lambda logging
data "aws_iam_policy_document" "lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/*"]
  }
}

resource "aws_iam_policy" "lambda_logging" {
  name   = "lambda-logging"
  policy = data.aws_iam_policy_document.lambda_logging.json
  tags = {
    "Modified By" = "AWSCC"
  }
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

resource "aws_lambda_function" "example" {
  filename      = "lambda_function.zip"
  function_name = "example-function"
  role          = aws_iam_role.lambda_role.arn
  handler       = "index.handler"
  runtime       = "nodejs18.x"

  lifecycle {
    ignore_changes = [filename]
  }
}

resource "aws_lambda_permission" "sns" {
  statement_id  = "AllowSNSInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.example.function_name
  principal     = "sns.amazonaws.com"
  source_arn    = "arn:aws:sns:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${awscc_sns_topic.example.topic_name}"
}

resource "awscc_sns_subscription" "example" {
  protocol  = "lambda"
  topic_arn = "arn:aws:sns:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${awscc_sns_topic.example.topic_name}"
  endpoint  = aws_lambda_function.example.arn
  filter_policy = jsonencode({
    "event_type" : ["order_placed", "order_cancelled"]
  })
  filter_policy_scope = "MessageAttributes"
  region              = data.aws_region.current.name

  depends_on = [aws_lambda_permission.sns]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `protocol` (String) The subscription's protocol.
- `topic_arn` (String) The ARN of the topic to subscribe to.

### Optional

- `delivery_policy` (String) The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.
- `endpoint` (String) The subscription's endpoint. The endpoint value depends on the protocol that you specify.
- `filter_policy` (String) The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.
- `filter_policy_scope` (String) This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.
- `raw_message_delivery` (Boolean) When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.
- `redrive_policy` (String) When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.
- `region` (String) For cross-region subscriptions, the region in which the topic resides.If no region is specified, AWS CloudFormation uses the region of the caller as the default.
- `replay_policy` (String) Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.
- `subscription_role_arn` (String) This property applies only to Amazon Data Firehose delivery stream subscriptions.

### Read-Only

- `arn` (String) Arn of the subscription
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_sns_subscription.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_sns_subscription.example "arn"
```
